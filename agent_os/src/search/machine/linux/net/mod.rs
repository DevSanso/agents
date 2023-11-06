mod tcp;
mod arp;
mod sockstat;
mod dev;

use std::io;

pub use tcp::read_tcp_stats;
pub use arp::read_arp_info;
pub use sockstat::read_sock_stat_info;
pub use dev::read_net_dev_info;

use crate::protos::net::ArpInfos;
use crate::protos::net::NetDevInfos;
use crate::protos::net::Tcp4Stats;
use crate::protos::net::SockStatInfo;

use crate::search::machine::traits::OsNet;


#[derive(Clone)]
pub struct OsLinuxNetStat {
    arps : ArpInfos,
    devs : NetDevInfos,
    tcps : Tcp4Stats,
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
    pub fn get_args(&self) -> ArpInfos {
        return self.arps.clone();
    }
    pub fn get_devs(&self) -> NetDevInfos {
        return self.devs.clone();
    }
    pub fn get_tcps(&self) -> Tcp4Stats {
        return self.tcps.clone();
    }
    pub fn get_sock_stat(&self) -> SockStatInfo {
        return self.sockstat.clone();
    }
}

impl OsNet<OsLinuxNetStat> for OsLinuxNetStat {
    fn total_rx(&self) -> u64 {
        self.devs.infos.iter().fold(0, |acc, val| {return acc + val.rx_bytes})
    }
    fn total_tx(&self) -> u64 {
        self.devs.infos.iter().fold(0, |acc, val| {return acc + val.rx_bytes})
    }
    fn total_use_sock(&self) -> u64 {
        self.sockstat.use_count
    }
    fn extend(&self) -> Box<OsLinuxNetStat> {
        Box::new(self.clone())
    }
}