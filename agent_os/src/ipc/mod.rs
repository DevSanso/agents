use std::io;
use std::iter::Iterator;

pub mod mmap;

pub trait IpcSequence {
    fn current(&self) -> Box<dyn Iterator<Item = u8>>;
    fn next(&mut self) ->Box<dyn Iterator<Item = u8>>;
}
pub trait IpcListener {
    fn get_stream(&mut self) -> io::Result<Box<dyn IpcSendStream>>;
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

pub fn new_listener(kind : ListenerKind) -> io::Result<Box<dyn IpcListener>> {
    let l = match kind {
        ListenerKind::Mmap(p, size) => mmap::MmapListener::new(p, size)
    }?;

        Ok(Box::new(l))
}
pub enum SequenceKind {
    Mmap(u64)
}

pub fn new_seq(kind : SequenceKind) -> Box<dyn IpcSequence> {
    match kind {
        SequenceKind::Mmap(start) => Box::new(mmap::MmapSequence::new(start))
    }
}