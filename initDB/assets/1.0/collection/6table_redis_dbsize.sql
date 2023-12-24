create table redis.redis_dbsize (
    target_id int,
    collection_time timestamp,
    size int8
)partition by range(collection_time);

create index redis_dbsize_idx on redis.redis_dbsize(target_id, collection_time);