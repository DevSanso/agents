create table spms.auths(
    target_id int,
    type varchar(3),
    username varchar(128),
    password varchar(256)
);

create index redis_auth_idx on spms.auths(target_id, type);