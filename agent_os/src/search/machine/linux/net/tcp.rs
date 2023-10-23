use std::io;
use std::fs;

use serde::Serialize;

use crate::utils::result::result_cast_to_io_result;
use crate::utils::option::opt_cast_to_io_result;

#[derive(Debug,Clone,Serialize)]
pub struct  Tcp4Stat {
    pub local_addr_hex : String,
    pub remote_addr_hex : String,
    pub connection_state : u8,
    pub tx_queue :  u64,
    pub rx_queue : u64,
    pub time_active : u8,
    pub jiffies_timer_expires : u64,
    pub rto : u64,
    pub uid : u32,
    pub zero_window_probes : u8,
    pub inode : u64,
    /* 
        - socket reference count
        - location of socket in memory
        - retransmit timeout
        - Predicted tick of soft clock
        - sending congestion window
        - slow start size threshold, 
     */
    pub etc : String
    
}

#[inline]
fn split_colon_ret_two_str<'a>(s : &&'a str) -> (&'a str, &'a str) {
    let mut s = s.split(":").take(2);

    (s.next().unwrap(),  s.next().unwrap())
}

pub fn read_tcp_stats() -> io::Result<Vec<Tcp4Stat>> {
    const PATH : &str = "/proc/net/tcp";
    let lines = fs::read_to_string(PATH)?;
    let datas = lines.split("\n").skip(1);

    let mut v = Vec::<Tcp4Stat>::new();
    
    for origin in datas {
        let mut data = origin.split_whitespace();
        
        if data.next()== None {continue;}

        let mut tok = data;

        let local_addr = tok.next().expect("local address is null");
        let local_addr_hex = String::from(local_addr);

        let remote_addr =tok.next().expect("remote address is null");
        let remote_addr_hex = String::from(remote_addr);

        let connection_state = result_cast_to_io_result(
            u8::from_str_radix(
                opt_cast_to_io_result(tok.next(),"connection_state is null")?,
                 16)
        )?;
        let q = tok.next().expect("queue is null");
        let tup = split_colon_ret_two_str(&q);
        
        let tx_queue = result_cast_to_io_result(u64::from_str_radix(tup.0, 16))?;//convert_str_to_u64(tup.0, 16)?;
        let rx_queue =result_cast_to_io_result(u64::from_str_radix(tup.1, 16))?;// convert_str_to_u64(tup.1, 16)?;

        let tr_tm_when = tok.next().expect("tr tm when is null");
        let tup2 = split_colon_ret_two_str(&tr_tm_when);

        let time_active = result_cast_to_io_result(u8::from_str_radix(tup2.0, 16))?;
        let jiffies_timer_expires = result_cast_to_io_result(u64::from_str_radix(tup2.1, 16))?;

        let rto =  result_cast_to_io_result(
            u64::from_str_radix(
                opt_cast_to_io_result(tok.next(),"rto is null")?,
                 16)
        )?;
        let uid =  result_cast_to_io_result(
            u32::from_str_radix(
                opt_cast_to_io_result(tok.next(),"uid is null")?,
                 10)
        )?;

        let zero_window_probes =   result_cast_to_io_result(
            u8::from_str_radix(
                opt_cast_to_io_result(tok.next(),"zero window probes is null")?,
                 16)
        )?;
        let inode = result_cast_to_io_result(
            u64::from_str_radix(opt_cast_to_io_result(tok.next(),"inode is null")?,10)
        )?;

        let etc = tok.fold(String::new(), |mut acc: String, x| {
            acc.push_str(" ");
            acc.push_str(x);
            acc
        });

        let ele : Tcp4Stat = Tcp4Stat {
            local_addr_hex ,
            remote_addr_hex ,
            connection_state,
            tx_queue,
            rx_queue,
            time_active,
            jiffies_timer_expires,
            rto,
            uid,
            zero_window_probes,
            inode,
            etc
        };
        v.push(ele);
    }
    Ok(v)
}

impl Default for Tcp4Stat {
    fn default() -> Self {
        Tcp4Stat { local_addr_hex: String::new(),
             remote_addr_hex: String::new(), 
             connection_state: 0,
            tx_queue: 0,
            rx_queue: 0,
            time_active: 0,
            jiffies_timer_expires: 0,
            rto: 0,
            uid: 0,
            zero_window_probes: 0,
            inode : 0,
            etc: String::new()
        }
    }
}



