// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    email, password
) VALUES (
    $1, $2
)
RETURNING id, email, password, created_at, updated_at, deleted_at
`

type CreateUserParams struct {
	Email    string
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteUserByEmail = `-- name: DeleteUserByEmail :exec
UPDATE users
SET deleted_at = NOW()
WHERE email = $1
`

func (q *Queries) DeleteUserByEmail(ctx context.Context, email string) error {
	_, err := q.db.Exec(ctx, deleteUserByEmail, email)
	return err
}

const deleteUserByID = `-- name: DeleteUserByID :exec
UPDATE users
SET deleted_at = NOW()
WHERE id = $1
`

func (q *Queries) DeleteUserByID(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteUserByID, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, password, created_at, updated_at, deleted_at from users
WHERE email = $1
AND deleted_at IS NULL
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, email, password, created_at, updated_at, deleted_at FROM users
WHERE id = $1
AND deleted_at IS NULL
`

func (q *Queries) GetUserByID(ctx context.Context, id pgtype.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, email, password, created_at, updated_at, deleted_at FROM users
WHERE deleted_at IS NULL
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Password,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserEmail = `-- name: UpdateUserEmail :exec
UPDATE users
SET email = $1,
updated_at = NOW()
WHERE id = $2
`

type UpdateUserEmailParams struct {
	Email string
	ID    pgtype.UUID
}

func (q *Queries) UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) error {
	_, err := q.db.Exec(ctx, updateUserEmail, arg.Email, arg.ID)
	return err
}

const updateUserPassword = `-- name: UpdateUserPassword :exec
UPDATE users
SET password = $1,
updated_at = NOW()
WHERE id = $2
`

type UpdateUserPasswordParams struct {
	Password string
	ID       pgtype.UUID
}

func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error {
	_, err := q.db.Exec(ctx, updateUserPassword, arg.Password, arg.ID)
	return err
}
