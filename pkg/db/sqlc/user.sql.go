// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const countUsers = `-- name: CountUsers :one
SELECT count(*) FROM users
`

func (q *Queries) CountUsers(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countUsers)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createUser = `-- name: CreateUser :execresult
INSERT INTO users (
  username,
  password,
  email
) VALUES (
  ?, ?, ?
)
`

type CreateUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser, arg.Username, arg.Password, arg.Email)
}

const deleteUser = `-- name: DeleteUser :execrows
DELETE FROM users
WHERE username = ?
`

func (q *Queries) DeleteUser(ctx context.Context, username string) (int64, error) {
	result, err := q.db.ExecContext(ctx, deleteUser, username)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, username, password, email, created_at, updated_at FROM users
WHERE email = ?
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, username, password, email, created_at, updated_at FROM users
WHERE id = ?
`

func (q *Queries) GetUserByID(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, password, email, created_at, updated_at FROM users
WHERE username = ?
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, username, password, email, created_at, updated_at FROM users
ORDER BY id
LIMIT ?
OFFSET ?
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
			&i.Password,
			&i.Email,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateUserPassword = `-- name: UpdateUserPassword :execresult
UPDATE users
SET
  password = ?
WHERE email = ?
`

type UpdateUserPasswordParams struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateUserPassword, arg.Password, arg.Email)
}

const updateUserUsername = `-- name: UpdateUserUsername :execresult
UPDATE users
SET
  username = ?
WHERE email = ?
`

type UpdateUserUsernameParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (q *Queries) UpdateUserUsername(ctx context.Context, arg UpdateUserUsernameParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateUserUsername, arg.Username, arg.Email)
}
