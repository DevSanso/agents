use std::io;

#[inline]
pub fn opt_cast_to_io_result<T>(opt : Option<T>, msg : &'_ str) -> io::Result<T> {
    match opt {
        Some(some) => Ok(some),
        None => Err(io::Error::new(io::ErrorKind::InvalidData, format!("opt cast_to_io_result - {}", msg)))
    }
}
