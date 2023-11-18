mod double_buffer;

use std::io;

pub  trait BufferController<T : Clone> {
    fn swtich(&mut self) -> io::Result<()>;
}

pub trait BufferAdder<T : Clone> {
    fn add(&mut self, data : T) -> io::Result<()>;
}


pub  trait BufferControllerAndReader<T : Clone> {
    fn swtich(&mut self) -> io::Result<()>;
    fn read(&self) -> io::Result<Vec<T>> ;
}

pub trait BufferReader<T : Clone> {
    fn read(&self) -> io::Result<Vec<T>> ;
}

pub use double_buffer::DoubleBuffer;