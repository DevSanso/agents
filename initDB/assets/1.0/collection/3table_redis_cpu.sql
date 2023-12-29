create table collection_redis.redis_cpu (
    target_id int,
    collection_time timestamp,
    use_cpu_sys numeric,
    use_cpu_user numeric,
    use_cpu_sys_child numeric,
    use_cpu_user_child numeric,
    use_cpu_sys_main numeric,
    use_cpu_user_main numeric
) partition by range(collection_time);

create index redis_cpu_idx on collection_redis.redis_cpu(target_id, collection_time);
