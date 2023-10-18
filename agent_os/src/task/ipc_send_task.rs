use std::thread;
use std::time::{SystemTime, UNIX_EPOCH};
use std::error::Error;

use bson::Bson;
use zstd;

use crate::task::Task;
use crate::buffer::Buffer;
use crate::ipc::{IpcListener};

const INTERVAL_SEC : u64 = 2;


pub struct IpcSendTask<'a, T : Buffer<(String,Bson)> + Send + Sync + ?Sized > {
    ipc :&'a dyn IpcListener,
    read_buf : &'a T,

    job : Option<thread::JoinHandle<()>>
}

impl<'a, T : Buffer<(String,Bson)> + Send + Sync + ?Sized > IpcSendTask<'a, T> {


    fn get_signal(&mut self) {}

    fn send_ipc(&mut self) {}

    fn control_task(&mut self) {}

    pub fn new(ipc :&'a dyn IpcListener,  read_buf : &'a T) -> Self {
        IpcSendTask {
            ipc,
            read_buf,
            job : None
        }
    }
}

impl<'a, T : Buffer<(String,Bson)> + Send + Sync + ?Sized > Task<'a> for IpcSendTask<'a, T> {
    fn start(&mut self) {
        
    }
    fn stop(&mut self) {
        
    }
    fn reflesh(&mut self) {
        
    }
    fn kill(&mut self) {
        
    }
}

pub fn new_ipc_send_task<'a,  T : Buffer<(String,Bson)> + Send + Sync + ?Sized  >(ipc :&'a dyn IpcListener,  read_buf : &'a  T) -> impl Task<'a> {
    IpcSendTask::<T>::new(ipc, read_buf)
}