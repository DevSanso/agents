use std::io;
use std::fs;

use crate::utils::convert_to_io_result;
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
        let ip_address =  convert_to_io_result!(option ,tok.next(),"arp ip_address is null")?.to_string();

        let hw_type_str : &'_ str = convert_to_io_result!(option, tok.next(),"arp hw_type is null")?;
        let hw_type = convert_to_io_result!(result, 
            u32::from_str_radix(hw_type_str 
                .replace("0x", "")
                .as_str(),
            16)
        )?;

        let flags = convert_to_io_result!(option, tok.next(),"arp flags is null")?.to_string();
        tok.next();
        tok.next();
        let device = convert_to_io_result!(option, tok.next(),"arp device is null")?.to_string();
        
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
