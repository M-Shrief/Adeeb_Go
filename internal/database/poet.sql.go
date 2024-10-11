// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: poet.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createPoet = `-- name: CreatePoet :one
INSERT INTO poets (name, bio, time_period) VALUES ($1, $2, $3) RETURNING id, name, bio, reviewed, time_period, created_at, updated_at
`

type CreatePoetParams struct {
	Name       string     `json:"name"`
	Bio        string     `json:"bio"`
	TimePeriod TimePeriod `json:"time_period"`
}

func (q *Queries) CreatePoet(ctx context.Context, arg CreatePoetParams) (Poet, error) {
	row := q.db.QueryRow(ctx, createPoet, arg.Name, arg.Bio, arg.TimePeriod)
	var i Poet
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Bio,
		&i.Reviewed,
		&i.TimePeriod,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPoetById = `-- name: GetPoetById :one

SELECT id, name, bio, reviewed, time_period, created_at, updated_at FROM poets WHERE id = $1 LIMIT 1
`

// like INSERT INTO users (name, password, roles) VALUES ('nameasf', 'sfaasffas', ARRAY['DBA']::role[]) RETURNING *;
func (q *Queries) GetPoetById(ctx context.Context, id pgtype.UUID) (Poet, error) {
	row := q.db.QueryRow(ctx, getPoetById, id)
	var i Poet
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Bio,
		&i.Reviewed,
		&i.TimePeriod,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}