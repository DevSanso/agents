mod tcp;
mod arp;
mod sockstat;
mod dev;

use std::io;

pub use tcp::read_tcp_stats;
pub use arp::read_arp_info;
pub use sockstat::read_sock_stat_info;
pub use dev::read_net_dev_info;

use arp::ArpInfo;
use tcp::Tcp4Stat;
use sockstat::SockStatInfo;
use dev::NetDevInfo;

use crate::search::machine::traits::OsNet;

#[derive(Clone)]
pub struct OsLinuxNetStat {
    arps : Vec<ArpInfo>,
    devs : Vec<NetDevInfo>,
    tcps : Vec<Tcp4Stat>,
    sockstat : SockStatInfo
}

impl OsLinuxNetStat {
    pub fn new() -> io::Result<Self> {
        let arps = read_arp_info()?;
        let devs = read_net_dev_info()?;
        let tcps = read_tcp_stats()?;
        let sockstat = read_sock_stat_info()?;
        Ok(OsLinuxNetStat {
            arps,
            devs,
            tcps,
            sockstat
        })
    } 
    pub fn get_args(&self) -> Vec<ArpInfo> {
        return self.arps.clone();
    }
    pub fn get_devs(&self) -> Vec<NetDevInfo> {
        return self.devs.clone();
    }
    pub fn get_tcps(&self) -> Vec<Tcp4Stat> {
        return self.tcps.clone();
    }
    pub fn get_sock_stat(&self) -> SockStatInfo {
        return self.sockstat.clone();
    }
}

impl OsNet<OsLinuxNetStat> for OsLinuxNetStat {
    fn total_rx(&self) -> u64 {
        self.devs.iter().fold(0, |acc, val| {return acc + val.rx_bytes})
    }
    fn total_tx(&self) -> u64 {
        self.devs.iter().fold(0, |acc, val| {return acc + val.rx_bytes})
    }
    fn total_use_sock(&self) -> u64 {
        self.sockstat.use_count
    }
    fn extend(&self) -> Box<OsLinuxNetStat> {
        Box::new(self.clone())
    }
}