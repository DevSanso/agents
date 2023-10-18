use std::sync::Mutex;

use crate::buffer::Buffer;
use crate::utils::result::result_cast_to_io_result;

pub struct DoubleBuffer<T> {
    w_index : usize,
    r_index : usize,

    slices : [Vec<T>;2],
    m : Mutex<()>
}

unsafe impl<T> Send for DoubleBuffer<T> {}
unsafe impl<T> Sync for DoubleBuffer<T> {}

impl<T : Clone> DoubleBuffer<T> {
    pub fn new() -> Self {
        DoubleBuffer { w_index: 0, r_index: 1, slices: [Vec::new(), Vec::new()], m: Mutex::new(()) }
    }
}

impl<T : Clone> Buffer<T> for DoubleBuffer<T> {
    fn swtich(&mut self) -> std::io::Result<()> {
        let g = result_cast_to_io_result(self.m.lock())?;
        let temp = self.w_index;
        self.w_index = self.r_index;
        self.r_index = temp;
        drop(g);

        Ok(())
    }

    fn add(&mut self, data : T) -> std::io::Result<()> {
        let g = result_cast_to_io_result(self.m.lock())?;
        let w = &mut self.slices[self.w_index];
        w.push(data);
        drop(g);

        Ok(())
    }
    
    fn read(&mut self) -> std::io::Result<Vec<T>> {
        let g = result_cast_to_io_result(self.m.lock())?;
        let r = self.slices[self.r_index].clone();
        drop(g);

        Ok(r)
    }
}