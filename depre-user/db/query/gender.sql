-- name: CreateGender :one
insert into genders (title)
values (@title)
returning *;

-- name: GetGender :one
select *
from genders
where id = @id
limit 1;

-- name: ListGenders :many
select *
from genders;

-- name: DeleteGender :exec
delete
from genders
where id = @id;
