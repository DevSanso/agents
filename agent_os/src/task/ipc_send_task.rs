use std::sync::Arc;
use std::sync::RwLock;

use bson::Bson;

use crate::buffer::BufferControllerAndReader;
use crate::ipc::IpcSendStream;
use crate::utils::result::result_change_err_is_string;

pub fn ipc_send_task_gen(
    mut stream: Box<dyn IpcSendStream + Sync + Send>,
    buf: Arc<RwLock<dyn BufferControllerAndReader<(String, Bson)> + Send + Sync>>,
) -> impl FnOnce(()) -> Result<(), String> + Send {

    let func = move |()| -> Result<(), String> {

        let controller = buf;

        let mut g = result_change_err_is_string(controller.write())?;

        result_change_err_is_string(g.swtich())?;

        drop(g);

        let g = result_change_err_is_string(controller.read())?;

        let send_data = result_change_err_is_string(g.read())?;

        let collected = send_data
            .into_iter()
            .fold(bson::Document::new(), |mut acc, x| {
                acc.insert(x.0, x.1);
                acc
            });

        let mut pack_data = Vec::<u8>::new();

        result_change_err_is_string(collected.to_writer(&mut pack_data))?;

        result_change_err_is_string(stream.send(pack_data.as_slice()))?;

        Ok(())
    };

    func
}
