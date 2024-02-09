use std::io;
use std::sync::{Mutex, Arc};

use crate::pool;
use crate::pool::thread_pool::thread_impl::{ThreadImpl,ThreadState};

#[derive(Debug)]
pub struct ThreadPool {
    max : usize,
    thread_p : Vec<ThreadImpl>,
}

impl Drop for ThreadPool {
    fn drop(&mut self) {
        for i in 0..self.thread_p.len() {
            drop(&self.thread_p[i]);
        }
        self.thread_p.clear();
        self.max = 0;
    }
}

impl ThreadPool {
    fn find_idle(&self) -> Option<usize> {

        for i in 0..self.thread_p.len() {
            let state = self.thread_p[i].get_state();
            if state == ThreadState::Idle {
                return Some(i);
            }
        }
        None
    }

    fn new_thread(&mut self) -> io::Result<usize> {
        if self.max <= self.alloc_size() {
            return Err(io::Error::new(io::ErrorKind::OutOfMemory, "thread pool is full"));
        }
        let t = ThreadImpl::new();
        self.thread_p.push(t);
        Ok(self.thread_p.len() - 1)
    }

    fn alloc_size(&self) -> usize {
        self.thread_p.len()
    }

    pub fn new(size : usize, max_size : usize) -> Arc<Mutex<Self>> {
        let mut tp = ThreadPool {max : max_size, thread_p : Vec::new()};

        for _ in 0..size {
            let t = ThreadImpl::new();
            tp.thread_p.push(t);
        }

        Arc::new(Mutex::new(tp))
    } 
}

impl pool::Pool<(), () , String> for ThreadPool {
    fn run_func<F : FnOnce(()) -> Result<(),String> + Send + 'static>(&mut self, arg : (), f : F) -> io::Result<()> {

        if self.alloc_size() >= self.max {
            return Err(io::Error::new(io::ErrorKind::InvalidInput, "index is out of range"));
        }
        let mut thread_index_opt = self.find_idle();
        
        if thread_index_opt.is_none() {
            thread_index_opt = Some(self.new_thread()?);
        }

        let thread_index = thread_index_opt.unwrap();
        let mut_thread = &mut self.thread_p[thread_index];
        mut_thread.call_push(Box::new(f))?;
        Ok(())
    }

    fn used_count(&self) -> usize {
        let mut count = 0;
        
        for i in 0..self.thread_p.len() {
            let state = self.thread_p[i].get_state();
            if state == ThreadState::Running {
                count += 1;
            }
        }
        count
    }

    fn full_count(&self) -> usize {
        self.max
    }
}