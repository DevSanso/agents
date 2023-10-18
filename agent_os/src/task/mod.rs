mod ipc_send_task;
pub trait Task<'a> {
    fn start(&mut self);
    fn kill(&mut self);
    fn stop(&mut self);
    fn reflesh(&mut self);
}

pub use ipc_send_task::new_ipc_send_task;
