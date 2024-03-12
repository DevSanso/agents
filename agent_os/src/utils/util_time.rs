use std::time::{Duration, SystemTime, UNIX_EPOCH};

#[inline]
pub fn get_unix_epoch_now() -> Duration {
    SystemTime::now()
        .duration_since(UNIX_EPOCH)
        .unwrap()
}

#[inline]
pub fn is_interval(now : Duration, interval_sec : Duration) -> bool {
    if now.as_secs() % interval_sec.as_secs() == 0 {
        if now.as_millis() % 1000 < 101 {
            return true;
        }
    }

    false
}