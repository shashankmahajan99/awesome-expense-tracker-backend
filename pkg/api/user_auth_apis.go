package apipkg

import (
	"context"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	AwesomeExpenseTrackerApi "github.com/shashankmahajan99/awesome-expense-tracker-backend/api"
	db "github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/db/sqlc"
	failuremanagement "github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/failure_management"
	"github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type authenticateWithGoogleResponse struct {
	URL string `json:"url"`
}

// UserAuthServer is the server API for UserAuthentication service.
type UserAuthServer interface {
	LoginUserAPI(ctx context.Context, req *AwesomeExpenseTrackerApi.LoginUserRequest) (*AwesomeExpenseTrackerApi.OAuth2Token, error)
	RegisterUserAPI(ctx context.Context, req *AwesomeExpenseTrackerApi.RegisterUserRequest) (*AwesomeExpenseTrackerApi.OAuth2Token, error)
	DeleteUserAPI(ctx context.Context, req *AwesomeExpenseTrackerApi.DeleteUserRequest) (*AwesomeExpenseTrackerApi.DeleteUserResponse, error)
	UpdateUserAPI(ctx context.Context, req *AwesomeExpenseTrackerApi.UpdateUserRequest) (*AwesomeExpenseTrackerApi.UpdateUserResponse, error)
	AuthenticateWithGoogleCallback(ctx context.Context, req *AwesomeExpenseTrackerApi.AuthenticateWithGoogleCallbackRequest) (*AwesomeExpenseTrackerApi.OAuth2Token, error)
}

// LoginUserAPI logs in a user.
func (s *Server) LoginUserAPI(ctx context.Context, req *AwesomeExpenseTrackerApi.LoginUserRequest) (res *AwesomeExpenseTrackerApi.OAuth2Token, err error) {

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
	token, err := s.authenticationManager.GenerateJWTToken(getUserResult.Email, "admin")
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "cannot generate access token: "+err.Error(), http.StatusInternalServerError)
	}

	// Parse the token and return the response
	res = &AwesomeExpenseTrackerApi.OAuth2Token{}
	res, idToken, err := s.oauthTokenParser(ctx, res, token)
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "cannot parse token: "+err.Error(), http.StatusInternalServerError)
	}
	res.Email = getUserResult.Email
	res.AuthProvider = req.AuthProvider
	s.SetTokenCookie(ctx, idToken)
	return res, nil
}

// RegisterUserAPI registers a user.
func (s *Server) RegisterUserAPI(ctx context.Context, req *AwesomeExpenseTrackerApi.RegisterUserRequest) (res *AwesomeExpenseTrackerApi.OAuth2Token, err error) {
	res = &AwesomeExpenseTrackerApi.OAuth2Token{}

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
	token, err := s.authenticationManager.GenerateJWTToken(req.Email, "admin")
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "cannot generate access token: "+err.Error(), http.StatusInternalServerError)
	}

	// Parse the token and return the response
	res, idToken, err := s.oauthTokenParser(ctx, res, token)
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "cannot parse token: "+err.Error(), http.StatusInternalServerError)
	}
	res.Email = req.Email
	res.AuthProvider = req.AuthProvider
	s.SetTokenCookie(ctx, idToken)
	return res, nil
}

// DeleteUserAPI deletes a user.
func (s *Server) DeleteUserAPI(ctx context.Context, req *AwesomeExpenseTrackerApi.DeleteUserRequest) (res *AwesomeExpenseTrackerApi.DeleteUserResponse, err error) {
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

// UpdateUserAPI updates a user.
func (s *Server) UpdateUserAPI(ctx context.Context, req *AwesomeExpenseTrackerApi.UpdateUserRequest) (res *AwesomeExpenseTrackerApi.UpdateUserResponse, err error) {
	res = &AwesomeExpenseTrackerApi.UpdateUserResponse{}

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
	url := s.authenticationManager.GcpOAuthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	res = &authenticateWithGoogleResponse{}
	res.URL = url
	return res, nil
}

// AuthenticateWithGoogleCallback authenticates a user with Google.
func (s *Server) AuthenticateWithGoogleCallback(ctx context.Context, req *AwesomeExpenseTrackerApi.AuthenticateWithGoogleCallbackRequest) (res *AwesomeExpenseTrackerApi.OAuth2Token, err error) {

	// Add user authentication with Google callback logic here
	token, err := s.authenticationManager.GcpOAuthConfig.Exchange(ctx, req.Code)
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to exchange token: "+err.Error(), http.StatusInternalServerError)
	}

	// Parse the token and return the response
	res = &AwesomeExpenseTrackerApi.OAuth2Token{
		AuthProvider: utils.GoogleAuthProvider,
	}

	res, idToken, err := s.oauthTokenParser(ctx, res, token)
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
	s.SetTokenCookie(ctx, idToken)
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

	if len(req.Password) < 8 {
		return failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "password should be atleast 8 characters long", http.StatusBadRequest)
	}

	return nil
}

// oauthTokenParser parses the oauth2.Token (t) and returns the AwesomeExpenseTrackerApi.OAuth2Token (v)
func (s *Server) oauthTokenParser(_ context.Context, v *AwesomeExpenseTrackerApi.OAuth2Token, t *oauth2.Token) (*AwesomeExpenseTrackerApi.OAuth2Token, string, error) {
	v.ExpiresAt = t.Expiry.String()
	v.TokenType = t.TokenType

	idToken, ok := t.Extra("id_token").(string)
	if !ok {
		return nil, "", failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to parse id_token", http.StatusInternalServerError)
	}

	userClaims, _ := jwt.Parse(idToken, nil)

	email, ok := userClaims.Claims.(jwt.MapClaims)["email"].(string)
	if !ok {
		return nil, "", failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to parse email", http.StatusInternalServerError)
	}
	v.Email = email

	return v, idToken, nil
}

// SetTokenCookie sets the token as a cookie.
func (s *Server) SetTokenCookie(ctx context.Context, idToken string) error {

	metadata.Pairs("id_token", idToken)
	cookie := http.Cookie{
		Name:     "id_token",
		Value:    idToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   3600, // 1 hour
	}

	grpc.SendHeader(ctx, metadata.Pairs("Set-Cookie", cookie.String()))
	return nil
}
