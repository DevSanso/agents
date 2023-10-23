use std::sync::Arc;

use agent_os::buffer::DoubleBuffer;
use agent_os::task::ipc_send_thread_gen;
use agent_os::task::total_net_stat_thread_gen;
use agent_os::buffer::BufferAdder;
use agent_os::ipc::new_listener;

use bson::Bson;

#[test]
fn a() {
    const path : &str = "/tmp/.task_test_linux.snap";
    let mut b  = DoubleBuffer::<(String,Bson)>::new();
    let r = Arc::clone(&b);

    let l =new_listener(agent_os::ipc::ListenerKind::Mmap(String::from(path), 0)).unwrap();

    let mut addr = b.lock().unwrap();

    addr.add((String::from("test1"),Bson::Int32(12)));
    addr.add((String::from("test2"),Bson::Int32(12)));
    addr.add((String::from("test3"),Bson::Int32(12)));

    
    let job = ipc_send_thread_gen(l, r);



    job.join();

    
}

