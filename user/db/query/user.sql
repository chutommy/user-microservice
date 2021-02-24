-- name: CreateUser :one
insert into users (email, hashed_password, first_name, last_name, birth_day, gender, phone_number)
values (@email, @hashed_password, @first_name, @last_name, @birth_day, @gender, @phone_number)
returning *;

-- name: GetUserByID :one
select *
from users
where id = @id
  and deleted_at is null
limit 1;

-- name: GetUserByEmail :one
select *
from users
where email = @email
  and deleted_at is null
limit 1;

-- name: UpdateUserEmail :one
update users
set email = @email
where id = @id
  and deleted_at is null
returning *;

-- name: UpdateUserPassword :one
update users
set hashed_password = @hashed_password
where id = @id
  and deleted_at is null
returning *;

-- name: UpdateUserInfo :one
update users
-- set first_name   = $2,
--     last_name    = $3,
--     birth_day    = $4,
--     gender       = $5,
--     phone_number = $6
set first_name   = case
                       when coalesce(@first_name, '') = '' then first_name
                       else @first_name
    end,
    last_name    = case
                       when coalesce(@last_name, '') = '' then last_name
                       else @last_name
        end,
    birth_day    = coalesce(@birth_day, birth_day),
    gender       = case
                       when coalesce(@gender, 0) = 0 then gender
                       else @gender
        end,
    phone_number = coalesce(@phone_number, phone_number)
where id = @id
  and deleted_at is null
returning *;

-- name: DeleteUserSoft :exec
update users
set deleted_at = now()
where id = @id
  and deleted_at is null;

-- name: RecoverUser :one
update users
set deleted_at = null
where id = @id
  and deleted_at is not null
returning *;

-- name: DeleteUserPermanent :exec
delete
from users
where id = @id
  and deleted_at is null;

-- name: GetHashedPassword :one
select hashed_password
from users
where id = @id
  and deleted_at is null
limit 1;
