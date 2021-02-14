// Code generated by sqlc. DO NOT EDIT.
// source: gender.sql

package repo

import (
	"context"
)

const createGender = `-- name: CreateGender :exec
insert into genders (title)
values ($1)
`

func (q *Queries) CreateGender(ctx context.Context, title string) error {
	_, err := q.db.ExecContext(ctx, createGender, title)
	return err
}

const deleteGender = `-- name: DeleteGender :exec
delete
from genders
where id = $1
`

func (q *Queries) DeleteGender(ctx context.Context, id int16) error {
	_, err := q.db.ExecContext(ctx, deleteGender, id)
	return err
}

const getGender = `-- name: GetGender :one
select id, title
from genders
where id = $1
   or title = $1
limit 1
`

func (q *Queries) GetGender(ctx context.Context, id int16) (Gender, error) {
	row := q.db.QueryRowContext(ctx, getGender, id)
	var i Gender
	err := row.Scan(&i.ID, &i.Title)
	return i, err
}

const listGenders = `-- name: ListGenders :many
select id, title
from genders
`

func (q *Queries) ListGenders(ctx context.Context) ([]Gender, error) {
	rows, err := q.db.QueryContext(ctx, listGenders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Gender{}
	for rows.Next() {
		var i Gender
		if err := rows.Scan(&i.ID, &i.Title); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
