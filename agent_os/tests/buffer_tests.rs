use std::io;

use agent_os::structure::buffer::DoubleBuffer;
use agent_os::structure::buffer::BufferAdder;
use agent_os::structure::buffer::BufferController;
use agent_os::structure::buffer::BufferReader;

use agent_os::utils::convert_to_io_result;

use simplelog::*;
use log::*;

#[derive(Clone)]
struct TestData {
    pub index : u64,
    pub _text : String
}

#[test]
pub fn ipc_mmap_write_test() -> io::Result<()>{
    let _ = CombinedLogger::init(vec![SimpleLogger::new(LevelFilter::Trace, Config::default())]);

    let origin = DoubleBuffer::<TestData>::new();
    let mut b = convert_to_io_result!(result, origin.write())?;
    for i in 0..10000 {
        b.add(TestData {index : i % 100, _text : String::from("text")})?;
    }
    b.swtich()?;

    let r = b.read()?;
    for i in 0..10000 {
        assert_eq!(r[i].index, (i  as u64) % 100);
    }
    Ok(())
}