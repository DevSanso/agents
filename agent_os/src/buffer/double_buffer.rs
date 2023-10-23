use std::sync::Arc;
use std::sync::Mutex;
use std::sync::atomic::AtomicUsize;
use std::sync::atomic::Ordering::Relaxed;

use crate::buffer;

use crate::utils::result::result_cast_to_io_result;

pub struct DoubleBuffer<T> {
    w_index : AtomicUsize,
    r_index : AtomicUsize,

    slices : [Vec<T>;2],
}

impl <T: Clone> buffer::BufferAdder<T> for DoubleBuffer<T>{
    fn add(&mut self, data : T) -> std::io::Result<()> {

        self.slices[self.w_index .load(Relaxed)].push(data);
        Ok(())

    }
}
impl <T:Clone> buffer::BufferControllerAndReader<T> for DoubleBuffer<T> {
    fn read(&self) -> std::io::Result<Vec<T>> {
        Ok(self.slices[self.r_index.load(Relaxed)].clone())
    }

    fn swtich(&mut self) -> std::io::Result<()> {
        let w = self.w_index.load(Relaxed);
        self.w_index.swap(self.r_index.swap(w, Relaxed), Relaxed);
        Ok(())
    }
}


impl <T: Clone> buffer::BufferReader<T> for DoubleBuffer<T> {
    fn read(&self) -> std::io::Result<Vec<T>> {
        Ok(self.slices[self.r_index.load(Relaxed)].clone())
    }
}

impl <T: Clone> buffer::BufferController<T> for DoubleBuffer<T> {
    fn swtich(&mut self) -> std::io::Result<()> {
        let w = self.w_index.load(Relaxed);
        self.w_index.swap(self.r_index.swap(w, Relaxed), Relaxed);
        Ok(())
    }
}

unsafe impl<T : Clone> Send for DoubleBuffer<T> {}
unsafe impl<T : Clone> Sync for DoubleBuffer<T> {}

impl<T : Clone > DoubleBuffer<T > {
    pub fn new() -> Arc<Mutex<DoubleBuffer< T >>>{
        let o = DoubleBuffer::<T> {
            w_index: AtomicUsize::new(0), r_index: AtomicUsize::new(1), slices: [Vec::new(), Vec::new()]
        }; 
        Arc::new(Mutex::new(o))
    }

}
