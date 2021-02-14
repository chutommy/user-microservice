-- name: CreateUser :one
insert into users (username, hashed_password, first_name, last_name, birth_day, gender, email, phone_number)
values ($1, $2, $3, $4, $5, $6, $7, $8)
returning *;

-- name: GetUserByID :one
select id,
       username,
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
select id,
       username,
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
select id,
       username,
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

-- name: UpdateUserUsername :one
update users
set username = $2
where id = $1
returning *;

-- name: UpdateUserEmail :one
update users
set email = $2
where id = $1
returning *;

-- name: UpdateUserPhoneNumber :one
update users
set phone_number = $2
where id = $1
returning *;

-- name: UpdateUserPassword :one
update users
set hashed_password = $2
where id = $1
returning *;

-- name: UpdateUserInfo :one
update users
set first_name = $2 and last_name = $3 and birth_day = $4 and gender = $5
where id = $1
returning *;

-- name: GetHashedPassword :one
select hashed_password
from users
where id = $1
limit 1;
