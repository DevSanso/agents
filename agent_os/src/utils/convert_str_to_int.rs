use std::io;



#[inline]
pub fn convert_str_to_u8(s : &'_ str, radix: u32) -> io::Result<u8> {
    match u8::from_str_radix(s,radix) {
        Ok(ok) => Ok(ok),
        Err(err) => Err(io::Error::new(io::ErrorKind::InvalidData, format!("convert_str_to_u8 - {}  [data:{}]",err.to_string(),s)))
    }
}

#[inline]
pub fn convert_str_to_u64(s : &'_ str, radix: u32) -> io::Result<u64> {
    match u64::from_str_radix(s, radix) {
        Ok(ok) => Ok(ok),
        Err(err) => Err(io::Error::new(io::ErrorKind::InvalidData, format!("convert_str_to_u64 - {}  [data:{}]",err.to_string(),s)))
    }
}

#[inline]
pub fn convert_str_to_u32(s : &'_ str, radix: u32) -> io::Result<u32> {
    match  u32::from_str_radix(s, radix){
        Ok(ok) => Ok(ok),
        Err(err) => Err(io::Error::new(io::ErrorKind::InvalidData, format!("convert_str_to_u32 - {}  [data:{}]",err.to_string(),s)))
    }
}
