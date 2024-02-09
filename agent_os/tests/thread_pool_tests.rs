use std::thread::sleep;
use std::time::Duration;
use std::{io, ops::Deref};

use agent_os::pool::ThreadPool;
use agent_os::pool::Pool;



fn print_ten(args : ()) -> Result<(),String>{
    for i in 0..10 {
        println!("print_ten : {}",i);
    }
    Ok(())
}

fn sleep_and_print(args : ()) -> Result<(), String> {
    std::thread::sleep(std::time::Duration::from_secs(10));
    println!("sleep_and_print");
    Ok(())
}

#[test]
pub fn thread_pool_test() -> io::Result<()>{
    let tp = ThreadPool::new(1, 4);
    let mut tp_mut = tp.lock().unwrap();

    tp_mut.run_func((), print_ten)?;
    sleep(Duration::from_secs(1));
    tp_mut.run_func((), sleep_and_print)?;
    sleep(Duration::from_secs(1));
    tp_mut.run_func((), sleep_and_print)?;
    sleep(Duration::from_secs(1));
    tp_mut.run_func((), print_ten)?;

    sleep(Duration::from_secs(1));
    println!("used count : {}", tp_mut.used_count());
    sleep(Duration::from_secs(20));


    Ok(())
}