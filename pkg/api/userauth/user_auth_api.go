package apipkg

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
	"unicode"

	"github.com/bufbuild/protovalidate-go"
	"github.com/golang-jwt/jwt/v5"
	AwesomeExpenseTrackerApi "github.com/shashankmahajan99/awesome-expense-tracker-backend/api"
	db "github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/db/sqlc"
	"github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserAuthServer is the server API for UserAuthentication service.
type UserAuthServer interface {
	// Login logs in a user.
	Login(context.Context, *AwesomeExpenseTrackerApi.LoginUserRequest) (*AwesomeExpenseTrackerApi.OAuth2Token, error)

	// Register registers a user.
	Register(context.Context, *AwesomeExpenseTrackerApi.RegisterUserRequest) (*AwesomeExpenseTrackerApi.OAuth2Token, error)

	// Delete deletes a user.
	DeleteUser(ctx context.Context, req *AwesomeExpenseTrackerApi.DeleteUserRequest) (res *AwesomeExpenseTrackerApi.DeleteUserResponse, err error)

	// AuthenticateWithGoogle authenticates a user with Google.
	AuthenticateWithGoogle(_ context.Context, _ *AwesomeExpenseTrackerApi.AuthenticateWithGoogleRequest) (res *AwesomeExpenseTrackerApi.AuthenticateWithGoogleResponse, err error)

	// AuthenticateWithGoogleCallback authenticates a user with Google.
	AuthenticateWithGoogleCallback(ctx context.Context, req *AwesomeExpenseTrackerApi.AuthenticateWithGoogleCallbackRequest) (res *AwesomeExpenseTrackerApi.OAuth2Token, err error)
}

// LoginUser logs in a user.
func (s *Server) LoginUser(ctx context.Context, req *AwesomeExpenseTrackerApi.LoginUserRequest) (res *AwesomeExpenseTrackerApi.OAuth2Token, err error) {
	// validate login details
	v, err := protovalidate.New()
	if err != nil {
		return nil, errors.New("failed to initialize validator: " + err.Error())
	}

	if err = v.Validate(req); err != nil {
		return nil, errors.New("failed to validate request: " + err.Error())
	}

	// Check if the username already exists in the database
	getUserResult, err := s.store.ListUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if getUserResult.Username == "" {
		return nil, errors.New("username doesn't exist")
	}

	// Check if the password is correct
	err = bcrypt.CompareHashAndPassword([]byte(getUserResult.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("password is incorrect")
	}

	// Generate JWT token
	token, err := s.generateJWTToken(req.Username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token: %v", err)
	}

	// Parse the token and return the response
	res = &AwesomeExpenseTrackerApi.OAuth2Token{}
	res, err = s.oauthTokenParser(res, token)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// RegisterUser registers a user.
func (s *Server) RegisterUser(ctx context.Context, req *AwesomeExpenseTrackerApi.RegisterUserRequest) (res *AwesomeExpenseTrackerApi.OAuth2Token, err error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, errors.New("failed to initialize validator: " + err.Error())
	}

	if err = v.Validate(req); err != nil {
		return nil, errors.New("failed to validate request: " + err.Error())
	}

	// validate registration details
	err = s.validateRegisterRequest(req)
	if err != nil {
		return nil, err
	}

	// Check if the username already exists in the database
	getUserResult, err := s.store.ListUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if getUserResult.Username != "" {
		return nil, errors.New("username already exists")
	}

	createUserParams := db.CreateUserParams{}
	createUserParams.Username = req.Username
	createUserParams.Email = req.Email

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	createUserParams.Password = string(hashedPassword)

	_, err = s.store.RegisterUser(ctx, createUserParams)
	if err != nil {
		return nil, err
	}
	token, err := s.generateJWTToken(req.Username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token: %v", err)
	}

	// Parse the token and return the response
	res = &AwesomeExpenseTrackerApi.OAuth2Token{}
	res, err = s.oauthTokenParser(res, token)
	if err != nil {
		return nil, err
	}

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
func (s *Server) AuthenticateWithGoogle(_ context.Context, _ *AwesomeExpenseTrackerApi.AuthenticateWithGoogleRequest) (res *AwesomeExpenseTrackerApi.AuthenticateWithGoogleResponse, err error) {
	// Add user authentication with Google logic here
	url := s.config.GcpOAuthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	res = &AwesomeExpenseTrackerApi.AuthenticateWithGoogleResponse{}
	res.Url = url
	return res, nil
}

// AuthenticateWithGoogleCallback authenticates a user with Google.
func (s *Server) AuthenticateWithGoogleCallback(ctx context.Context, req *AwesomeExpenseTrackerApi.AuthenticateWithGoogleCallbackRequest) (res *AwesomeExpenseTrackerApi.OAuth2Token, err error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, errors.New("failed to initialize validator: " + err.Error())
	}

	if err = v.Validate(req); err != nil {
		return nil, errors.New("failed to validate request: " + err.Error())
	}

	// Add user authentication with Google callback logic here
	token, err := s.config.GcpOAuthConfig.Exchange(ctx, req.Code)
	if err != nil {
		return nil, err
	}

	// Parse the token and return the response
	res = &AwesomeExpenseTrackerApi.OAuth2Token{}
	res, err = s.oauthTokenParser(res, token)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Write all utility functions here

func (s *Server) validateRegisterRequest(req *AwesomeExpenseTrackerApi.RegisterUserRequest) error {
	if req.Password != req.ConfirmPassword {
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

func (s *Server) generateJWTToken(username string) (*oauth2.Token, error) {
	// Create a new token object for the ID token
	idToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // ID token expires after 24 hours
	})

	// Sign and get the complete encoded token as a string using the secret
	idTokenString, err := idToken.SignedString(s.config.JwtKey)
	if err != nil {
		return nil, err
	}

	// Create a new token object for the access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Access token expires after 1 hour
	})

	// Sign and get the complete encoded token as a string using the secret
	accessTokenString, err := accessToken.SignedString(s.config.JwtKey)
	if err != nil {
		return nil, err
	}

	// Create a new token object for the refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // Refresh token expires after 72 hours
	})

	// Sign and get the complete encoded token as a string using the secret
	refreshTokenString, err := refreshToken.SignedString(s.config.JwtKey)
	if err != nil {
		return nil, err
	}

	token := &oauth2.Token{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
		Expiry:       time.Now().Add(time.Hour * 1),
		TokenType:    "Bearer",
	}

	token.WithExtra(map[string]interface{}{
		"id_token": idTokenString,
	})
	return token, nil
}

// oauthTokenParser parses the oauth2.Token (t) and returns the AwesomeExpenseTrackerApi.OAuth2Token (v)
func (s *Server) oauthTokenParser(v *AwesomeExpenseTrackerApi.OAuth2Token, t *oauth2.Token) (*AwesomeExpenseTrackerApi.OAuth2Token, error) {
	v.AccessToken = t.AccessToken
	v.RefreshToken = t.RefreshToken
	v.ExpiresAt = t.Expiry.String()
	v.IdToken = t.Extra("id_token").(string)
	v.TokenType = t.TokenType

	httpResponse, err := http.Get(utils.GoogleUserInfoURL + t.AccessToken)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	body, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	email, ok := data["email"].(string)
	if !ok {
		return nil, errors.New("email not found or not a string")
	}
	v.Email = email

	name, ok := data["name"].(string)
	if !ok {
		return nil, errors.New("name not found or not a string")
	}
	v.Name = name

	picture, ok := data["picture"].(string)
	if !ok {
		return nil, errors.New("profile picture not found or not a string")
	}
	v.ProfilePic = picture

	return v, nil
}
