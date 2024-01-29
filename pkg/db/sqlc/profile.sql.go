// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: profile.sql

package db

import (
	"context"
	"database/sql"
)

const countprofiles = `-- name: Countprofiles :one
SELECT count(*) 
FROM profiles
`

func (q *Queries) Countprofiles(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countprofiles)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createProfile = `-- name: CreateProfile :execresult
INSERT INTO profiles (
  user_id,
  bio,
  profile_name,
  profile_picture
) VALUES (
  (SELECT id FROM Users WHERE email = ?), ?, ?, ?
)
`

type CreateProfileParams struct {
	Email          string `json:"email"`
	Bio            string `json:"bio"`
	ProfileName    string `json:"profile_name"`
	ProfilePicture string `json:"profile_picture"`
}

func (q *Queries) CreateProfile(ctx context.Context, arg CreateProfileParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createProfile,
		arg.Email,
		arg.Bio,
		arg.ProfileName,
		arg.ProfilePicture,
	)
}

const deleteProfile = `-- name: DeleteProfile :exec
DELETE FROM profiles
WHERE user_id = (SELECT id FROM Users WHERE email = ?)
`

func (q *Queries) DeleteProfile(ctx context.Context, email string) error {
	_, err := q.db.ExecContext(ctx, deleteProfile, email)
	return err
}

const getProfileByEmail = `-- name: GetProfileByEmail :one
SELECT profiles.id, profiles.user_id, profiles.bio, profiles.profile_name, profiles.profile_picture, profiles.created_at, profiles.updated_at
FROM profiles
JOIN Users ON profiles.user_id = Users.id
WHERE Users.email = ?
`

type GetProfileByEmailRow struct {
	Profile Profile `json:"profile"`
}

func (q *Queries) GetProfileByEmail(ctx context.Context, email string) (GetProfileByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, getProfileByEmail, email)
	var i GetProfileByEmailRow
	err := row.Scan(
		&i.Profile.ID,
		&i.Profile.UserID,
		&i.Profile.Bio,
		&i.Profile.ProfileName,
		&i.Profile.ProfilePicture,
		&i.Profile.CreatedAt,
		&i.Profile.UpdatedAt,
	)
	return i, err
}

const getProfileByID = `-- name: GetProfileByID :one
SELECT id, user_id, bio, profile_name, profile_picture, created_at, updated_at
FROM profiles
WHERE id = ?
`

func (q *Queries) GetProfileByID(ctx context.Context, id int32) (Profile, error) {
	row := q.db.QueryRowContext(ctx, getProfileByID, id)
	var i Profile
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Bio,
		&i.ProfileName,
		&i.ProfilePicture,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listprofiles = `-- name: Listprofiles :many
SELECT id, user_id, bio, profile_name, profile_picture, created_at, updated_at 
FROM profiles
ORDER BY id
LIMIT ?
OFFSET ?
`

type ListprofilesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) Listprofiles(ctx context.Context, arg ListprofilesParams) ([]Profile, error) {
	rows, err := q.db.QueryContext(ctx, listprofiles, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Profile{}
	for rows.Next() {
		var i Profile
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Bio,
			&i.ProfileName,
			&i.ProfilePicture,
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

const updateProfileBio = `-- name: UpdateProfileBio :execresult
UPDATE profiles
JOIN Users ON profiles.user_id = Users.id
SET profiles.bio = ?
WHERE Users.email = ?
`

type UpdateProfileBioParams struct {
	Bio   string `json:"bio"`
	Email string `json:"email"`
}

func (q *Queries) UpdateProfileBio(ctx context.Context, arg UpdateProfileBioParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateProfileBio, arg.Bio, arg.Email)
}

const updateProfileName = `-- name: UpdateProfileName :execresult
UPDATE profiles
JOIN Users ON profiles.user_id = Users.id
SET profiles.profile_name = ?
WHERE Users.email = ?
`

type UpdateProfileNameParams struct {
	ProfileName string `json:"profile_name"`
	Email       string `json:"email"`
}

func (q *Queries) UpdateProfileName(ctx context.Context, arg UpdateProfileNameParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateProfileName, arg.ProfileName, arg.Email)
}

const updateProfileProfilePicture = `-- name: UpdateProfileProfilePicture :execresult
UPDATE profiles
JOIN Users ON profiles.user_id = Users.id
SET profiles.profile_picture = ?
WHERE Users.email = ?
`

type UpdateProfileProfilePictureParams struct {
	ProfilePicture string `json:"profile_picture"`
	Email          string `json:"email"`
}

func (q *Queries) UpdateProfileProfilePicture(ctx context.Context, arg UpdateProfileProfilePictureParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateProfileProfilePicture, arg.ProfilePicture, arg.Email)
}
