use std::io;

use agent_os::buffer::DoubleBuffer;
use agent_os::buffer::BufferAdder;
use agent_os::buffer::BufferController;
use agent_os::buffer::BufferReader;
use agent_os::utils::result::result_cast_to_io_result;

#[derive(Clone)]
struct TestData {
    pub index : u64,
    pub _text : String
}
#[test]
pub fn ipc_mmap_write_test() -> io::Result<()>{
    let origin = DoubleBuffer::<TestData>::new();
    let mut b = result_cast_to_io_result(origin.write())?;
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