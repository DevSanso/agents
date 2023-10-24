pub mod thraed_pool;

use std::io;
use std::time;

pub trait Pool<T : Send + Sync,R,E> {
    fn use_item<F : 'static + FnOnce(T) -> Result<(),String> + Send >(&mut self, object : T,  f : F) -> io::Result<()>;
    fn used_count(&self) -> usize;
    fn full_count(&self) -> usize;
}