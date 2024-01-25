const 	P : u8 = 58;

#[inline]
pub fn make_format(size : usize,  seq : Box<dyn Iterator<Item = u8>>, data : &[u8]) -> Vec<u8> {
    let mut ret = Vec::new();
    let mut byte_arr : [u8;4] = [0;4] ;

    byte_arr.clone_from_slice((size as u32).to_le_bytes().as_slice());
    ret.extend(byte_arr);
    ret.push(P);
    ret.extend(seq);
    ret.push(P);
    ret.extend_from_slice("OS   ".as_bytes());
    ret.push(P);
    ret.extend(data);
    
    ret
}