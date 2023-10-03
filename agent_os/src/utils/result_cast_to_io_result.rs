use std::io;
use std::fmt::{Debug, Display};

#[inline]
pub fn result_cast_to_io_result<T,E : Debug + Display >(res : Result<T,E>) -> io::Result<T> {
    match res {
        Ok(ok) => Ok(ok),
        Err(err) => Err(io::Error::new(io::ErrorKind::InvalidData, format!("result_cast_to_io_result - {}", err)))
    }
}