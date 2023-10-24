use std::sync::Arc;
use std::sync::RwLock;

use bson::Bson;
use zstd;

use crate::buffer::BufferControllerAndReader;
use crate::ipc::IpcListener;
use crate::utils::result::result_change_err_is_string;

pub fn ipc_send_task_gen(
    mut listen: Box<dyn IpcListener + Sync + Send>,
    buf: Arc<RwLock<dyn BufferControllerAndReader<(String, Bson)> + Send + Sync>>,
) -> impl FnOnce(()) -> Result<(), String> + Send {

    let func = move |()| -> Result<(), String> {

        let controller = buf;

        let mut stream = result_change_err_is_string(listen.get_stream())?;

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

        let compress_data =
            result_change_err_is_string(zstd::encode_all(pack_data.as_slice(), 19))?;

        result_change_err_is_string(stream.send(compress_data.as_slice()))?;

        Ok(())
    };

    func
}
