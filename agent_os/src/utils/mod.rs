mod opt_cast_io_result;
mod result_cast_to_io_result;
mod result_change_err_is_string;



pub mod option {
    use super::opt_cast_io_result;
    pub use opt_cast_io_result::opt_cast_to_io_result;
}

pub mod result {
    use super::result_cast_to_io_result;
    use super::result_change_err_is_string;

    pub use result_cast_to_io_result::result_cast_to_io_result;
    pub use result_change_err_is_string::result_change_err_is_string;
}

#[macro_export]
macro_rules! convert_to_io_result {
    (option, $opt:ident, $msg:expr) => {
        {
            use log::*;
            #[cfg(test)] use agent_os::utils::option::opt_cast_to_io_result;
            #[cfg(not(test))] use crate::utils::option::opt_cast_to_io_result;

            if $opt.is_none()  {debug!("[FILE:{}] - [LINE:{}] - Option Is None - MSG:{}", file!(), line!(), $msg)}
            opt_cast_to_io_result($opt, $msg)
        }
    };
    (option, $opt:expr, $msg:expr) => {
        {
            use log::*;
            #[cfg(test)] use agent_os::utils::option::opt_cast_to_io_result;
            #[cfg(not(test))] use crate::utils::option::opt_cast_to_io_result;
            let opt_ret = $opt;
            if opt_ret.is_none()  {debug!("[FILE:{}] - [LINE:{}] - Option Is None - MSG:{}", file!(), line!(), $msg)}
            opt_cast_to_io_result(opt_ret, $msg)
        }
    };

    (result, $ret:expr) => {
        {
            use log::*;
            #[cfg(test)] use agent_os::utils::result::result_cast_to_io_result;
            #[cfg(not(test))] use crate::utils::result::result_cast_to_io_result;

            if $ret.is_err() {debug!("[FILE:{}] - [LINE:{}] - Result Is Error", file!(), line!())}
            result_cast_to_io_result($ret)
        }
    };
}
pub use convert_to_io_result;

pub mod util_time;
pub mod format;


