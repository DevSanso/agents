use std::thread::{self, JoinHandle};
use std::time::{SystemTime};
use std::sync::Arc;
use std::sync::Mutex;

use bson::Bson;

use crate::buffer::BufferAdder;
use crate::search::machine::get_machine;
use crate::utils::result::result_change_err_is_string;

pub fn  total_net_stat_thread_gen( 
    buf : Arc<Mutex<dyn BufferAdder<(String,Bson)> + Sync + Send>>
)  -> JoinHandle<Result<(),String>>{

    let job = thread::spawn(move || -> Result<(), String> {
        const INTERVAL_SEC : u128 = 2000;
        let search = result_change_err_is_string(get_machine())?;
        let addr = buf;

        loop {

            let now = result_change_err_is_string(
                SystemTime::now().duration_since(SystemTime::UNIX_EPOCH))?;

            if now.as_millis() % INTERVAL_SEC >= 10 {continue;}

            {
                let data = result_change_err_is_string(search.os_net_stat())?;
                
                let doc = bson::doc! {
                    "tx" : data.total_tx() as i64,
                    "rx" : data.total_rx() as  i64,
                    "use_sock" : data.total_use_sock() as i64
                };

                let mut g = result_change_err_is_string(
                    addr.lock())?;

                result_change_err_is_string(
                    g.add(
                        (String::from("network_total"), Bson::Document(doc))
                ))?;
            }
        }

        Ok(())
    });

    job
}

#[cfg(target_os =  "linux")]
pub fn  os_details_net_stat_thread_gen( 
    buf : Arc<dyn BufferAdder<(String,Bson)> + Sync + Send>
)  -> JoinHandle<Result<(),String>>{

    let mut thread_buf = buf;

    let job = thread::spawn(move || -> Result<(), String> {
        const INTERVAL_SEC : u128 = 4000;
        let search = result_change_err_is_string(get_machine())?;
        let addr = Arc::get_mut(&mut thread_buf).unwrap();

        loop {

            let now = result_change_err_is_string(
                SystemTime::now().duration_since(SystemTime::UNIX_EPOCH))?;

            if now.as_millis() % INTERVAL_SEC >= 10 {continue;}

            {
                let data = result_change_err_is_string(search.os_net_stat())?;
                let details_data = data.extend();

                let args = result_change_err_is_string(
                    bson::to_bson(details_data.get_args().as_slice())
                )?;
                let devs =  result_change_err_is_string(
                    bson::to_bson(details_data.get_devs().as_slice())
                )?;
                let sock_stat =  result_change_err_is_string(
                    bson::to_bson(&details_data.get_sock_stat())
                )?;
                let tcps =  result_change_err_is_string(
                    bson::to_bson(details_data.get_tcps().as_slice())
                )?;

                let doc = bson::doc! {
                    "args" : args,
                    "devs" : devs,
                    "sock_stat" : sock_stat,
                    "tcps" : tcps
                };

                result_change_err_is_string(
                    addr.add(
                        (String::from("network_os_linux"), Bson::Document(doc))
                ))?;
            }
        }

        Ok(())
    });

    job
}