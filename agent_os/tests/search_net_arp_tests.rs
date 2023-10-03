use std::io;

#[cfg(target_os =  "linux")]
use agent_os::search::machine::linux::net::reead_arp_info;

#[test]
#[cfg(target_os =  "linux")]
pub fn linux_arp_search_test() -> io::Result<()>{
    reead_arp_info()?;
    Ok(())
}