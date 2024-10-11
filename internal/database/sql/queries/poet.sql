
-- name: CreatePoet :one
INSERT INTO poets (name, bio, time_period) VALUES ($1, $2, $3) RETURNING *;
-- like INSERT INTO users (name, password, roles) VALUES ('nameasf', 'sfaasffas', ARRAY['DBA']::role[]) RETURNING *;

-- name: GetPoets :many
SELECT * FROM poets;

-- name: GetPoetById :one
SELECT * FROM poets WHERE id = $1 LIMIT 1;

-- -- name: UpdatePoet :exec
-- UPDATE users SET
--   name = COALESCE(NULLIF($2::varchar, ''), name),
--   password = COALESCE(NULLIF($3::varchar, ''), password),
--   roles = COALESCE(NULLIF($4::role[], ARRAY[]::role[]), roles),
--   update_at = CURRENT_TIMESTAMP
-- WHERE id = $1;

-- -- name: DeleteUser :exec
-- DELETE FROM users WHERE id = $1;