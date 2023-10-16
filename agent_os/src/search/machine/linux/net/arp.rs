use std::io;
use std::fs;

use crate::utils::option::opt_cast_to_io_result;
use crate::utils::result::result_cast_to_io_result;

#[derive(Debug,Clone)]
pub struct ArpInfo {
    pub ip_address: String,
    pub hw_type: u32,
    pub flags: String,
    //hw_address: String,
    //mask: String,
    pub device: String,
}

impl Default for ArpInfo {
    fn default() -> Self {
        ArpInfo {
            ip_address : String::new(),
            hw_type : 0,
            flags : String::new(),
            device : String::new()
        }
    }
}

pub fn read_arp_info() -> io::Result<Vec<ArpInfo>> {
    const PATH : &str = "/proc/net/arp";
    let file = fs::read_to_string(PATH)?;
    let data = file
                                                        .split("\n")
                                                        .skip(1)
                                                        .filter(|x| x.len() > 0);

    let ret : io::Result<Vec<ArpInfo>> = data.map(|line| {
        let mut tok = line.split_whitespace();

        let ip_address =  opt_cast_to_io_result(tok.next(),"arp ip_address is null")?.to_string();

        let hw_type = result_cast_to_io_result(
            u32::from_str_radix(
                opt_cast_to_io_result(tok.next(),"arp hw_type is null")?
                    .replace("0x", "")
                    .as_str(),
                 16)
        )?;

        let flags = opt_cast_to_io_result(tok.next(),"arp flags is null")?.to_string();
        tok.next();
        tok.next();
        let device = opt_cast_to_io_result(tok.next(),"arp device is null")?.to_string();

        Ok(ArpInfo {
            ip_address ,
            hw_type ,
            flags,
            device
        })
    }).collect();

    ret

}
