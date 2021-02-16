-- name: CreateGender :one
insert into genders (title)
values ($1)
returning *;

-- name: GetGender :one
select *
from genders
where id = $1
limit 1;

-- name: ListGenders :many
select *
from genders;

-- name: DeleteGender :exec
delete
from genders
where id = $1;
