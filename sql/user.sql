create database if not exists test;
use test;
create table users (
    id int not null auto_increment primary key,
    name varchar(255),
    age int
);
insert into users (name, age) values
("Michaele", 20),
("Lesia", 21),
("Ambrose", 22),
("Domenic", 23),
("Bong", 24),
("Versie", 25),
("Angeli", 25),
("Vince", 24),
("Britni", 24),
("Faye", 23);
