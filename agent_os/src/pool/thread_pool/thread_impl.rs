use std::ops::Deref;
use std::thread::{self, sleep, JoinHandle};
use std::io;
use std::sync::{Mutex, Arc};
use std::sync::mpsc::{channel, Sender, Receiver, TryRecvError};
use std::time::Duration;

type ThreadFunc = Box<dyn FnOnce(()) -> Result<(),String> + Send>;

#[derive(Debug, Clone, PartialEq)]
pub enum ThreadSignal {
    Running,
    Idle,
    Error,
    Drop
}

impl Default for ThreadSignal {
    fn default() -> Self {
        ThreadSignal::Idle
    }
}

#[derive(Debug, Default)]
struct ThreadState {
    signal : ThreadSignal,
    sender_cnt : usize
}

impl ThreadState {
    pub(super) fn set_signal(&mut self, signal : ThreadSignal) {
        self.signal = signal;
    }
    pub(super) fn incr_sender_cnt(&mut self) {
        self.sender_cnt += 1;
    }
    pub(super) fn decr_sender_cnt(&mut self) {
        self.sender_cnt -= 1;
    }
    pub(super) fn get_signal(&self) -> ThreadSignal {
        self.signal.clone()
    }
    pub(super) fn get_sender_cnt(&self) -> usize {
        self.sender_cnt
    }
}

#[derive(Debug)]
pub struct ThreadImpl{
    join_handle : Option<JoinHandle<()>>,
    func_sender : Sender<ThreadFunc>,
    state : Arc<Mutex<ThreadState>>
}

impl Drop for ThreadImpl {
    fn drop(&mut self) {
        let mut s = self.state.lock().unwrap();
        s.set_signal(ThreadSignal::Drop);
    }
}

impl ThreadImpl {
    fn switch_states(s_mutex : &Arc<Mutex<ThreadState>>, s : ThreadSignal) {
        let mut state = s_mutex.lock().unwrap();
        if state.get_signal() == ThreadSignal::Drop {
            return;
        }
        state.set_signal(s);
    }
    fn thread_main(ptr :  Arc<Mutex<ThreadState>>, recv : Receiver<ThreadFunc>) {
        loop {
            {
                let state = ptr.lock().unwrap();
                if state.get_signal() == ThreadSignal::Drop {
                    return;
                }
            }

            let f_ret = recv.try_recv();
            ThreadImpl::switch_states(&ptr, ThreadSignal::Running);

            if f_ret.is_err() {
                if f_ret.err().unwrap() == TryRecvError::Empty {
                    continue;
                }
                ThreadImpl::switch_states(&ptr, ThreadSignal::Error);
                return;
            }

            let f = f_ret.unwrap();
            let _ = f(());
            ThreadImpl::switch_states(&ptr, ThreadSignal::Idle);
            sleep(Duration::from_micros(10));
        }
    }

    fn thread_handler(recv : Receiver<ThreadFunc>, state_p : Arc<Mutex<ThreadState>>) -> JoinHandle<()> {
        return thread::spawn(move || {
            ThreadImpl::thread_main(state_p, recv);
        });
    }

    pub fn new() -> ThreadImpl {
        let (sender, receiver) = channel::<ThreadFunc>();
        let mut obj = ThreadImpl {
            join_handle : None,
            func_sender : sender,
            state : Arc::new(Mutex::new(ThreadState::default()))
        };
        obj.join_handle = Some(ThreadImpl::thread_handler(receiver, obj.state.clone()));
        return obj;
    }

    pub fn call_push(&mut self, thread_func : ThreadFunc) -> io::Result<()>  {
        let mut s = self.state.lock().unwrap();
        s.incr_sender_cnt();
        drop(s);
        
        self.func_sender.send(Box::new(thread_func));
        Ok(())
    }

    pub fn get_signal(&self) -> ThreadSignal {
        let state = self.state.lock().unwrap();
        return state.get_signal()
    }
}