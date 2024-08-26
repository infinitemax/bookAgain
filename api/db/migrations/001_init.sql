create extension if not exists "uuid-ossp";

drop schema if exists books cascade;
create schema books;

create table books.users(
    id uuid primary key default uuid_generate_v4(),
    user_name text not null unique,
    email text not null unique,
    created_at timestamp without time zone not null default (now() at time zone 'utc'),
    deleted_at timestamp without time zone default null
);

create table books.books(
    id uuid primary key default uuid_generate_v4(),
    title text not null,
    author text not null,
    user_id UUID NOT NULL REFERENCES books.users(id) ON DELETE CASCADE,
    created_at timestamp without time zone not null default (now() at time zone 'utc'),
    deleted_at timestamp without time zone default null
);

create table books.notes(
    id uuid primary key default uuid_generate_v4(),
    title text,
    note text not null,
    book uuid not null references books.books(id),
    page_number integer not null,
    created_by uuid not null references books.books(id),
    created_at timestamp without time zone not null default (now() at time zone 'utc'),
    deleted_at timestamp without time zone default null
);

---- create above / drop below ----

drop table books.notes;
drop table books.books;
drop table books.users;
