-- name: createuser :one
insert into users (email, hashed_password, first_name, last_name, birth_day, gender, phone_number)
values ($1, $2, $3, $4, $5, $6, $7)
returning *;

-- name: getuserbyid :one
select *
from users
where id = $1
  and deleted_at is null
limit 1;

-- name: getuserbyemail :one
select *
from users
where email = $1
  and deleted_at is null
limit 1;

-- name: updateuseremail :one
update users
set email = $2
where id = $1
  and deleted_at is null
returning *;

-- name: updateuserpassword :one
update users
set hashed_password = $2
where id = $1
  and deleted_at is null
returning *;

-- name: updateuserinfo :one
update users
set first_name   = $2,
    last_name    = $3,
    birth_day    = $4,
    gender       = $5,
    phone_number = $6
where id = $1
  and deleted_at is null
returning *;

-- name: deleteusersoft :exec
update users
set deleted_at = now()
where id = $1
  and deleted_at is null;

-- name: recoverdeleteduser :one
update users
set deleted_at = null
where id = $1
  and deleted_at is not null
returning *;

-- name: deleteuserpermanent :exec
delete
from users
where id = $1
  and deleted_at is null;

-- name: gethashedpassword :one
select hashed_password
from users
where id = $1
  and deleted_at is null
limit 1;
