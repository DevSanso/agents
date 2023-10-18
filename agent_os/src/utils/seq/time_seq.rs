use std::sync::{Mutex,Arc};
use std::time;

use crate::utils::seq::Sequence;

pub struct UnixEpocheSequence {
    unix_time_mut : Mutex<time::Duration>
}

impl UnixEpocheSequence {
    pub(crate) fn new() -> Self {
        UnixEpocheSequence  {
            unix_time_mut : Mutex::new(time::Duration::new(0, 0))
        }
    }
}

impl Sequence for UnixEpocheSequence {
    fn current(&self) -> Box<dyn Iterator<Item = u8>> {
        let g = self.unix_time_mut.lock().unwrap();
        Box::new(g.as_secs().to_le_bytes().into_iter())
    }

    fn update(&mut self) {
        let mut g = self.unix_time_mut.lock().unwrap();
        *g = time::SystemTime::now()
            .duration_since(time::UNIX_EPOCH)
            .unwrap();
    }

    fn next(&mut self) ->Box<dyn Iterator<Item = u8>> {
        let mut g = self.unix_time_mut.lock().unwrap();
        *g = time::SystemTime::now()
            .duration_since(time::UNIX_EPOCH)
            .unwrap();
        let ret = Box::new(g.as_secs().to_le_bytes().into_iter());
        ret
    }
}