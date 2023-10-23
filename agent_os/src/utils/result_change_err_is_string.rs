use std::error::Error;

#[inline]
pub fn result_change_err_is_string<T,E : Error>(origin : Result<T,E> ) -> Result<T, String> {
    if origin.is_err() {
        return Err(origin.err().unwrap().to_string());
    }

    Ok(origin.ok().unwrap())
}