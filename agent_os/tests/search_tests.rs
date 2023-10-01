use std::io;
use agent_os::search::machine::net::tcp::read_tcp4_stats;

#[test]
pub fn tcp4_search_test() -> io::Result<()>{
    read_tcp4_stats()?;
    Ok(())
}