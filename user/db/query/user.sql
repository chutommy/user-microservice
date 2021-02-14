-- name: CreateUser :one
insert into users (username, hashed_password, first_name, last_name, birth_day, gender, email, phone_number)
values ($1, $2, $3, $4, $5, $6, $7, $8)
returning *;

-- name: GetUserByID :one
select username,
       first_name,
       last_name,
       birth_day,
       gender,
       email,
       phone_number,
       updated_at,
       created_at
from users
where id = $1
limit 1;

-- name: GetUserByUsername :one
select username,
       first_name,
       last_name,
       birth_day,
       gender,
       email,
       phone_number,
       updated_at,
       created_at
from users
where username = $1
limit 1;

-- name: GetUserByEmail :one
select username,
       first_name,
       last_name,
       birth_day,
       gender,
       email,
       phone_number,
       updated_at,
       created_at
from users
where email = $1
limit 1;

-- name: UpdateUserUsername :exec
update users
set username = $2
where id = $1;

-- name: UpdateUserPassword :exec
update users
set hashed_password = $2
where id = $1;

-- name: UpdateUserInfo :exec
update users
set first_name = $2 and last_name = $3 and birth_day = $4 and gender = $5
where id = $1;
