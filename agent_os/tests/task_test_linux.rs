use std::sync::Arc;
use std::thread;

use agent_os::buffer::DoubleBuffer;
use agent_os::task::ipc_send_task_gen;
use agent_os::buffer::BufferAdder;
use agent_os::ipc::new_listener;
use agent_os::pool::thraed_pool;
use agent_os::pool::Pool;

use bson::Bson;

#[test]
fn ipc_send_task_test() {
    const PATH : &str = "/tmp/.task_test_linux.snap";
    let b  = DoubleBuffer::<(String,Bson)>::new();

    let tp = thraed_pool::ThreadPool::new(20);
    
    let mut i = 0;
    loop {
        i += 1;
        let mut addr = b.write().unwrap();

        let _ = addr.add((String::from("test1"),Bson::Int32(12)));
        let _ = addr.add((String::from("test2"),Bson::Int32(12)));
        let _ = addr.add((String::from("test3"),Bson::Int32(i)));
        drop(addr);

        thread::sleep(std::time::Duration::from_secs(1));

        let r = Arc::clone(&b);

        let mut l =new_listener(agent_os::ipc::ListenerKind::Mmap(String::from(PATH), 256)).unwrap();
    
        let fun = ipc_send_task_gen(l.get_stream().unwrap(),r);
        
        {
            let mut g = tp.lock().unwrap();
            let _ = g.use_item((), fun);
        }
    }
    
    
}

