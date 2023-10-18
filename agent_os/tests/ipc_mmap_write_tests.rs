use std::fs::File;
use std::io;

use memmap2;

use agent_os::ipc::new_listener;
use agent_os::ipc::ListenerKind;

use agent_os::ipc::make_format;
use agent_os::utils::seq::new_seq;
use agent_os::utils::seq::SequenceKind;
#[test]
pub fn ipc_mmap_write_test() -> io::Result<()>{
    const MSG : &str= "hello world!";
    
    let mut seq = new_seq(SequenceKind::U64(0));
    
    let mut output = make_format(seq.next(), MSG.as_bytes());
    let mut output2 = make_format(seq.next(), MSG.as_bytes());
    let mut listener =new_listener(ListenerKind::Mmap(String::from("/tmp/ipc_mmap_write_test_mmap"), 32))?;
    let mut  stream = listener.get_stream()?;

    let read = unsafe {
        let f = File::open("/tmp/ipc_mmap_write_test_mmap")?;
        memmap2::MmapOptions::new().map(&f)?
    };

    stream.send(MSG.as_bytes())?;
    let read_data = read.as_ref();
    output.resize(32, 0);
    assert_eq!(read_data, output.as_slice());

    stream.send(MSG.as_bytes())?;
    let read_data2 = read.as_ref();
    output2.resize(32, 0);
    assert_eq!(read_data2, output2.as_slice());

    Ok(())
}

