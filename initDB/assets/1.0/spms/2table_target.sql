create table spms.target (
    target_id int primary key,
    ip char(16),
    port int,
    target_class_id int
);

create index target_idx on spms.target(target_id, target_class_id);