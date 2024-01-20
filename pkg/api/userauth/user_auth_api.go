package apipkg

import (
	"context"
	"errors"
	"unicode"

	AwesomeExpenseTrackerApi "github.com/shashankmahajan99/awesome-expense-tracker-backend/api"
	db "github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/db/sqlc"
	"golang.org/x/crypto/bcrypt"
)

// UserAuthServer is the server API for UserAuthentication service.
type UserAuthServer interface {
	// Login logs in a user.
	Login(context.Context, *AwesomeExpenseTrackerApi.LoginUserRequest) (*AwesomeExpenseTrackerApi.LoginUserResponse, error)
	// Register registers a user.
	Register(context.Context, *AwesomeExpenseTrackerApi.RegisterUserRequest) (*AwesomeExpenseTrackerApi.RegisterUserResponse, error)
}

// LoginUser logs in a user.
func (s *Server) LoginUser(ctx context.Context, req *AwesomeExpenseTrackerApi.LoginUserRequest) (res *AwesomeExpenseTrackerApi.LoginUserResponse, err error) {
	// Add user login logic here
	err = s.validateLoginRequest(req)
	if err != nil {
		return nil, err
	}

	getUserResult, err := s.store.ListUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if getUserResult.Username == "" {
		return nil, errors.New("username doesn't exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(getUserResult.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("password is incorrect")
	}

	res = &AwesomeExpenseTrackerApi.LoginUserResponse{}
	res.Token = "token"
	return res, nil
}

// RegisterUser registers a user.
func (s *Server) RegisterUser(ctx context.Context, req *AwesomeExpenseTrackerApi.RegisterUserRequest) (res *AwesomeExpenseTrackerApi.RegisterUserResponse, err error) {
	// Add user registration logic here
	// validate registration details
	err = s.validateRegisterRequest(req)
	if err != nil {
		return nil, err
	}

	getUserResult, err := s.store.ListUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if getUserResult.Username != "" {
		return nil, errors.New("username already exists")
	}

	// Check if the username already exists in the database
	createUserParams := db.CreateUserParams{}
	createUserParams.Username = req.Username
	createUserParams.Email = req.Email
	createUserParams.Password = req.Password

	result, err := s.store.RegisterUser(ctx, createUserParams)
	if err != nil {
		return nil, err
	}
	res = &AwesomeExpenseTrackerApi.RegisterUserResponse{}
	res.Username = result.Username
	res.Email = result.Email

	return res, nil
}

// DeleteUser deletes a user.
func (s *Server) DeleteUser(ctx context.Context, req *AwesomeExpenseTrackerApi.DeleteUserRequest) (res *AwesomeExpenseTrackerApi.DeleteUserResponse, err error) {
	// Add user deletion logic here
	user, err := s.store.ListUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err

	}

	if user.Username == "" {
		return nil, errors.New("username doesn't exist")
	}

	err = s.store.DeleteUser(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	res = &AwesomeExpenseTrackerApi.DeleteUserResponse{}
	res.Username = req.Username

	return res, nil
}

func (s *Server) validateLoginRequest(req *AwesomeExpenseTrackerApi.LoginUserRequest) error {
	if req.Username == "" || req.Password == "" {
		return errors.New("username or password cannot be empty")
	}

	return nil
}

func (s *Server) validateRegisterRequest(req *AwesomeExpenseTrackerApi.RegisterUserRequest) error {
	if req.Username == "" || req.Password == "" {
		return errors.New("username or password cannot be empty")
	}

	if req.Email == "" {
		return errors.New("email cannot be empty")
	}

	if req.ConfirmPassword == "" || req.Password != req.ConfirmPassword {
		return errors.New("passwords do not match")
	}

	// Check if the password is strong enough
	if !isStrongPassword(req.Password) {
		return errors.New("password is not strong enough")
	}

	return nil
}

func isStrongPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasUppercase := false
	hasLowercase := false
	hasNumber := false
	hasSpecialChar := false

	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUppercase = true
		} else if unicode.IsLower(char) {
			hasLowercase = true
		} else if unicode.IsNumber(char) {
			hasNumber = true
		} else if unicode.IsPunct(char) || unicode.IsSymbol(char) {
			hasSpecialChar = true
		}
	}

	return hasUppercase && hasLowercase && hasNumber && hasSpecialChar
}
