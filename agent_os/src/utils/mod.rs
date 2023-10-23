mod opt_cast_io_result;
mod result_cast_to_io_result;
mod result_change_err_is_string;

pub mod option {
    pub use super::opt_cast_io_result::opt_cast_to_io_result;
}
pub mod result {
    pub use super::result_cast_to_io_result::result_cast_to_io_result;
    pub use super::result_change_err_is_string::result_change_err_is_string;
}

pub mod seq;

