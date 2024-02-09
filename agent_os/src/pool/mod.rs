pub mod thread_pool;
use std::io;

pub trait Pool<T : Send + Sync,R,E> {
    fn run_func<F : FnOnce(T) -> Result<R,E> + Send + 'static>(&mut self, arg : T, f : F) -> io::Result<()>;
    fn used_count(&self) -> usize;
    fn full_count(&self) -> usize;
}

pub use thread_pool::ThreadPool;