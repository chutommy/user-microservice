-- name: CreateUser :one
insert into users (hashed_password, first_name, last_name, birth_day, gender, email, phone_number)
values ($1, $2, $3, $4, $5, $6, $7)
returning *;

-- name: GetUserByID :one
select *
from users
where id = $1
  and deleted_at IS NULL
limit 1;

-- name: GetUserByEmail :one
select *
from users
where email = $1
  and deleted_at IS NULL
limit 1;

-- name: UpdateUserEmail :one
update users
set email = $2
where id = $1
  and deleted_at IS NULL
returning *;

-- name: UpdateUserPhoneNumber :one
update users
set phone_number = $2
where id = $1
  and deleted_at IS NULL
returning *;

-- name: UpdateUserPassword :one
update users
set hashed_password = $2
where id = $1
  and deleted_at IS NULL
returning *;

-- name: UpdateUserInfo :one
update users
set first_name = $2,
    last_name  = $3,
    birth_day  = $4,
    gender     = $5
where id = $1
  and deleted_at IS NULL
returning *;

-- name: DeleteUserSoft :exec
update users
set deleted_at = now()
where id = $1
  and deleted_at IS NULL;

-- name: RecoverDeletedUser :one
update users
set deleted_at = null
where id = $1
  and deleted_at IS NOT NULL
returning *;

-- name: DeleteUserPermanent :exec
delete
from users
where id = $1
  and deleted_at IS NULL;

-- name: GetHashedPassword :one
select hashed_password
from users
where id = $1
  and deleted_at IS NULL
limit 1;
