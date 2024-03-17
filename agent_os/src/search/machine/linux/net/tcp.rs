use std::io;
use std::fs;

use crate::utils::convert_to_io_result;
use crate::protos::net::{Tcp4Stat, Tcp4Stats};

#[inline]
fn split_colon_ret_two_str<'a>(s : &&'a str) -> (&'a str, &'a str) {
    let mut s = s.split(":").take(2);

    (s.next().unwrap(),  s.next().unwrap())
}

pub fn read_tcp_stats() -> io::Result<Tcp4Stats> {
    const PATH : &str = "/proc/net/tcp";
    let lines = fs::read_to_string(PATH)?;
    let datas = lines.split("\n").skip(1);

    let mut v = Tcp4Stats::new();
    
    for origin in datas {
        let mut data = origin.split_whitespace();
        
        if data.next()== None {continue;}

        let mut tok = data;

        let local_addr = tok.next().expect("local address is null");
        let local_addr_hex = String::from(local_addr);

        let remote_addr =tok.next().expect("remote address is null");
        let remote_addr_hex = String::from(remote_addr);

        let connection_state_opt = convert_to_io_result!(option, tok.next(),"connection_state is null")?;
        let connection_state = convert_to_io_result!(result, u32::from_str_radix(connection_state_opt,16))?;
        
        let q = tok.next().expect("queue is null");
        let tup = split_colon_ret_two_str(&q);
        
        let tx_queue = convert_to_io_result!(result, u64::from_str_radix(tup.0, 16))?;//convert_str_to_u64(tup.0, 16)?;
        let rx_queue =convert_to_io_result!(result, u64::from_str_radix(tup.1, 16))?;// convert_str_to_u64(tup.1, 16)?;

        let tr_tm_when = tok.next().expect("tr tm when is null");
        let tup2 = split_colon_ret_two_str(&tr_tm_when);

        let time_active = convert_to_io_result!(result, u32::from_str_radix(tup2.0, 16))?;
        let jiffies_timer_expires = convert_to_io_result!(result, u64::from_str_radix(tup2.1, 16))?;

        let rto_opt = convert_to_io_result!(option, tok.next(),"rto is null")?;
        let rto =  convert_to_io_result!(result, u64::from_str_radix(rto_opt,16))?;

        let uid_opt = convert_to_io_result!(option, tok.next(),"uid is null")?;
        let uid =  convert_to_io_result!(result, u32::from_str_radix(uid_opt,10))?;

        let zero_window_probes_opt = convert_to_io_result!(option, tok.next(),"zero window probes is null")?;
        let zero_window_probes =   convert_to_io_result!(result, u32::from_str_radix(zero_window_probes_opt, 16))?;

        let inode_opt = convert_to_io_result!(option, tok.next(),"inode is null")?;
        let inode = convert_to_io_result!(result, u64::from_str_radix(inode_opt,10))?;

        let etc = tok.fold(String::new(), |mut acc: String, x| {
            acc.push_str(" ");
            acc.push_str(x);
            acc
        });
        let mut tcpstat = Tcp4Stat::new();
        tcpstat.local_addr_hex = local_addr_hex;
        tcpstat.remote_addr_hex = remote_addr_hex;
        tcpstat.connection_state = connection_state;
        tcpstat.tx_queue = tx_queue;
        tcpstat.rx_queue = rx_queue;
        tcpstat.time_active = time_active;
        tcpstat.jiffies_timer_expires = jiffies_timer_expires;
        tcpstat.rto = rto;
        tcpstat.uid = uid;
        tcpstat.zero_window_probes = zero_window_probes;
        tcpstat.inode = inode;
        tcpstat.etc = etc;

        v.stats.push(tcpstat);
    }
    Ok(v)
}
