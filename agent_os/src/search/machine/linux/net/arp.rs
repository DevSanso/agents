use std::io;
use std::fs;

use crate::utils::option::opt_cast_to_io_result;
use crate::utils::result::result_cast_to_io_result;
use crate::protos::net::{ArpInfo,ArpInfos};


pub fn read_arp_info() -> io::Result<ArpInfos> {
    const PATH : &str = "/proc/net/arp";
    let file = fs::read_to_string(PATH)?;
    let data = file
                                                        .split("\n")
                                                        .skip(1)
                                                        .filter(|x| x.len() > 0);

    let arp_vec : io::Result<Vec<ArpInfo>> = data.map(|line| {
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
        let mut arp_info = ArpInfo::new();
        arp_info.ip_address = ip_address;
        arp_info.flags = flags;
        arp_info.device = device;
        arp_info.hw_type = hw_type;
        Ok(arp_info)
    }).collect();

    let mut ret = ArpInfos::new();

    ret.infos.clone_from(&arp_vec?);
    Ok(ret)    

}
