use std::fs::File;
use std::io;
use std::io::Read;

use memmap2;

use agent_os::ipc::new_listener;
use agent_os::ipc::ListenerKind;

use agent_os::utils::format::make_format;
use agent_os::utils::seq::new_seq;
use agent_os::utils::seq::SequenceKind;
#[test]
pub fn ipc_mmap_write_test() -> io::Result<()>{
    const MSG : &str= "hello world! ";
    
    let mut seq = new_seq(SequenceKind::U64(0));

    let mut output = make_format(MSG.len(),seq.next(), MSG.as_bytes());
    let mut listener =new_listener(ListenerKind::Mmap(String::from("/tmp/ipc_mmap_write_test_mmap"), 36))?;
    let mut  stream = listener.get_stream()?;

    let read = unsafe {
        let f = File::open("/tmp/ipc_mmap_write_test_mmap")?;
        memmap2::MmapOptions::new().map(&f)?
    };

    stream.send(MSG.as_bytes())?;
    let mut read_data = Vec::new();
    let file_date = read.as_ref();
    assert_eq!(file_date.len(), 36);
    read_data.extend_from_slice(file_date);
    assert_eq!(read_data.as_slice(), output.as_slice());

    Ok(())
}

