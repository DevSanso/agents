use std::path::Path;
use std::io;
use std::fs::read_to_string;

use serde::Deserialize;
use toml;

use crate::utils::result::result_cast_to_io_result;

#[derive(Deserialize)]
pub struct Config {
    pub ipc_path : String
}

pub fn read_config<P : AsRef<Path>>(path : P) -> io::Result<Config> {
    let cfg = read_to_string(path)?;
    let c : Config = result_cast_to_io_result(toml::from_str(cfg.as_str()))?;

    Ok(c)
}