use std::thread::{self, JoinHandle};
use std::time::{SystemTime};
use std::sync::Arc;
use std::sync::Mutex;

use bson::Bson;
use zstd;

use crate::buffer::{BufferControllerAndReader};
use crate::ipc::{IpcListener};
use crate::utils::result::result_change_err_is_string;

pub fn ipc_send_thread_gen(
    mut listen : Box<dyn IpcListener + Sync + Send>, 
    buf : Arc<Mutex<dyn BufferControllerAndReader<(String,Bson)>  + Send + Sync  >>
)  -> JoinHandle<Result<(),String>>{

    let job = thread::spawn(move || -> Result<(), String> {
        const INTERVAL_SEC : u128 = 2000;

        let controller = buf;

        let mut stream= result_change_err_is_string(
            listen.get_stream())?;
            
        loop {

            let now = result_change_err_is_string(
                SystemTime::now().duration_since(SystemTime::UNIX_EPOCH))?;

            if now.as_millis() % INTERVAL_SEC >= 10 {continue;}

            {
                let mut g = result_change_err_is_string(
                    controller.lock())?;
    
                result_change_err_is_string(g.swtich())?;

                let send_data = result_change_err_is_string(g.read())?;
            
                let collected = send_data.into_iter().fold(bson::Document::new(), |mut acc, x| {
                    acc.insert(x.0,  x.1);
                    acc
                });


                let mut pack_data = Vec::<u8>::new();
                
                result_change_err_is_string(collected.to_writer(&mut pack_data))?;
            
                let compress_data = result_change_err_is_string(
                    zstd::encode_all(pack_data.as_slice(), 19))?;
                
                result_change_err_is_string(stream.send(compress_data.as_slice()))?;
            }
        }

        Ok(())
    });

    job
}