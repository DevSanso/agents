use std::io;
use std::fs;

use crate::utils::option::opt_cast_to_io_result;
use crate::utils::result::result_cast_to_io_result;
use crate::protos::net::SockStatInfo;


pub fn read_sock_stat_info() -> io::Result<SockStatInfo> {
    const PATH: &str = "/proc/net/sockstat";
    let file = fs::read_to_string(PATH)?;
    let mut data = file
                                                        .split("\n")
                                                        .filter(|x| x.len() > 0);

    let use_socket_line = opt_cast_to_io_result(data.next(),"use socket line not exists")?;
    let tcp_line = opt_cast_to_io_result(data.next(),"tcp socket line not exists")?;
    
    let mut use_sock_tok = use_socket_line.split_whitespace().skip(2);
    let use_count = result_cast_to_io_result(
        u64::from_str_radix(
            opt_cast_to_io_result(use_sock_tok.next(),"use socket num not exists")?, 10))?;
    
    let tcp_line_tok : Vec<&str> = tcp_line.split_whitespace().collect();
    let in_use =result_cast_to_io_result(
        u64::from_str_radix(tcp_line_tok[2], 10)
    )?;
    let orphan = result_cast_to_io_result(
        u64::from_str_radix(tcp_line_tok[4], 10)
    )?;
    let tw = result_cast_to_io_result(
        u64::from_str_radix(tcp_line_tok[6], 10)
    )?;
    let alloc = result_cast_to_io_result(
        u64::from_str_radix(tcp_line_tok[8], 10)
    )?;
    let mem_kb = result_cast_to_io_result(
        u64::from_str_radix(tcp_line_tok[10], 10)
    )? as f64 / 1024.0;

    let mut stat = SockStatInfo::new();
    stat.use_count = use_count;
    stat.in_use = in_use;
    stat.orphan = orphan;
    stat.tw = tw;
    stat.alloc = alloc;
    stat.mem_kb = mem_kb;

    Ok(stat)
}