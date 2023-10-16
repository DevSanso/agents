use std::io;

pub trait OsNet<T> {
    fn total_tx(&self) -> u64;
    fn total_rx(&self) -> u64;
    fn total_use_sock(&self) -> u64;
    fn extend(&self) -> Box<T>;
}
pub trait OSMachine<NetT> {
    fn os_net_stat(&self) -> io::Result<Box<dyn OsNet<NetT>>>;
}