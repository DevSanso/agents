pub mod buffer;
pub mod config;
pub mod ipc;
pub mod pool;
pub mod search;
pub mod task;
pub mod utils;



use std::env;
use std::error::Error;
use std::io;
use std::sync::Arc;
use std::time;

use bson::Bson;

use pool::Pool;

const TOTAL_NET_INTERVAL: f64 = 4.0;
const OS_DETAILS_NET_INTERVAL: f64 = 6.0;
const IPC_SEND_INTERVAL: f64 = 2.0;

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

    let buf = buffer::DoubleBuffer::<(String, Bson)>::new();

    loop {
        {
            let now = time::SystemTime::now().duration_since(time::UNIX_EPOCH)?;

            let now_sec = now.as_secs();

            let tp_g_res = tp.lock();

            if tp_g_res.is_err() {
                return Err(Box::new(io::Error::new(
                    io::ErrorKind::Interrupted,
                    "thread pool mutex lock is failed",
                )));
            }

            let mut tp_g = tp_g_res.unwrap();

            if now_sec as f64 % IPC_SEND_INTERVAL <= 0.005 {
                let buf_clone = Arc::clone(&buf);
                let fun = task::ipc_send_task_gen(ipc_listener.get_stream()?, buf_clone);

                tp_g.use_item((), fun)?;
            }

            if now_sec as f64 % TOTAL_NET_INTERVAL <= 0.005 {
                let buf_clone = Arc::clone(&buf);
                let fun = task::total_net_stat_thread_gen(buf_clone);

                tp_g.use_item((), fun)?;
            }

            if now_sec as f64 % OS_DETAILS_NET_INTERVAL <= 0.005 {
                let buf_clone = Arc::clone(&buf);
                let fun = task::os_details_net_stat_thread_gen(buf_clone);
                tp_g.use_item((), fun)?;
            }
        }
    }

    Ok(())
}
