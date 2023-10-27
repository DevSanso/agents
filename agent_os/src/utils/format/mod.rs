use zstd;

const 	VERTICAL_BAR : u8 = 124;

#[inline]
pub fn make_format(size : usize,  seq : Box<dyn Iterator<Item = u8>>, data : &[u8]) -> Vec<u8> {
    let mut ret = Vec::new();
    let mut byte_arr : [u8;4] = [0;4] ;

    byte_arr.clone_from_slice((size as u32).to_le_bytes().as_slice());
    ret.extend(byte_arr);
    ret.push(VERTICAL_BAR);
    ret.extend(seq);
    ret.push(VERTICAL_BAR);
    ret.extend(
        zstd::encode_all(data, 19).unwrap());
    
    ret
}