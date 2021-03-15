-- name: CreateUser :one
insert into users (id, email, phone_number, hashed_password, first_name, last_name, gender, birth_day)
values (@id, @email, @phone_number, @hashed_password, @first_name, @last_name, @gender, @birth_day)
returning *;

-- name: GetUser :one
select *
from users
where id = @id
limit 1;

-- name: UpdateUser :one
update users
set email           = case when coalesce(@email::varchar(64), '') = '' then email else @email end,
    phone_number    = case when coalesce(@phone_number::varchar(32), '') = '' then phone_number else @phone_number end,
    hashed_password = case
                          when coalesce(@hashed_password::varchar, '') = '' then hashed_password
                          else @hashed_password end,
    first_name      = case when coalesce(@first_name::varchar(64), '') = '' then first_name else @first_name end,
    last_name       = case when coalesce(@last_name::varchar(64), '') = '' then last_name else @last_name end,
    gender          = case when coalesce(@gender::smallint, 0) = 0 then gender else @gender end,
    birth_day       = case when @birth_day::date = '0001-01-01' then birth_day else @birth_day end
where id = @id
returning *;

-- name: DeleteUser :one
with deleted as
         (
             delete from users
                 where id = @id
                 returning *
         )
select count(*)
from deleted;
