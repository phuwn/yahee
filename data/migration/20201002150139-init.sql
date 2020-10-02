
-- +migrate Up
CREATE FUNCTION gen_random_uuid() RETURNS uuid
    LANGUAGE c
    AS '$libdir/pgcrypto', 'pg_random_uuid';

create table users (
	id uuid default gen_random_uuid() not null unique primary key,
	created_at timestamp(6) with time zone default now(),
	updated_at timestamp(6) with time zone,
	deleted_at timestamp(6) with time zone,
	name varchar(200),
	email varchar(200)
);

create table connections (
	user_id_1 uuid references users (id) on delete cascade not null,
	user_id_2 uuid references users (id) on delete cascade not null,
	created_at timestamp(6) with time zone default now(),
	updated_at timestamp(6) with time zone,
	deleted_at timestamp(6) with time zone,
	primary key (user_id_1, user_id_2)
);

create table messages (
	id uuid default gen_random_uuid() not null unique primary key,
	sender_id uuid references users (id) on delete cascade not null,
	receiver_id uuid references users (id) on delete cascade not null,
	created_at timestamp(6) with time zone default now(),
	updated_at timestamp(6) with time zone,
	deleted_at timestamp(6) with time zone,
	content text
);

-- +migrate Down
drop table messages;
drop table connections;
drop table users;
drop function gen_random_uuid;