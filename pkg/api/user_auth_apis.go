package apipkg

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/bufbuild/protovalidate-go"
	"github.com/golang-jwt/jwt/v5"
	AwesomeExpenseTrackerApi "github.com/shashankmahajan99/awesome-expense-tracker-backend/api"
	db "github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/db/sqlc"
	failuremanagement "github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/failure_management"
	"github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
)

type authenticateWithGoogleResponse struct {
	URL string `json:"url"`
}

// UserAuthServer is the server API for UserAuthentication service.
type UserAuthServer interface {
	// Login logs in a user.
	Login(context.Context, *AwesomeExpenseTrackerApi.LoginUserRequest) (*AwesomeExpenseTrackerApi.OAuth2Token, error)

	// Register registers a user.
	Register(context.Context, *AwesomeExpenseTrackerApi.RegisterUserRequest) (*AwesomeExpenseTrackerApi.OAuth2Token, error)

	// Delete deletes a user.
	DeleteUser(ctx context.Context, req *AwesomeExpenseTrackerApi.DeleteUserRequest) (res *AwesomeExpenseTrackerApi.DeleteUserResponse, err error)

	// AuthenticateWithGoogleCallback authenticates a user with Google.
	AuthenticateWithGoogleCallback(ctx context.Context, req *AwesomeExpenseTrackerApi.AuthenticateWithGoogleCallbackRequest) (res *AwesomeExpenseTrackerApi.OAuth2Token, err error)
}

// LoginUser logs in a user.
func (s *Server) LoginUser(ctx context.Context, req *AwesomeExpenseTrackerApi.LoginUserRequest) (res *AwesomeExpenseTrackerApi.OAuth2Token, err error) {
	// validate login details
	v, err := protovalidate.New()
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to initialize validator: "+err.Error(), http.StatusInternalServerError)
	}

	if err = v.Validate(req); err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "failed to validate request: "+err.Error(), http.StatusBadRequest)
	}

	if req.AuthProvider == utils.GoogleAuthProvider {
		gcpOauthRes, _ := s.authenticateWithGoogle(ctx)
		res = &AwesomeExpenseTrackerApi.OAuth2Token{}
		res.AuthUrl = gcpOauthRes.URL

		return res, nil
	}

	if req.Username == "" || req.Password == "" {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "invalid username or password", http.StatusBadRequest)
	}

	// Check if the username doesn't exists in the database
	getUserResult, err := s.store.ListUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "unknown error: "+err.Error(), http.StatusInternalServerError)
	}
	if getUserResult.Username == "" {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "user doesn't exist", http.StatusBadRequest)
	}

	// Check if the password is correct
	err = bcrypt.CompareHashAndPassword([]byte(getUserResult.Password), []byte(req.Password))
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "incorrect password", http.StatusBadRequest)
	}

	// Generate JWT token
	token, err := s.generateJWTToken(getUserResult.Email)
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "cannot generate access token: "+err.Error(), http.StatusInternalServerError)
	}

	// Parse the token and return the response
	res = &AwesomeExpenseTrackerApi.OAuth2Token{}
	res, err = s.oauthTokenParser(ctx, res, token)
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "cannot parse token: "+err.Error(), http.StatusInternalServerError)
	}
	res.Email = getUserResult.Email
	res.AuthProvider = req.AuthProvider

	return res, nil
}

// RegisterUser registers a user.
func (s *Server) RegisterUser(ctx context.Context, req *AwesomeExpenseTrackerApi.RegisterUserRequest) (res *AwesomeExpenseTrackerApi.OAuth2Token, err error) {
	res = &AwesomeExpenseTrackerApi.OAuth2Token{}
	// validate register details
	v, err := protovalidate.New()
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to initialize validator: "+err.Error(), http.StatusInternalServerError)
	}

	if err = v.Validate(req); err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "failed to validate request: "+err.Error(), http.StatusBadRequest)
	}

	// validate custom registration details
	if req.AuthProvider != utils.GoogleAuthProvider {
		err = s.validateRegisterRequest(req)
		if err != nil {
			return nil, failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "invalid registration details: "+err.Error(), http.StatusBadRequest)
		}
	}

	// Check if the email already exists in the database
	getUserResult, err := s.store.ListUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "Unknown error: "+err.Error(), http.StatusInternalServerError)
	}
	if getUserResult.Email != "" {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "email already exists", http.StatusBadRequest)
	}

	// Check if the username already exists in the database
	getUserResult, err = s.store.ListUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "Unknown error: "+err.Error(), http.StatusInternalServerError)
	}
	if getUserResult.Username != "" {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "username already exists", http.StatusBadRequest)
	}

	if req.AuthProvider == utils.GoogleAuthProvider {
		gcpOauthRes, _ := s.authenticateWithGoogle(ctx)

		res.AuthUrl = gcpOauthRes.URL
		return res, nil
	}
	createUserParams := db.CreateUserParams{}
	createUserParams.Username = req.Username
	createUserParams.Email = req.Email

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "incorrect password", http.StatusBadRequest)
	}
	createUserParams.Password = string(hashedPassword)

	_, err = s.store.RegisterUser(ctx, createUserParams)
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "Unknown error: "+err.Error(), http.StatusInternalServerError)
	}
	token, err := s.generateJWTToken(req.Email)
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "cannot generate access token: "+err.Error(), http.StatusInternalServerError)
	}

	// Parse the token and return the response
	res, err = s.oauthTokenParser(ctx, res, token)
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "cannot parse token: "+err.Error(), http.StatusInternalServerError)
	}
	res.Email = req.Email
	res.Name = req.Name
	res.ProfilePic = req.ProfilePic
	res.AuthProvider = req.AuthProvider
	return res, nil
}

