package db

import (
	"context"
)

// AddProfile creates a new user profile.
func (store *MySQLStore) AddProfile(ctx context.Context, arg CreateProfileParams) (Profile, error) {
	var result Profile
	var profileID int32
	err := store.execTx(ctx, func(q *Queries) error {
		res, err := q.CreateProfile(ctx, arg)
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
	var profileID int32
	var result Profile
	err := store.execTx(ctx, func(q *Queries) error {
		res, err := q.UpdateProfileBio(ctx, arg)
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
	result, err = store.GetProfileByID(ctx, int32(profileID))
	if err != nil {
		return result, err
	}

	return result, nil
}

// ModifyProfilePicture updates a user profile.
func (store *MySQLStore) ModifyProfilePicture(ctx context.Context, arg UpdateProfileProfilePictureParams) (Profile, error) {
	var profileID int32
	var result Profile
	err := store.execTx(ctx, func(q *Queries) error {
		res, err := q.UpdateProfileProfilePicture(ctx, arg)
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
	result, err = store.GetProfileByID(ctx, int32(profileID))
	if err != nil {
		return result, err
	}

	return result, nil
}

// ModifyProfileName updates a user profile.
func (store *MySQLStore) ModifyProfileName(ctx context.Context, arg UpdateProfileNameParams) (Profile, error) {
	var profileID int32
	var result Profile
	err := store.execTx(ctx, func(q *Queries) error {
		res, err := q.UpdateProfileName(ctx, arg)
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
	result, err = store.GetProfileByID(ctx, int32(profileID))
	if err != nil {
		return result, err
	}

	return result, nil
}
