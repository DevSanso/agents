pub mod config;
pub mod ipc;
pub mod pool;
pub mod search;
pub mod task;
pub mod utils;
pub mod protos;

use std::env;
use std::error::Error;
use std::io;
use std::sync::Arc;
use std::time;

use pool::Pool;
use utils::buffer;

const OS_DETAILS_NET_INTERVAL: u64 = 6;
const IPC_SEND_INTERVAL: u64 = 2;

fn main() -> Result<(), Box<dyn Error>> {
    let config_path = env::args().skip(1).next();

    if config_path == None {
        return Err(Box::new(io::Error::new(
            io::ErrorKind::InvalidData,
            "process args not exists",
        )));
    }

    let config = config::read_config(config_path.unwrap())?;

    let mut ipc_listener =
        ipc::new_listener(ipc::ListenerKind::Mmap(config.ipc_path, config.ipc_size))?;

    let tp = pool::thraed_pool::ThreadPool::new(200);

    let buf = buffer::DoubleBuffer::<protos::message::Data>::new();

    loop {
        {
            let now = utils::util_time::get_unix_epoch_now();

            let tp_g_res = tp.lock();

            if tp_g_res.is_err() {
                return Err(Box::new(io::Error::new(
                    io::ErrorKind::Interrupted,
                    "thread pool mutex lock is failed",
                )));
            }

            let mut tp_g = tp_g_res.unwrap();
            
            if utils::util_time::is_interval(now, time::Duration::from_secs(IPC_SEND_INTERVAL)) {
                let buf_clone = Arc::clone(&buf);
                let fun = task::ipc_send_task_gen(ipc_listener.get_stream()?, buf_clone);

                tp_g.use_item((), fun)?;
            }
            
            if utils::util_time::is_interval(now, time::Duration::from_secs(OS_DETAILS_NET_INTERVAL)) {
                let buf_clone = Arc::clone(&buf);
                let fun = task::os_details_net_stat_thread_gen(buf_clone);
                tp_g.use_item((), fun)?;
            }
        }
    }

    //Ok(())
}
