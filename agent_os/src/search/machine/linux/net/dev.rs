use std::io;
use std::fs;

use crate::utils::result::result_cast_to_io_result;

pub struct NetDevInfo {
    pub interface: String,
    pub rx_bytes: u64,
    pub rx_packets: u64,
    pub rx_errs: u64,
    pub rx_drop: u64,
    pub rx_fifo: u64,
    pub rx_frame: u64,
    pub rx_compressed: u64,
    pub tx_bytes: u64,
    pub tx_packets: u64,
    pub tx_errs: u64,
    pub tx_drop: u64,
    pub tx_fifo: u64,
    pub tx_frame: u64,
    pub tx_compressed: u64,
}

impl Default for NetDevInfo {
    fn default() -> Self {
        NetDevInfo {
            interface: String::new(),
            rx_bytes: 0,
            rx_packets: 0,
            rx_errs: 0,
            rx_drop: 0,
            rx_fifo: 0,
            rx_frame: 0,
            rx_compressed: 0,
            tx_bytes: 0,
            tx_packets: 0,
            tx_errs: 0,
            tx_drop: 0,
            tx_fifo: 0,
            tx_frame: 0,
            tx_compressed: 0,
        }
    }
}

pub fn read_net_dev_info() -> io::Result<Vec<NetDevInfo>> {
    const PATH: &str = "/proc/net/dev";
    let file = fs::read_to_string(PATH)?;
    let data = file.split("\n")
                                                                .skip(2)
                                                                .filter(|x| x.len() >= 17);
    
    let ret : io::Result<Vec<NetDevInfo>> = data.map(|x| {
        let tok : Vec<&str> = x.split_whitespace().collect();
        
        let interface = tok[0].trim_matches(':').to_string();

        let rx_bytes = result_cast_to_io_result(
            tok[1].parse::<u64>()
        )?;

        let rx_packets =  result_cast_to_io_result(
            tok[2].parse::<u64>()
        )?;

        let rx_errs =  result_cast_to_io_result(
            tok[3].parse::<u64>()
        )?;

        let rx_drop =  result_cast_to_io_result(
            tok[4].parse::<u64>()
        )?;

        let rx_fifo =  result_cast_to_io_result(
            tok[5].parse::<u64>()
        )?;

        let rx_frame =  result_cast_to_io_result(
            tok[6].parse::<u64>()
        )?;

        let rx_compressed =  result_cast_to_io_result(
            tok[7].parse::<u64>()
        )?;

        let tx_bytes = result_cast_to_io_result(
            tok[9].parse::<u64>()
        )?;

        let tx_packets =  result_cast_to_io_result(
            tok[10].parse::<u64>()
        )?;

        let tx_errs =  result_cast_to_io_result(
            tok[11].parse::<u64>()
        )?;

        let tx_drop =  result_cast_to_io_result(
            tok[12].parse::<u64>()
        )?;

        let tx_fifo =  result_cast_to_io_result(
            tok[13].parse::<u64>()
        )?;

        let tx_frame =  result_cast_to_io_result(
            tok[14].parse::<u64>()
        )?;

        let tx_compressed =  result_cast_to_io_result(
            tok[15].parse::<u64>()
        )?;

        Ok(NetDevInfo {
            interface,
            rx_bytes,
            rx_packets,
            rx_errs,
            rx_drop,
            rx_fifo,
            rx_frame,
            rx_compressed,
            tx_bytes,
            tx_packets,
            tx_errs,
            tx_drop,
            tx_fifo,
            tx_frame,
            tx_compressed,
        })

    }).collect();

    ret
}