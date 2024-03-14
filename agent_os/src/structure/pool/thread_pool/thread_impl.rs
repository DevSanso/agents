use std::sync::atomic::{self,AtomicU16};
use std::thread::{self, sleep, JoinHandle};
use std::io;
use std::sync::{Mutex, Arc};
use std::sync::mpsc::{channel, Sender, Receiver, TryRecvError};
use std::time::Duration;

use crate::utils::convert_to_io_result;

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
    sender_cnt : AtomicU16
}

impl ThreadState {
    pub(super) fn set_signal(&mut self, signal : ThreadSignal) {
        self.signal = signal;
    }
    pub(super) fn incr_sender_cnt(&mut self) {
        self.sender_cnt.fetch_add(1, atomic::Ordering::SeqCst);
    }
    pub(super) fn decr_sender_cnt(&mut self) {
        self.sender_cnt.fetch_sub(1, atomic::Ordering::SeqCst);
    }
    pub(super) fn get_signal(&self) -> ThreadSignal {
        self.signal.clone()
    }
    pub(super) fn get_sender_cnt(&self) -> usize {
        self.sender_cnt.load(atomic::Ordering::Relaxed) as usize
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

            if f_ret.is_err() {
                if f_ret.err().unwrap() == TryRecvError::Empty {
                    continue;
                }
                ThreadImpl::switch_states(&ptr, ThreadSignal::Error);
                return;
            }
            ThreadImpl::switch_states(&ptr, ThreadSignal::Running);
            let mut locked = ptr.lock().unwrap();
            locked.decr_sender_cnt();
            drop(locked);

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
        
        let ret =self.func_sender.send(thread_func);
        convert_to_io_result!(result, ret)?;
        Ok(())
    }

    pub fn get_signal(&self) -> ThreadSignal {
        let state = self.state.lock().unwrap();
        return state.get_signal()
    }

    pub fn get_thread_func_count(&self) -> usize {
        return self.state.lock().unwrap().get_sender_cnt();
    }
}