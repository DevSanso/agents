create table spms.target_class (
    class_id int primary key,
    name varchar(64)
)

create index target_class_idx on spms.target_class(class_id);