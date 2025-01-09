// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: users.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (name,password, roles) VALUES ($1, $2, $3) RETURNING id,name,roles
`

type CreateUserParams struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Roles    []Role `json:"roles"`
}

type CreateUserRow struct {
	ID    pgtype.UUID `json:"id"`
	Name  string      `json:"name"`
	Roles []Role      `json:"roles"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Name, arg.Password, arg.Roles)
	var i CreateUserRow
	err := row.Scan(&i.ID, &i.Name, &i.Roles)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getUserByName = `-- name: GetUserByName :one

SELECT id, name, password, roles FROM users WHERE name = $1 LIMIT 1
`

type GetUserByNameRow struct {
	ID       pgtype.UUID `json:"id"`
	Name     string      `json:"name"`
	Password string      `json:"password"`
	Roles    []Role      `json:"roles"`
}

// like INSERT INTO users (name, password, roles) VALUES ('nameasf', 'sfaasffas', ARRAY['DBA']::role[]) RETURNING *;
func (q *Queries) GetUserByName(ctx context.Context, name string) (GetUserByNameRow, error) {
	row := q.db.QueryRow(ctx, getUserByName, name)
	var i GetUserByNameRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Password,
		&i.Roles,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users SET
  name = COALESCE(NULLIF($2::varchar, ''), name),
  password = COALESCE(NULLIF($3::varchar, ''), password),
  roles = COALESCE(NULLIF($4::role[], ARRAY[]::role[]), roles),
  updated_at = CURRENT_TIMESTAMP
WHERE id = $1
`

type UpdateUserParams struct {
	ID      pgtype.UUID `json:"id"`
	Column2 string      `json:"column_2"`
	Column3 string      `json:"column_3"`
	Column4 []Role      `json:"column_4"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser,
		arg.ID,
		arg.Column2,
		arg.Column3,
		arg.Column4,
	)
	return err
}