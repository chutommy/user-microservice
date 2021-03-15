create table if not exists users
(
    id              uuid primary key,
    email           varchar(64) not null unique,
    phone_number    varchar(32) unique,
    hashed_password varchar     not null,
    first_name      varchar(64) not null,
    last_name       varchar(64) not null,
    gender          smallint,
    birth_day       date,
    updated_at      timestamptz,
    created_at      timestamptz not null default now()
);

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
    before update
    on users
    for each row
execute procedure update_updated_at();
