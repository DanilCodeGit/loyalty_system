-- +goose Up
create table users(
    id serial primary key ,
    login varchar unique not null,
    password varchar not null,
    balance int
);

-- +goose Down
drop table users;
