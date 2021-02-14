create table if not exists genders
(
    id    smallserial primary key,
    title varchar(32) not null
);

create table if not exists users
(
    id              bigserial primary key,
    username        varchar(64),
    hashed_password varchar     not null,
    first_name      varchar(64) not null,
    last_name       varchar(64) not null,
    birth_day       date,
    gender          smallint,
    email           varchar(64) not null,
    phone_number    varchar(32),
    updated_at      timestamptz,
    created_at      timestamptz not null default now()
);

alter table users
    add foreign key (gender) references genders (id);

create or replace function update_updated_at()
    returns trigger
    language plpgsql
as
$$
begin
    new.updated_at = now();
    return new;
end;
$$;

create trigger update_trigger
    after update
    on users
    for each row
execute procedure update_updated_at();
