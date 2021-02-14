-- name: CreateGender :exec
insert into genders (title)
values ($1);

-- name: GetGender :one
select *
from genders
where id = $1
   or title = $1
limit 1;

-- name: ListGenders :many
select *
from genders;

-- name: DeleteGender :exec
delete
from genders
where id = $1;
