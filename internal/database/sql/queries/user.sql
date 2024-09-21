
-- name: CreateUser :one
INSERT INTO users (name,password, roles) VALUES ($1, $2, $3) RETURNING id,name,roles;
-- like INSERT INTO users (name, password, roles) VALUES ('nameasf', 'sfaasffas', ARRAY['DBA']::role[]) RETURNING *;

-- name: GetUserByName :one
SELECT id, name, password, roles FROM users WHERE name = $1 LIMIT 1;

-- name: UpdateUser :exec
UPDATE users SET
  name = COALESCE(NULLIF($2::varchar, ''), name),
  password = COALESCE(NULLIF($3::varchar, ''), password),
  roles = COALESCE(NULLIF($4::role[], ARRAY[]::role[]), roles),
  update_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
