use std::sync::Arc;
use std::sync::RwLock;

use protobuf::{EnumOrUnknown, Message};

use crate::utils::buffer::BufferAdder;
use crate::search::machine::get_machine;
use crate::utils::result::result_change_err_is_string;
use crate::protos::os_snap::{Data, DataFormat};

#[cfg(target_os = "linux")]
pub fn os_details_net_stat_thread_gen(
    buf: Arc<RwLock<dyn BufferAdder<Data> + Sync + Send>>,
) -> impl FnOnce(()) -> Result<(), String> + Send {
    let fun = move |_tup| -> Result<(), String> {
        let search = result_change_err_is_string(get_machine())?;
        let addr = buf;
        {
            let data = result_change_err_is_string(search.os_net_stat())?;
            let details_data = data.extend();

            let mut g = result_change_err_is_string(addr.write())?;

            {
                let mut args_data = Data::new();
                args_data.format = EnumOrUnknown::<DataFormat>::new(DataFormat::NetArp);
                args_data.raw_data = result_change_err_is_string(
                    Message::write_to_bytes(&details_data.get_args()))?;

                g.add(args_data).unwrap();
            }

            {
                let mut dev_data = Data::new();
                dev_data.format = EnumOrUnknown::<DataFormat>::new(DataFormat::NetDev);
                dev_data.raw_data = result_change_err_is_string(
                    Message::write_to_bytes(&details_data.get_devs()))?;
                
                g.add(dev_data).unwrap();
            }

            {
                let mut sock_stat_data = Data::new();
                sock_stat_data.format = EnumOrUnknown::<DataFormat>::new(DataFormat::NetSockStat);
                sock_stat_data.raw_data = result_change_err_is_string(
                    Message::write_to_bytes(&details_data.get_sock_stat()))?;
                
                g.add(sock_stat_data).unwrap();
            }

            {
                let mut net_tcp_4_data = Data::new();
                net_tcp_4_data.format = EnumOrUnknown::<DataFormat>::new(DataFormat::NetTcp4Stat);
                net_tcp_4_data.raw_data = result_change_err_is_string(
                    Message::write_to_bytes(&details_data.get_tcps()))?;
                
                g.add(net_tcp_4_data).unwrap();
            }
        }

        Ok(())
    };

    fun
}
