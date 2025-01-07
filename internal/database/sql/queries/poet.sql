
-- name: CreatePoet :one
INSERT INTO poets (name, bio, time_period) VALUES (sqlc.narg('name'), sqlc.narg('bio'), sqlc.narg('time_period')) RETURNING id,name,bio,time_period;

-- name: GetPoets :many
SELECT id,name,bio,time_period FROM poets;

-- name: GetPoetById :one
SELECT id,name,bio,time_period FROM poets WHERE id = $1 LIMIT 1;

-- -- name: DeletePoet :exec
-- DELETE FROM poets WHERE id = $1;