use std::io;

use simplelog::*;
use agent_os::search::machine;

#[test]
pub fn machine_test() -> io::Result<()>{
    let _ = CombinedLogger::init(vec![SimpleLogger::new(LevelFilter::Trace, Config::default())]);
    let m = machine::get_machine()?;
    for _i in 0..100 {
        m.os_net_stat()?;
    }
    Ok(())
}