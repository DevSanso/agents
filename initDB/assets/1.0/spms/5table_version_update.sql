create table spms.version_update (
    major int,
    minor int
    update_time timestamp,
    name varchar(256)
)

create index version_update_idx on spms.version_update(major, minor);