// DeleteUser deletes a user.
func (s *Server) DeleteUser(ctx context.Context, req *AwesomeExpenseTrackerApi.DeleteUserRequest) (res *AwesomeExpenseTrackerApi.DeleteUserResponse, err error) {
	// validate delete details
	v, err := protovalidate.New()
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to initialize validator: "+err.Error(), http.StatusInternalServerError)
	}

	if err = v.Validate(req); err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "failed to validate request: "+err.Error(), http.StatusBadRequest)
	}
	// Add user deletion logic here
	user, err := s.store.ListUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to list user: "+err.Error(), http.StatusInternalServerError)
	}

	if user.Username == "" {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "username doesn't exist", http.StatusBadRequest)
	}

	err = s.store.DeleteUser(ctx, req.Username)
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "unknown error: "+err.Error(), http.StatusInternalServerError)
	}

	res = &AwesomeExpenseTrackerApi.DeleteUserResponse{}
	res.Username = req.Username

	return res, nil
}

// UpdateUser updates a user.
func (s *Server) UpdateUser(ctx context.Context, req *AwesomeExpenseTrackerApi.UpdateUserRequest) (res *AwesomeExpenseTrackerApi.UpdateUserResponse, err error) {
	res = &AwesomeExpenseTrackerApi.UpdateUserResponse{}

	v, err := protovalidate.New()
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to initialize validator: "+err.Error(), http.StatusInternalServerError)
	}

	if err = v.Validate(req); err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "failed to validate request: "+err.Error(), http.StatusBadRequest)
	}

	if req.NewPassword == "" && req.NewUsername == "" {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "no user was updated", http.StatusBadRequest)
	}

	getUserResult, err := s.store.ListUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "unknown error: "+err.Error(), http.StatusInternalServerError)
	}

	if getUserResult.Email == "" {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "email doesn't exist", http.StatusBadRequest)
	}

	// Populate current values
	res.Username = getUserResult.Username
	res.Email = getUserResult.Email

	if req.NewUsername != "" {
		updateUserParams := db.UpdateUserUsernameParams{}
		updateUserParams.Email = req.Email
		updateUserParams.Username = req.NewUsername
		err = s.store.ModifyUserUsername(ctx, updateUserParams)
		if err != nil {
			return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "unknown error: "+err.Error(), http.StatusInternalServerError)
		}
		res.Username = req.NewUsername
	}

	if req.NewPassword != "" {
		err := bcrypt.CompareHashAndPassword([]byte(getUserResult.Password), []byte(req.NewPassword))
		if err == nil {
			return nil, failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "new password cannot be same as old password", http.StatusBadRequest)
		}

		if len(req.NewPassword) < 8 {
			return nil, failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "password should be atleast 8 characters long", http.StatusBadRequest)
		}

		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			return nil, failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "incorrect password", http.StatusBadRequest)
		}
		updateUserParams := db.UpdateUserPasswordParams{}
		updateUserParams.Email = req.Email
		updateUserParams.Password = string(hashedPassword)
		err = s.store.ModifyUserPassword(ctx, updateUserParams)
		if err != nil {
			return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "unknown error: "+err.Error(), http.StatusInternalServerError)
		}
	}

	res.Email = req.Email
	return res, nil
}

// AuthenticateWithGoogle authenticates a user with Google.
func (s *Server) authenticateWithGoogle(_ context.Context) (res *authenticateWithGoogleResponse, err error) {
	// Add user authentication with Google logic here
	url := s.config.GcpOAuthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	res = &authenticateWithGoogleResponse{}
	res.URL = url
	return res, nil
}

