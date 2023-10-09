use std::sync::{Mutex,Arc};
use std::fs::OpenOptions;
use std::io;
use std::path::Path;

use memmap2;

use crate::ipc::make_format;
use crate::utils::result::result_cast_to_io_result;
use crate::ipc::{IpcSequence, IpcListener, IpcSendStream};

pub struct MmapSequence {
    current_val : Mutex<u64>
}

impl MmapSequence {
    pub(crate) fn new(start : u64) -> Self {
        MmapSequence { current_val:Mutex::new(start) }
    }
}

impl IpcSequence for MmapSequence {
    fn current(&self) -> Box<dyn Iterator<Item = u8>> {
        let g = self.current_val.lock().unwrap();
        Box::new(g.to_le_bytes().into_iter())
    }

    fn next(&mut self) ->Box<dyn Iterator<Item = u8>> {
        let mut g = self.current_val.lock().unwrap();
        let ret = Box::new(g.to_le_bytes().into_iter());
        *g = *g + 1;
        ret
    }
}

pub struct MmapListener {
    writer : Arc<Mutex<(memmap2::MmapMut, MmapSequence)>>,
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
        let seq = MmapSequence::new(1);
        
        let field: Arc<Mutex<(memmap2::MmapMut, MmapSequence)>> = Arc::new(Mutex::new((mmap,seq)));
        Ok(MmapListener { writer: field , file_size : file_size as usize})
    }
}

impl IpcListener for MmapListener {
    fn get_stream(&mut self) -> std::io::Result<Box<dyn super::IpcSendStream>> {
        let clone_arc = Arc::clone(&self.writer);

        Ok({
            Box::new(MmapSendStream::new(clone_arc, self.file_size))
        })
    }
}

pub struct MmapSendStream {
    f : Arc<Mutex<(memmap2::MmapMut, MmapSequence)>>,
    file_size: usize
}
impl MmapSendStream {
    pub(crate) fn new(arc :  Arc<Mutex<(memmap2::MmapMut, MmapSequence)>>, size : usize) -> Self {
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