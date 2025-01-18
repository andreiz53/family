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