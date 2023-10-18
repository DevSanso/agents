use std::sync::{Mutex,Arc};

use crate::utils::seq::Sequence;

pub struct U64Sequence {
    current_val : Mutex<u64>
}

impl U64Sequence {
    pub(crate) fn new(start : u64) -> Self {
        U64Sequence { current_val:Mutex::new(start) }
    }
}

impl Sequence for U64Sequence {
    fn current(&self) -> Box<dyn Iterator<Item = u8>> {
        let g = self.current_val.lock().unwrap();
        Box::new(g.to_le_bytes().into_iter())
    }

    fn update(&mut self) {
        let mut g = self.current_val.lock().unwrap();
        *g = *g + 1;
    }

    fn next(&mut self) ->Box<dyn Iterator<Item = u8>> {
        let mut g = self.current_val.lock().unwrap();
        *g = *g + 1;
        let ret = Box::new(g.to_le_bytes().into_iter());
        ret
    }
}