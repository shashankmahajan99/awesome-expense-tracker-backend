package db

import (
	"context"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

// RegisterUserResult is the result of registering a new user.
type RegisterUserResult struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// RegisterUser registers a new user in the database.
func (store *MySQLStore) RegisterUser(ctx context.Context, arg CreateUserParams) (RegisterUserResult, error) {
	var userID int32
	registerUserResult := RegisterUserResult{}

	// Hash the password before storing it in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(arg.Password), bcrypt.DefaultCost)
	if err != nil {
		return registerUserResult, err
	}
	arg.Password = string(hashedPassword)

	err = store.execTx(ctx, func(q *Queries) error {
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
func (store *MySQLStore) DeleteUser(ctx context.Context, username string) error {
	return store.execTx(ctx, func(q *Queries) error {
		return q.DeleteUser(ctx, username)
	})
}

// ListUserByUsername gets a user from the database by username.
func (store *MySQLStore) ListUserByUsername(ctx context.Context, username string) (User, error) {
	user, err := store.GetUserByUsername(ctx, username)
	if err != nil {
		if err == sql.ErrNoRows {
			// No rows found, return empty User and no error
			return user, nil
		}
		// Some other error occurred
		return user, err
	}
	return user, nil
}
