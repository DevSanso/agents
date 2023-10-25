use std::sync::Arc;
use std::sync::RwLock;

use bson::Bson;

use crate::buffer::BufferAdder;
use crate::search::machine::get_machine;
use crate::utils::result::result_change_err_is_string;

pub fn total_net_stat_thread_gen(
    buf: Arc<RwLock<dyn BufferAdder<(String, Bson)> + Sync + Send>>,
) -> impl FnOnce(()) -> Result<(), String> + Send {
    let fun = move |_tup| -> Result<(), String> {
        let search = result_change_err_is_string(get_machine())?;
        let addr = buf;
        {
            let data = result_change_err_is_string(search.os_net_stat())?;

            let doc = bson::doc! {
                "tx" : data.total_tx() as i64,
                "rx" : data.total_rx() as  i64,
                "use_sock" : data.total_use_sock() as i64
            };

            let mut g = result_change_err_is_string(addr.write())?;

            result_change_err_is_string(
                g.add((String::from("network_total"), Bson::Document(doc))),
            )?;
        }

        Ok(())
    };

    fun
}

#[cfg(target_os = "linux")]
pub fn os_details_net_stat_thread_gen(
    buf: Arc<RwLock<dyn BufferAdder<(String, Bson)> + Sync + Send>>,
) -> impl FnOnce(()) -> Result<(), String> + Send {
    let fun = move |_tup| -> Result<(), String> {
        let search = result_change_err_is_string(get_machine())?;
        let addr = buf;

        {
            let data = result_change_err_is_string(search.os_net_stat())?;
            let details_data = data.extend();

            let args =
                result_change_err_is_string(bson::to_bson(details_data.get_args().as_slice()))?;
            let devs =
                result_change_err_is_string(bson::to_bson(details_data.get_devs().as_slice()))?;
            let sock_stat =
                result_change_err_is_string(bson::to_bson(&details_data.get_sock_stat()))?;
            let tcps =
                result_change_err_is_string(bson::to_bson(details_data.get_tcps().as_slice()))?;

            let doc = bson::doc! {
                "args" : args,
                "devs" : devs,
                "sock_stat" : sock_stat,
                "tcps" : tcps
            };

            let mut g = result_change_err_is_string(addr.write())?;

            result_change_err_is_string(
                g.add((String::from("network_os_linux"), Bson::Document(doc))),
            )?;
        }

        Ok(())
    };

    fun
}
