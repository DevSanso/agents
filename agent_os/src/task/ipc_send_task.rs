use std::sync::Arc;
use std::sync::RwLock;

use protobuf;

use crate::utils::buffer::BufferControllerAndReader;
use crate::ipc::IpcSendStream;
use crate::utils::result::result_change_err_is_string;
use crate::protos::os_snap::{Data, AgentOsSnap};
use crate::utils::util_time;

pub fn ipc_send_task_gen(
    mut stream: Box<dyn IpcSendStream + Sync + Send>,
    buf: Arc<RwLock<dyn BufferControllerAndReader<Data> + Send + Sync>>,
) -> impl FnOnce(()) -> Result<(), String> + Send {

    let func = move |()| -> Result<(), String> {

        let controller = buf;

        {
            let mut g = result_change_err_is_string(controller.write())?;
            result_change_err_is_string(g.swtich())?;
        }

        let g = result_change_err_is_string(controller.read())?;

        let send_data = result_change_err_is_string(g.read())?;

        if send_data.len() <= 0 {
            return Ok(());
        }

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

        Ok(())
    };

    func
}
