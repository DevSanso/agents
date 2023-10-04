use std::io;

#[cfg(target_os =  "linux")]
use agent_os::search::machine::linux::net::read_net_dev_info;

#[test]
#[cfg(target_os =  "linux")]
pub fn linux_arp_search_test() -> io::Result<()>{
    read_net_dev_info()?;
    Ok(())
}