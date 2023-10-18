use std::sync::{Mutex,Arc};
use std::fs::OpenOptions;
use std::io;
use std::path::Path;

use memmap2;

use crate::ipc::make_format;
use crate::utils::result::result_cast_to_io_result;
use crate::ipc::{IpcListener, IpcSendStream};
use crate::utils::seq::{new_seq, Sequence, SequenceKind};

pub struct MmapListener {
    writer : Arc<Mutex<(memmap2::MmapMut, Box<dyn Sequence>)>>,
    file_size : usize
}

impl MmapListener {
    pub(crate) fn new<P : AsRef<Path>>(path : P, file_size : u64) -> io::Result<Self> {
        let file_res = OpenOptions::new()
        .read(true)
        .write(true)
        .create(true)
        .open(&path);
        
        let file = result_cast_to_io_result(file_res)?;
        file.set_len(file_size)?;

        let mmap = unsafe { memmap2::MmapMut::map_mut(&file)? };
        let seq = new_seq(SequenceKind::U64(0));
        
        let field: Arc<Mutex<(memmap2::MmapMut, Box<dyn Sequence>)>> = Arc::new(Mutex::new((mmap,seq)));
        Ok(MmapListener { writer: field , file_size : file_size as usize})
    }
}

impl IpcListener for MmapListener {
    fn get_stream(&mut self) -> std::io::Result<Box<dyn super::IpcSendStream + Send>> {
        let clone_arc = Arc::clone(&self.writer);

        Ok({
            Box::new(MmapSendStream::new(clone_arc, self.file_size))
        })
    }
}

pub struct MmapSendStream {
    f : Arc<Mutex<(memmap2::MmapMut, Box<dyn Sequence>)>>,
    file_size: usize
}

unsafe impl Send for MmapSendStream {}

impl MmapSendStream {
    pub(crate) fn new(arc :  Arc<Mutex<(memmap2::MmapMut, Box<dyn Sequence>)>>, size : usize) -> Self {
        MmapSendStream { f:arc, file_size : size }
    }
}

impl IpcSendStream for MmapSendStream {
    fn send(&mut self, data : &'_ [u8]) -> std::io::Result<()> {
        let mut g = result_cast_to_io_result(self.f.lock())?;
        
        let origin = make_format(g.1.next(), data);
        let msg : Vec<u8> = (|| {
            let mut ext = origin.clone();
            ext.resize(self.file_size, 0);
            ext
        })();
        g.0.copy_from_slice(msg.as_slice());
        Ok(())
    }
}