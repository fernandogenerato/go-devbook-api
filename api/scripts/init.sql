CREATE DATABASE IF NOT EXISTS devbook;
use devbook;

DROP TABLE IF EXISTS users;


create table devbook.users
(
    id        int auto_increment primary key,
    name      varchar(50) not null,
    nick      varchar(50) not null unique,
    email     varchar(50) not null unique,
    password  varchar(25) not null unique,
    createdAt timestamp default current_timestamp()
) engine = innodb;

