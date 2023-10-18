use std::io;
use agent_os::buffer::new_buffer;
use agent_os::buffer::BufferKind;

#[derive(Clone)]
struct TestData {
    pub index : u64,
    pub text : String
}
#[test]
pub fn ipc_mmap_write_test() -> io::Result<()>{
    let mut b = new_buffer::<TestData>(BufferKind::DoubleBuffer);

    for i in 0..10000 {
        b.add(TestData {index : i % 100, text : String::from("text")})?;
    }
    b.swtich()?;

    let r = b.read()?;
    for i in 0..10000 {
        assert_eq!(r[i].index, (i  as u64) % 100);
    }
    Ok(())
}