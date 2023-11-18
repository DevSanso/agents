use std::io;
use std::fs;

use crate::utils::result::result_cast_to_io_result;
use crate::protos::net::{NetDevInfo,NetDevInfos};

pub fn read_net_dev_info() -> io::Result<NetDevInfos> {
    const PATH: &str = "/proc/net/dev";
    let file = fs::read_to_string(PATH)?;
    let data = file.split("\n")
                                                                .skip(2)
                                                                .filter(|x| x.len() >= 17);
    
    let dev_slice : io::Result<Vec<NetDevInfo>> = data.map(|x| {
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
        let mut devi = NetDevInfo::new();

        devi.interface = interface;
        devi.rx_bytes = rx_bytes;
        devi.rx_packets = rx_packets;
        devi.rx_errs = rx_errs;
        devi.rx_drop = rx_drop;
        devi.rx_fifo = rx_fifo;
        devi.rx_frame = rx_frame;
        devi.rx_compressed = rx_compressed;
        devi.tx_bytes = tx_bytes;
        devi.tx_packets = tx_packets;
        devi.tx_errs = tx_errs;
        devi.tx_drop = tx_drop;
        devi.tx_fifo = tx_fifo;
        devi.tx_frame = tx_frame;
        devi.tx_compressed = tx_compressed;

        Ok(devi)

    }).collect();
    let mut ret = NetDevInfos::new();
    ret.infos.clone_from(&dev_slice?);
    Ok(ret)
}