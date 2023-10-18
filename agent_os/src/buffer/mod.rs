mod double_buffer;

use std::io;

pub trait Buffer<T : Clone> {
    fn swtich(&mut self) -> io::Result<()>;
    fn add(&mut self, data : T) -> io::Result<()>;
    fn read(&mut self) -> io::Result<Vec<T>> ;
}

pub enum BufferKind {
    DoubleBuffer
}

pub fn new_buffer<T : Clone + 'static >(k : BufferKind) -> Box<dyn Buffer<T> + Send + Sync> {
    match k {
        BufferKind::DoubleBuffer => Box::new(double_buffer::DoubleBuffer::new())
    }
}