pub mod net;


use std::io;

use crate::search::machine::traits::OSMachine;
use crate::search::machine::traits::OsNet;

pub struct LinuxMachine;

impl LinuxMachine {
    pub fn new() -> io::Result<Box<dyn  OSMachine<net::OsLinuxNetStat>>> {
        Ok(Box::new(LinuxMachine))
    }
}

impl OSMachine<net::OsLinuxNetStat> for LinuxMachine {
    fn os_net_stat(&self) -> io::Result<Box<dyn OsNet<net::OsLinuxNetStat>>> {
        Ok(Box::new( net::OsLinuxNetStat::new()?))
    }
}


