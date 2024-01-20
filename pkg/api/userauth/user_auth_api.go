package apipkg

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
	"unicode"

	"github.com/golang-jwt/jwt/v5"
	AwesomeExpenseTrackerApi "github.com/shashankmahajan99/awesome-expense-tracker-backend/api"
	db "github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/db/sqlc"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": getUserResult.Username,
		"exp":      time.Now().Add(time.Hour).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(s.config.JwtKey)
	res = &AwesomeExpenseTrackerApi.LoginUserResponse{}
	res.AccessToken = tokenString
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

// AuthenticateWithGoogle authenticates a user with Google.
func (s *Server) AuthenticateWithGoogle(ctx context.Context, req *AwesomeExpenseTrackerApi.AuthenticateWithGoogleRequest) (res *AwesomeExpenseTrackerApi.AuthenticateWithGoogleResponse, err error) {
	// Add user authentication with Google logic here
	url := s.config.GcpOAuthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	res = &AwesomeExpenseTrackerApi.AuthenticateWithGoogleResponse{}
	res.Url = url
	return res, nil
}

// AuthenticateWithGoogleCallback authenticates a user with Google.
func (s *Server) AuthenticateWithGoogleCallback(ctx context.Context, req *AwesomeExpenseTrackerApi.AuthenticateWithGoogleCallbackRequest) (res *AwesomeExpenseTrackerApi.AuthenticateWithGoogleCallbackResponse, err error) {
	// Add user authentication with Google callback logic here
	token, err := s.config.GcpOAuthConfig.Exchange(ctx, req.Code)
	if err != nil {
		return nil, err
	}

	res = &AwesomeExpenseTrackerApi.AuthenticateWithGoogleCallbackResponse{}
	res.AccessToken = token.AccessToken
	res.RefreshToken = token.RefreshToken
	res.ExpiresAt = token.Expiry.String()
	res.IdToken = token.Extra("id_token").(string)
	res.TokenType = token.TokenType

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	res.Email = data["email"].(string)
	res.Name = data["name"].(string)
	res.ProfilePic = data["picture"].(string)

	return res, nil
}

// Write all utility functions here

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
