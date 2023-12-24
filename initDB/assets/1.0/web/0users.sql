create table web.users(
    email varchar(128) primary key,
    user_uuid uuid default uuid_generate_v4() unique,
    id varchar(16) unique not null,
    password varchar(256) not null
);