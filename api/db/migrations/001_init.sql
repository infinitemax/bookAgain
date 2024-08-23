create extension if not exists "uuid-ossp";

drop schema if exists books cascade;
create schema books;

create table users(
    id uuid primary key default uuid_generate_v4(),
    user_name text not null unique,
    email text not null unique,
    created_at timestamp without time zone not null default (now() at time zone 'utc'),
    deleted_at timestamp without time zone default null
);

create table books(
    id uuid primary key default uuid_generate_v4(),
    title text,
    author text
);

create table notes(
    id uuid primary key default uuid_generate_v4(),
    title text,
    note text,
    book uuid not null references "books" (id),
    created_by uuid not null references "users" (id)
);

---- create above / drop below ----

drop table users;
