use std::ops::Deref;
use std::thread::{self, sleep, JoinHandle};
use std::io;
use std::sync::{Mutex, Arc};
use std::sync::mpsc::{channel, Sender, Receiver, TryRecvError};
use std::time::Duration;

type ThreadFunc = Box<dyn FnOnce(()) -> Result<(),String> + Send>;

#[derive(Debug, Clone, PartialEq)]
pub enum ThreadState {
    Running,
    Idle,
    Error,
    Drop
}

#[derive(Debug)]
pub struct ThreadImpl{
    join_handle : Option<JoinHandle<()>>,
    func_sender : Sender<ThreadFunc>,
    state : Arc<Mutex<ThreadState>>
}

impl Drop for ThreadImpl {
    fn drop(&mut self) {
        *self.state.lock().unwrap() = ThreadState::Drop;
    }
}

impl ThreadImpl {
    fn switch_states(s_mutex : &Arc<Mutex<ThreadState>>, s : ThreadState) {
        let mut state = s_mutex.lock().unwrap();
        if *state == ThreadState::Drop {
            return;
        }
        *state = s;
    }
    fn thread_main(ptr :  Arc<Mutex<ThreadState>>, recv : Receiver<ThreadFunc>) {
        loop {
            {
                let state = ptr.lock().unwrap();
                if *state.deref() == ThreadState::Drop {
                    return;
                }
            }

            let f_ret = recv.try_recv();
            ThreadImpl::switch_states(&ptr, ThreadState::Running);

            if f_ret.is_err() {
                if f_ret.err().unwrap() == TryRecvError::Empty {
                    continue;
                }
                ThreadImpl::switch_states(&ptr, ThreadState::Error);
                return;
            }

            let f = f_ret.unwrap();
            let _ = f(());
            ThreadImpl::switch_states(&ptr, ThreadState::Idle);
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
            state : Arc::new(Mutex::new(ThreadState::Idle))
        };
        obj.join_handle = Some(ThreadImpl::thread_handler(receiver, obj.state.clone()));
        return obj;
    }

    pub fn call_push(&mut self, thread_func : ThreadFunc) -> io::Result<()>  {
        self.func_sender.send(Box::new(thread_func));
        Ok(())
    }

    pub fn get_state(&self) -> ThreadState {
        let state = self.state.lock().unwrap();
        return state.clone()
    }
}