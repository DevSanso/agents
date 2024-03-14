use std::sync::Arc;
use std::sync::RwLock;
use std::time;

use protobuf;

use crate::utils;
use crate::structure::buffer::BufferControllerAndReader;
use crate::ipc::IpcSendStream;
use crate::utils::result::result_change_err_is_string;
use crate::protos::os_snap::{Data, AgentOsSnap};
use crate::utils::util_time;

const IPC_SEND_INTERVAL: u64 = 2;

pub fn ipc_send_task_gen(
    stream: Box<dyn IpcSendStream + Sync + Send>,
    buf: Arc<RwLock<dyn BufferControllerAndReader<Data> + Send + Sync>>,
) -> impl FnOnce(()) -> Result<(), String> + Send {

    let func = move |()| -> Result<(), String> {
        ipc_send_task(stream, buf)
    };

    func
}

fn ipc_send_task(
    mut stream: Box<dyn IpcSendStream + Sync + Send>,
    buf: Arc<RwLock<dyn BufferControllerAndReader<Data> + Send + Sync>>
) -> Result<(), String> {

    loop {
        let now = utils::util_time::get_unix_epoch_now();
        if !utils::util_time::is_interval(now, time::Duration::from_secs(IPC_SEND_INTERVAL))  { continue; }
        {
            let mut g = result_change_err_is_string(buf.write())?;
            result_change_err_is_string(g.swtich())?;
        }

        let g = result_change_err_is_string(buf.read())?;

        let send_data = result_change_err_is_string(g.read())?;

        if send_data.len() <= 0 { continue; }

        let mut collected = send_data
        .into_iter()
        .fold(AgentOsSnap::new(), |mut acc, x| {
            acc.datas.push(x);
            acc
        });

        collected.unix_epoch = util_time::get_unix_epoch_now().as_secs();
        
        let pack_data = result_change_err_is_string(
            protobuf::Message::write_to_bytes(&collected))?;
        
        result_change_err_is_string(stream.send(pack_data.as_slice()))?;
    }

    //Ok(())
}