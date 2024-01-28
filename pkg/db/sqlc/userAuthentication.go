package db

import (
	"context"
	"database/sql"
)

// RegisterUserResult is the result of registering a new user.
type RegisterUserResult struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// UpdateUserResult is the result of updating a user.
type UpdateUserResult struct {
	Email          string `json:"email"`
	Name           string `json:"name"`
	Bio            string `json:"bio"`
	ProfilePicture string `json:"profile_picture"`
}

// RegisterUser registers a new user in the database.
func (store *MySQLStore) RegisterUser(ctx context.Context, arg CreateUserParams) (RegisterUserResult, error) {
	var userID int32
	registerUserResult := RegisterUserResult{}

	err := store.execTx(ctx, func(q *Queries) error {
		res, err := q.CreateUser(ctx, arg)
		if err != nil {
			return err
		}
		lastInsertID, err := res.LastInsertId()
		if err != nil {
			return err
		}
		userID = int32(lastInsertID)
		return nil
	})
	if err != nil {
		return registerUserResult, err
	}
	result, err := store.GetUserByID(ctx, int32(userID))
	if err != nil {
		return registerUserResult, err
	}
	registerUserResult.Username = result.Username
	registerUserResult.Email = result.Email
	return registerUserResult, nil
}

// DeleteUser deletes a user from the database.
func (store *MySQLStore) DeleteUser(ctx context.Context, username string) (int64, error) {
	numberOfRows := int64(0)
	err := store.execTx(ctx, func(q *Queries) error {
		numRows, err := q.DeleteUser(ctx, username)
		if err != nil {
			return err
		}
		numberOfRows = numRows
		return nil
	})
	return numberOfRows, err
}

// ListUserByUsername gets a user from the database by username.
func (store *MySQLStore) ListUserByUsername(ctx context.Context, username string) (User, error) {
	user, err := store.GetUserByUsername(ctx, username)
	if err == sql.ErrNoRows {
		// No rows found, return empty User and no error
		return user, nil
	}

	return user, err
}

// ListUserByEmail gets a user from the database by email.
func (store *MySQLStore) ListUserByEmail(ctx context.Context, email string) (User, error) {
	user, err := store.GetUserByEmail(ctx, email)
	if err == sql.ErrNoRows {
		// No rows found, return empty User and no error
		return user, nil
	}

	return user, err
}

// ModifyUserUsername updates a user's username in the database.
func (store *MySQLStore) ModifyUserUsername(ctx context.Context, arg UpdateUserUsernameParams) error {
	err := store.execTx(ctx, func(q *Queries) error {
		_, err := q.UpdateUserUsername(ctx, arg)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// ModifyUserPassword updates a user's password in the database.
func (store *MySQLStore) ModifyUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error {
	err := store.execTx(ctx, func(q *Queries) error {
		_, err := q.UpdateUserPassword(ctx, arg)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
