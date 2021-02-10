create table genders
(
    id    smallserial primary key,
    title varchar(32)
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

create table if not exists companies
(
    id           bigserial primary key,
    owner_id     bigint      not null,
    title        varchar(64),
    phone_number varchar(64) not null,
    email        varchar(64) not null,
    location     varchar(128),
    description  text,
    website      text,
    work_from    time,
    work_to      time,
    created_at   timestamptz not null default now()
);

alter table companies
    add foreign key (owner_id) references users (id);

create table if not exists clients
(
    client_id  bigint      not null,
    company_id bigint      not null,
    created_at timestamptz not null default now(),
    primary key (client_id, company_id)
);

alter table clients
    add foreign key (client_id) references users (id);

alter table clients
    add foreign key (company_id) references companies (id);

create table if not exists employees
(
    employee_id bigint      not null,
    company_id  bigint      not null,
    created_at  timestamptz not null default now(),
    deleted_at  timestamptz,
    primary key (employee_id, company_id)
);

alter table employees
    add foreign key (employee_id) references users (id);

alter table employees
    add foreign key (company_id) references companies (id);

create table if not exists services
(
    id            bigserial primary key,
    company_id    bigint      not null,
    title         varchar(64) not null,
    description   text,
    price         numeric(14, 2),
    time_duration interval,
    created_at    timestamptz not null default now(),
    deleted_at    timestamptz
);

alter table services
    add foreign key (company_id) references companies (id);

create table if not exists conversations
(
    id         bigserial primary key,
    user_id    bigint      not null,
    company_id bigint      not null,
    created_at timestamptz not null default now()
);

alter table conversations
    add foreign key (user_id) references users (id);

alter table conversations
    add foreign key (company_id) references companies (id);

create table if not exists messages
(
    id              bigserial primary key,
    conversation_id bigint      not null,
    sender_id       bigint      not null,
    receiver_id     bigint      not null,
    message         text        not null,
    status          varchar(16) not null,
    posted_at       timestamptz not null default now()
);

alter table messages
    add foreign key (conversation_id) references conversations (id);

alter table messages
    add foreign key (sender_id) references users (id);

alter table messages
    add foreign key (receiver_id) references users (id);

create table if not exists reservations
(
    id             bigserial primary key,
    user_id        bigint      not null,
    service_id     bigint      not null,
    status         varchar(32) not null,
    reservation_at timestamptz not null,
    created_at     timestamptz not null default now(),
    canceled_at    timestamptz
);

alter table reservations
    add foreign key (user_id) references users (id);

alter table reservations
    add foreign key (service_id) references companies (id);

create table if not exists reviews
(
    id              bigserial primary key,
    reservation_id  bigint not null,
    company_id      bigint not null,
    text            text   not null,
    rate            decimal(2, 2),
    company_comment text
);

alter table reviews
    add foreign key (reservation_id) references reservations (id);

alter table reviews
    add foreign key (company_id) references companies (id);
