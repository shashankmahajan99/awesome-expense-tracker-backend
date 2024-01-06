// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: profile.sql

package db

import (
	"context"
	"database/sql"
)

const countProfiles = `-- name: CountProfiles :one
SELECT count(*) FROM Profiles
`

func (q *Queries) CountProfiles(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countProfiles)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createProfile = `-- name: CreateProfile :execresult
INSERT INTO Profiles (
  user_id,
  bio
) VALUES (
  ?, ?
)
`

type CreateProfileParams struct {
	UserID sql.NullInt32  `json:"user_id"`
	Bio    sql.NullString `json:"bio"`
}

func (q *Queries) CreateProfile(ctx context.Context, arg CreateProfileParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createProfile, arg.UserID, arg.Bio)
}

const deleteProfile = `-- name: DeleteProfile :exec
DELETE FROM Profiles
WHERE id = ?
`

func (q *Queries) DeleteProfile(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteProfile, id)
	return err
}

const getProfile = `-- name: GetProfile :one
SELECT id, user_id, bio FROM Profiles
WHERE id = ?
`

func (q *Queries) GetProfile(ctx context.Context, id int32) (Profile, error) {
	row := q.db.QueryRowContext(ctx, getProfile, id)
	var i Profile
	err := row.Scan(&i.ID, &i.UserID, &i.Bio)
	return i, err
}

const listProfiles = `-- name: ListProfiles :many
SELECT id, user_id, bio FROM Profiles
ORDER BY id
LIMIT ?
OFFSET ?
`

type ListProfilesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListProfiles(ctx context.Context, arg ListProfilesParams) ([]Profile, error) {
	rows, err := q.db.QueryContext(ctx, listProfiles, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Profile{}
	for rows.Next() {
		var i Profile
		if err := rows.Scan(&i.ID, &i.UserID, &i.Bio); err != nil {
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

const updateProfile = `-- name: UpdateProfile :exec
UPDATE Profiles
SET
  user_id = ?,
  bio = ?
WHERE id = ?
`

type UpdateProfileParams struct {
	UserID sql.NullInt32  `json:"user_id"`
	Bio    sql.NullString `json:"bio"`
	ID     int32          `json:"id"`
}

func (q *Queries) UpdateProfile(ctx context.Context, arg UpdateProfileParams) error {
	_, err := q.db.ExecContext(ctx, updateProfile, arg.UserID, arg.Bio, arg.ID)
	return err
}