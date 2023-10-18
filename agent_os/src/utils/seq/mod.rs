mod time_seq;
mod u64_seq;

pub enum SequenceKind {
    U64(u64),
    UnixEpoche
}

pub trait Sequence {
    fn current(&self) -> Box<dyn Iterator<Item = u8>>;
    fn update(&mut self);
    fn next(&mut self) ->Box<dyn Iterator<Item = u8>>;
}

pub fn new_seq(k : SequenceKind) -> Box<dyn Sequence> {
    match k {
        SequenceKind::U64(start) => Box::new(u64_seq::U64Sequence::new(start)),
        SequenceKind::UnixEpoche => Box::new(time_seq::UnixEpocheSequence::new())
    }
}

