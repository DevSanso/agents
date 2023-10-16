use std::io;


use agent_os::search::machine;

#[test]
pub fn machine_test() -> io::Result<()>{
    let m = machine::get_machine()?;
    m.os_net_stat()?;
    Ok(())
}