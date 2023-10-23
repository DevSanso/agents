use std::io;
use std::iter::Iterator;

pub mod mmap;

pub trait IpcListener {
    fn get_stream(&mut self) -> io::Result<Box<dyn IpcSendStream + Send + Sync>>;
}


pub trait IpcSendStream {
    fn send(&mut self, data : &'_ [u8]) -> io::Result<()>;
}
#[inline]
pub fn make_format(seq : Box<dyn Iterator<Item = u8>>, data : &'_ [u8] )-> Vec<u8> {
    let mut ret = Vec::from(data);
    ret.extend_from_slice("|".as_bytes());
    ret.extend(seq);
    
    ret
}

pub enum ListenerKind {
    Mmap(String,u64)
}

pub fn new_listener(kind : ListenerKind) -> io::Result<Box<dyn IpcListener  + Send + Sync>> {
    let l = match kind {
        ListenerKind::Mmap(p, size) => mmap::MmapListener::new(p, size)
    }?;

    Ok(Box::new(l))
}
