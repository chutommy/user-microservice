--  this script run on the database's initialization.
-- create a function which updates the updated_at timestamp
create
or replace function trigger_set_timestamp() returns trigger as $ $ begin new.updated_at = now();
return new;
end;
$ $ language plpgsql;
-- create a table
create table if not exists accounts (
  id uuid primary key,
  username varchar(32) unique,
  email varchar(255) unique not null,
  phone varchar(33) unique,
  hpassword varchar(128) not null,
  first_name varchar(64),
  last_name varchar(64),
  birth_day date,
  perm_address varchar(255),
  mail_address varchar(255),
  created_at timestamp not null default now(),
  updated_at timestamp not null default now(),
  deleted_at timestamp default null
);
-- create a trigger
create trigger set_timestamp before
update
  on accounts for each row execute procedure trigger_set_timestamp();