// AuthenticateWithGoogleCallback authenticates a user with Google.
func (s *Server) AuthenticateWithGoogleCallback(ctx context.Context, req *AwesomeExpenseTrackerApi.AuthenticateWithGoogleCallbackRequest) (res *AwesomeExpenseTrackerApi.OAuth2Token, err error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to initialize validator: "+err.Error(), http.StatusInternalServerError)
	}

	if err = v.Validate(req); err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "failed to validate request: "+err.Error(), http.StatusBadRequest)
	}

	// Add user authentication with Google callback logic here
	token, err := s.config.GcpOAuthConfig.Exchange(ctx, req.Code)
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to exchange token: "+err.Error(), http.StatusInternalServerError)
	}

	// Parse the token and return the response
	res = &AwesomeExpenseTrackerApi.OAuth2Token{
		AuthProvider: utils.GoogleAuthProvider,
	}

	res, err = s.oauthTokenParser(ctx, res, token)
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "cannot parse token: "+err.Error(), http.StatusInternalServerError)
	}

	// Check if the email already exists in the database
	getUserResult, err := s.store.ListUserByEmail(ctx, res.Email)
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "unknown error: "+err.Error(), http.StatusInternalServerError)
	}

	// If the user doesn't exist, create a new user
	if getUserResult.Email == "" {
		createUserParams := db.CreateUserParams{}
		createUserParams.Email = res.Email
		createUserParams.Username = res.Email
		_, err = s.store.RegisterUser(ctx, createUserParams)
		if err != nil {
			return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "unknown error: "+err.Error(), http.StatusInternalServerError)
		}
		log.Default().Println("User: " + res.Email + " created successfully")
	}
	return res, nil
}

// Write all utility functions here

func (s *Server) validateRegisterRequest(req *AwesomeExpenseTrackerApi.RegisterUserRequest) error {
	if req.Username == "" || req.Password == "" {
		return failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "username or password cannot be empty", http.StatusBadRequest)
	}

	if req.Email == "" {
		return failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "email cannot be empty", http.StatusBadRequest)
	}

	if req.ConfirmPassword == "" || req.Password != req.ConfirmPassword {
		return failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "passwords do not match", http.StatusBadRequest)
	}

	if req.Name == "" {
		return failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "name cannot be empty", http.StatusBadRequest)
	}

	if !(len(req.Name) < 8 || len(req.Name) > 20) {
		return failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "name should be between 8 and 20 characters", http.StatusBadRequest)
	}

	if len(req.Password) < 8 {
		return failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "password should be atleast 8 characters long", http.StatusBadRequest)
	}

	return nil
}

func (s *Server) generateJWTToken(email string) (*oauth2.Token, error) {
	// Create a new token object for the access token
	idToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"iss":   "awesome-expense-tracker",
		"exp":   time.Now().Add(time.Hour * 1).Unix(), // Access token expires after 1 hour
	})

	// Sign and get the complete encoded token as a string using the secret
	idTokenString, err := idToken.SignedString(s.config.JwtKey)
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to sign access token: "+err.Error(), http.StatusInternalServerError)
	}

	token := &oauth2.Token{
		AccessToken:  "",
		RefreshToken: "",
		Expiry:       time.Now().Add(time.Hour * 1),
		TokenType:    "Bearer",
	}
	extra := make(map[string]interface{})
	extra["id_token"] = idTokenString
	token = token.WithExtra(extra)

	return token, nil
}

// oauthTokenParser parses the oauth2.Token (t) and returns the AwesomeExpenseTrackerApi.OAuth2Token (v)
func (s *Server) oauthTokenParser(_ context.Context, v *AwesomeExpenseTrackerApi.OAuth2Token, t *oauth2.Token) (*AwesomeExpenseTrackerApi.OAuth2Token, error) {
	v.ExpiresAt = t.Expiry.String()
	v.TokenType = t.TokenType

	idToken, ok := t.Extra("id_token").(string)
	if !ok {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to parse id_token", http.StatusInternalServerError)
	}
	v.IdToken = idToken

	userClaims, _ := jwt.Parse(idToken, nil)

	email, ok := userClaims.Claims.(jwt.MapClaims)["email"].(string)
	if !ok {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to parse email", http.StatusInternalServerError)
	}
	v.Email = email

	name, ok := userClaims.Claims.(jwt.MapClaims)["name"].(string)
	if !ok {
		v.Name = ""
	}
	v.Name = name

	profilePic, ok := userClaims.Claims.(jwt.MapClaims)["picture"].(string)
	if !ok {
		v.ProfilePic = ""
	}
	v.ProfilePic = profilePic

	return v, nil
}
