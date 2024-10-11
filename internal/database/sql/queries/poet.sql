
-- name: CreatePoet :one
INSERT INTO poets (name, bio, time_period) VALUES ($1, $2, $3) RETURNING id,name,bio,time_period;

-- name: GetPoets :many
SELECT id,name,bio,time_period FROM poets;

-- name: GetPoetById :one
SELECT id,name,bio,time_period FROM poets WHERE id = $1 LIMIT 1;

-- -- name: UpdatePoet :exec
-- UPDATE poets SET
--   name = COALESCE(NULLIF($2::varchar, ''), name),
--   password = COALESCE(NULLIF($3::varchar, ''), password),
--   roles = COALESCE(NULLIF($4::role[], ARRAY[]::role[]), roles),
--   update_at = CURRENT_TIMESTAMP
-- WHERE id = $1;

-- -- name: DeletePoet :exec
-- DELETE FROM poets WHERE id = $1;