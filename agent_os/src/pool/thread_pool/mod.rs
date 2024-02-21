mod thread_impl;
mod thread_pool;

pub(in crate::pool::thread_pool) use thread_impl::ThreadImpl;
pub(in crate::pool::thread_pool) use thread_impl::ThreadSignal;

pub use thread_pool::ThreadPool;