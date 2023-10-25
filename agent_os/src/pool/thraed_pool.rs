use std::thread::{self, JoinHandle};
use std::io;
use std::sync::{Mutex, Arc};
use crate::pool;

type ThreadItem = Option<JoinHandle<Result<(),String>>>;

#[derive(Debug)]
pub struct ThreadPool {
    max : usize,
    thread_p : Vec<ThreadItem>
}

impl ThreadPool {
    fn clear(&mut self) {

        for i in 0..self.thread_p.len() {
            {
                let buf = self.thread_p[i].as_ref();

                if buf.is_none() { continue;}
    
                let t= buf.unwrap();
                if t.is_finished() == false {continue;}
            }

            let ele = self.thread_p.get_mut(i).unwrap();
            let j = ele.take().unwrap();
            let _ = j.join();
            
            self.thread_p[i] = None;
        }
    }
    pub fn new(size : usize) -> Arc<Mutex<Self>> {
        let mut tp = ThreadPool {max : size, thread_p : Vec::with_capacity(size)};
        unsafe {tp.thread_p.set_len(size);}

        Arc::new(Mutex::new(tp))
    } 
}

impl pool::Pool<(), () , String> for ThreadPool {
    fn use_item<F : 'static +  FnOnce(()) -> Result<(),String> + Send>(&mut self, object : (),  f :  F) -> io::Result<()> {
        let index_opt = self.thread_p.iter().position(|x| x.is_none());

        if index_opt.is_none() {
            return Err(io::Error::new(io::ErrorKind::UnexpectedEof, "pool is full"));
        }
        
        let handle = thread::spawn(move || {
            let ret = f(object);
            ret
        });

        let index = index_opt.unwrap();
        self.thread_p[index] = Some(handle);

        if (self.used_count() as f64 / self.full_count() as f64) * 100.0 > 20.0 {
            self.clear();
        }

        Ok(())
    }

    fn used_count(&self) -> usize {
        self.thread_p.iter()
            .filter(|&x| {x.is_some()})
            .count()
    }

    fn full_count(&self) -> usize {
        self.max
    }
}