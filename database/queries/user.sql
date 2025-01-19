-- name: CreateUser :one
INSERT INTO users (
    email, password
) VALUES (
    $1, $2
)
RETURNING *;

-- name: GetUsers :many
SELECT * FROM users
WHERE deleted_at IS NULL;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1
AND deleted_at IS NULL;

-- name: GetUserByEmail :one
SELECT * from users
WHERE email = $1
AND deleted_at IS NULL;

-- name: DeleteUserByID :exec
UPDATE users
SET deleted_at = NOW()
WHERE id = $1;

-- name: DeleteUserByEmail :exec
UPDATE users
SET deleted_at = NOW()
WHERE email = $1;

-- name: UpdateUserEmail :exec
UPDATE users
SET email = $1
WHERE id = $2;

-- name: UpdateUserPassword :exec
UPDATE users
SET password = $1
WHERE id = $2;
