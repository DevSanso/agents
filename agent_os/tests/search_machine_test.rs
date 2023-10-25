use std::io;


use agent_os::search::machine;

#[test]
pub fn machine_test() -> io::Result<()>{
    let m = machine::get_machine()?;
    for _i in 0..100 {
        m.os_net_stat()?;
    }
    Ok(())
}