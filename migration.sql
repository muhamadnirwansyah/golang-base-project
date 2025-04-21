create table if not exists account(
    id bigserial primary key,
    full_name varchar(200)not null,
    email varchar(200) not null unique,
    phone_number varchar(20) not null unique, 
    password varchar(255) not null
);
