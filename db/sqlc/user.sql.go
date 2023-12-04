// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (username,email,hashed_password)
VALUES ($1,$2,$3)
RETURNING id, username, email, hashed_password, password_changed_at, created_at
`

type CreateUserParams struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Username, arg.Email, arg.HashedPassword)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, username, email, hashed_password, password_changed_at, created_at FROM users
WHERE id = $1 
LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, username, email, hashed_password, password_changed_at, created_at FROM users
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.HashedPassword,
			&i.PasswordChangedAt,
			&i.CreatedAt,
		); err != nil {
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

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET username= $2,
email = $3,
hashed_password = $4
WHERE id = $1
RETURNING id, username, email, hashed_password, password_changed_at, created_at
`

type UpdateUserParams struct {
	ID             int64  `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.ID,
		arg.Username,
		arg.Email,
		arg.HashedPassword,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}
