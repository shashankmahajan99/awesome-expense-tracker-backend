package db

import (
	"context"
	"database/sql"
)

// ListProfileByEmail gets a user profile by given email.
func (store *MySQLStore) ListProfileByEmail(ctx context.Context, email string) (Profile, error) {
	var result Profile
	var err error
	res, err := store.GetProfileByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			// No rows found, return empty Profile and no error
			return result, nil
		}
		return result, err
	}
	result.Bio = res.Bio
	result.ProfileName = res.ProfileName
	result.ProfilePicture = res.ProfilePicture
	result.UserID = res.UserID
	result.CreatedAt = res.CreatedAt
	result.UpdatedAt = res.UpdatedAt

	return result, nil
}

// AddProfile creates a new user profile.
func (store *MySQLStore) AddProfile(ctx context.Context, arg CreateProfileParams) (Profile, error) {
	var result Profile
	var profileID int32
	err := store.execTx(ctx, func(q *Queries) error {
		res, err := q.CreateProfile(ctx, arg)
		if err != nil {
			return err
		}
		lastInsertID, err := res.LastInsertId()
		if err != nil {
			return err
		}
		profileID = int32(lastInsertID)
		return nil
	})
	if err != nil {
		return result, err
	}
	result, err = store.GetProfileByID(ctx, profileID)
	if err != nil {
		return result, err
	}

	return result, nil
}

// ModifyProfileBio updates a user profile.
func (store *MySQLStore) ModifyProfileBio(ctx context.Context, arg UpdateProfileBioParams) (Profile, error) {
	var result Profile
	err := store.execTx(ctx, func(q *Queries) error {
		res, err := q.UpdateProfileBio(ctx, arg)
		if err != nil {
			return err
		}
		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return nil
		}
		return nil
	})
	if err != nil {
		return result, err
	}
	res, err := store.GetProfileByEmail(ctx, arg.Email)
	if err != nil {
		return result, err
	}
	result.Bio = res.Bio
	result.ProfileName = res.ProfileName
	result.ProfilePicture = res.ProfilePicture
	result.UserID = res.UserID
	result.CreatedAt = res.CreatedAt
	result.UpdatedAt = res.UpdatedAt

	return result, nil
}

// ModifyProfilePicture updates a user profile.
func (store *MySQLStore) ModifyProfilePicture(ctx context.Context, arg UpdateProfileProfilePictureParams) (Profile, error) {
	var result Profile
	err := store.execTx(ctx, func(q *Queries) error {
		res, err := q.UpdateProfileProfilePicture(ctx, arg)
		if err != nil {
			return err
		}
		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return nil
		}
		return nil
	})
	if err != nil {
		return result, err
	}
	res, err := store.GetProfileByEmail(ctx, arg.Email)
	if err != nil {
		return result, err
	}
	result.Bio = res.Bio
	result.ProfileName = res.ProfileName
	result.ProfilePicture = res.ProfilePicture
	result.UserID = res.UserID
	result.CreatedAt = res.CreatedAt
	result.UpdatedAt = res.UpdatedAt

	return result, nil
}

// ModifyProfileName updates a user profile.
func (store *MySQLStore) ModifyProfileName(ctx context.Context, arg UpdateProfileNameParams) (Profile, error) {
	var result Profile
	err := store.execTx(ctx, func(q *Queries) error {
		res, err := q.UpdateProfileName(ctx, arg)
		if err != nil {
			return err
		}
		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return nil
		}
		return nil
	})
	if err != nil {
		return result, err
	}
	res, err := store.GetProfileByEmail(ctx, arg.Email)
	if err != nil {
		return result, err
	}
	result.Bio = res.Bio
	result.ProfileName = res.ProfileName
	result.ProfilePicture = res.ProfilePicture
	result.UserID = res.UserID
	result.CreatedAt = res.CreatedAt
	result.UpdatedAt = res.UpdatedAt

	return result, nil
}
