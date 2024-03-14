use std::sync::Arc;
use std::sync::RwLock;
use std::time;

use protobuf::{EnumOrUnknown, Message};

use crate::utils;
use crate::search::machine::linux::net::OsLinuxNetStat;
use crate::structure::buffer::BufferAdder;
use crate::search::machine::get_machine;
use crate::utils::result::result_change_err_is_string;
use crate::protos::os_snap::{Data, DataFormat};

const OS_DETAILS_NET_INTERVAL: u64 = 6;

#[cfg(target_os = "linux")]
pub fn os_details_net_stat_thread_gen(
    buf: Arc<RwLock<dyn BufferAdder<Data> + Sync + Send>>,
) -> impl FnOnce(()) -> Result<(), String> + Send {

    let fun = move |_tup| -> Result<(), String> {
        os_details_net_stat_thread(buf)
    };

    fun
}

#[inline]
fn get_net_detail_stats(details : &Box<OsLinuxNetStat>, format : DataFormat) -> Result<Data, String> {
    let mut data = Data::new();
    let raw_data = match &format {
        DataFormat::NetArp => {
            result_change_err_is_string(
                Message::write_to_bytes(&details.get_args()))?
        },
        DataFormat::NetDev => {
            result_change_err_is_string(
                Message::write_to_bytes(&details.get_devs()))?
        },
        DataFormat::NetSockStat => {
            result_change_err_is_string(
                Message::write_to_bytes(&details.get_sock_stat()))?
        },
        DataFormat::NetTcp4Stat => {
            result_change_err_is_string(
                Message::write_to_bytes(&details.get_tcps()))?
        },
    };

    data.format = EnumOrUnknown::<DataFormat>::new(format);
    data.raw_data = raw_data;

    Ok(data)
}

fn os_details_net_stat_thread(buf: Arc<RwLock<dyn BufferAdder<Data> + Sync + Send>>) -> Result<(), String> {
    let search = result_change_err_is_string(get_machine())?;
    let stat = result_change_err_is_string(search.os_net_stat())?;
    let details_stat = stat.extend();
    let mut g = result_change_err_is_string(buf.write())?;

    loop {
        let now = utils::util_time::get_unix_epoch_now();
        if !utils::util_time::is_interval(now, time::Duration::from_secs(OS_DETAILS_NET_INTERVAL))  { continue; }

        {
            let data = get_net_detail_stats(&details_stat, DataFormat::NetArp)?;
            g.add(data).unwrap();
        }

        {
            let data = get_net_detail_stats(&details_stat, DataFormat::NetDev)?;
            g.add(data).unwrap();
        }

        {
            let data = get_net_detail_stats(&details_stat, DataFormat::NetSockStat)?;
            g.add(data).unwrap();
        }

        {
            let data = get_net_detail_stats(&details_stat, DataFormat::NetTcp4Stat)?;
            g.add(data).unwrap();
        }
    }

    //Ok(())
}