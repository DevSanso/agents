pub mod traits;
#[cfg(target_os =  "linux")]
pub mod linux;

use std::io;

#[cfg(target_os =  "linux")]
type LinuxOsMachine  = dyn traits::OSMachine<linux::net::OsLinuxNetStat>;

#[cfg(target_os =  "linux")]
pub fn get_machine() -> io::Result<Box<LinuxOsMachine>> {
    linux::LinuxMachine::new()
}