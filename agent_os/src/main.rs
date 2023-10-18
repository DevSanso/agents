pub mod search;
pub mod ipc;
pub mod buffer;
pub mod task;
pub mod config;
pub mod utils;

use std::env;
use std::error::Error;

fn main() -> Result<(),Box<dyn Error>> {
    //let args = env::args();

    //let config = config::read_config(args[1])?;

    //let ipc_listener = ipc::new_listener(ipc::ListenerKind::Mmap(config.ipc_path, 1))?;
    

    //setting tasks

    //run signal handler

    //run tasks

    //wait

    Ok(())
}
