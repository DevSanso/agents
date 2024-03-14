pub mod config;
pub mod ipc;
pub mod search;
pub mod task;
pub mod utils;
pub mod protos;
pub mod structure;

use std::{env, thread};
use std::error::Error;
use std::io;
use std::sync::Arc;
use std::time;

use simplelog::*;

use structure::pool;
use structure::pool::Pool;
use structure::buffer;

fn convert_string_to_filter(level: &String) -> LevelFilter {
    match level.as_str() {
        "trace" => LevelFilter::Trace,
        "debug" => LevelFilter::Debug,
        "info" => LevelFilter::Info,
        "warn" => LevelFilter::Warn,
        "error" => LevelFilter::Error,
        _ => LevelFilter::Error
    }
}

fn init_logger(config : &config::Config) -> Result<(), log::SetLoggerError>{

    let loggers = (||{
        let mut ret :Vec<Box<dyn SharedLogger + 'static>> = vec![];

        ret.push(SimpleLogger::new(LevelFilter::Trace, Config::default()));
        
        if config.log_path.is_some() {
            let level = config.log_level.as_ref().unwrap();
            ret.push(WriteLogger::new(
                convert_string_to_filter(level),
                Config::default(),
                std::fs::File::create(config.log_path.as_ref().unwrap()).unwrap(),
            ));
        }

        ret
    })();
    
    CombinedLogger::init(loggers)
}

fn main() -> Result<(), Box<dyn Error>> {
    let config_path = env::args().skip(1).next();

    if config_path == None {
        return Err(Box::new(io::Error::new(
            io::ErrorKind::InvalidData,
            "process args not exists",
        )));
    }

    let config = config::read_config(config_path.unwrap())?;

    init_logger(&config)?;

    let mut ipc_listener =
        ipc::new_listener(ipc::ListenerKind::Mmap(config.ipc_path, config.ipc_size))?;

    let tp = pool::ThreadPool::new(30,200);

    let buf = buffer::DoubleBuffer::<protos::os_snap::Data>::new();

    {
        let mut tp_g = tp.lock().unwrap();
        
        let buf_ipc = Arc::clone(&buf);
        let buf_stat = Arc::clone(&buf);

        let ipc = task::ipc_send_task_gen(ipc_listener.get_stream()?, buf_ipc);
        let net_stat = task::os_details_net_stat_thread_gen(buf_stat);
        
        tp_g.run_func((), ipc)?;
        tp_g.run_func((), net_stat)?;
    }

    loop {
        thread::sleep(time::Duration::from_secs(1));
    }

    //Ok(())
}
