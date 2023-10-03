use std::io;

#[cfg(target_os =  "linux")]
use agent_os::search::machine::linux::net::read_tcp_stats;

#[test]
#[cfg(target_os =  "linux")]
pub fn linux_tcp_search_test() -> io::Result<()>{
    read_tcp_stats()?;
    Ok(())
}