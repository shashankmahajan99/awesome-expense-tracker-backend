package apipkg

import (
	"context"
	"net/http"

	AwesomeExpenseTrackerApi "github.com/shashankmahajan99/awesome-expense-tracker-backend/api"
	db "github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/db/sqlc"
	failuremanagement "github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/failure_management"
	"github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/utils"
)

// UserProfileServer is the interface that provides user profile methods.
type UserProfileServer interface {
	GetUserProfileAPI(ctx context.Context, req *AwesomeExpenseTrackerApi.GetUserProfileRequest) (*AwesomeExpenseTrackerApi.GetUserProfileResponse, error)
	CreateUserProfileAPI(ctx context.Context, req *AwesomeExpenseTrackerApi.CreateUserProfileRequest) (*AwesomeExpenseTrackerApi.CreateUserProfileResponse, error)
	UpdateUserProfileAPI(ctx context.Context, req *AwesomeExpenseTrackerApi.UpdateUserProfileRequest) (*AwesomeExpenseTrackerApi.UpdateUserProfileResponse, error)
}

// GetUserProfileAPI gets a user profile.
func (s *Server) GetUserProfileAPI(ctx context.Context, req *AwesomeExpenseTrackerApi.GetUserProfileRequest) (res *AwesomeExpenseTrackerApi.GetUserProfileResponse, err error) {
	res = &AwesomeExpenseTrackerApi.GetUserProfileResponse{}

	userProfileResult, err := s.store.ListProfileByEmail(ctx, req.Email)
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to get user profile: "+err.Error(), http.StatusInternalServerError)
	}
	res.Bio = userProfileResult.Bio
	res.ProfileName = userProfileResult.ProfileName
	res.ProfilePicture = userProfileResult.ProfilePicture

	return res, nil
}

// CreateUserProfileAPI creates a new user profile.
func (s *Server) CreateUserProfileAPI(ctx context.Context, req *AwesomeExpenseTrackerApi.CreateUserProfileRequest) (res *AwesomeExpenseTrackerApi.CreateUserProfileResponse, err error) {
	res = &AwesomeExpenseTrackerApi.CreateUserProfileResponse{}

	// Check if user profile already exists
	listUserProfileResult, err := s.store.ListProfileByEmail(ctx, req.Email)
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to get user profile: "+err.Error(), http.StatusInternalServerError)
	}
	if listUserProfileResult.UserID != 0 {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "user profile already exists", http.StatusBadRequest)
	}

	createProfileParams := db.CreateProfileParams{
		Email:          req.Email,
		ProfileName:    req.ProfileName,
		ProfilePicture: req.ProfilePicture,
		Bio:            req.Bio,
	}
	userProfileResult, err := s.store.AddProfile(ctx, createProfileParams)
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to create user profile: "+err.Error(), http.StatusInternalServerError)
	}
	res.Bio = userProfileResult.Bio
	res.ProfileName = userProfileResult.ProfileName
	res.ProfilePicture = userProfileResult.ProfilePicture

	return res, nil
}

// UpdateUserProfileAPI updates a user profile.
func (s *Server) UpdateUserProfileAPI(ctx context.Context, req *AwesomeExpenseTrackerApi.UpdateUserProfileRequest) (res *AwesomeExpenseTrackerApi.UpdateUserProfileResponse, err error) {
	res = &AwesomeExpenseTrackerApi.UpdateUserProfileResponse{}

	// Check if user profile doesn't exists
	listUserProfileResult, err := s.store.ListProfileByEmail(ctx, req.Email)
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to get user profile: "+err.Error(), http.StatusInternalServerError)
	}
	if listUserProfileResult.UserID == 0 {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "user profile doesn't exists", http.StatusBadRequest)
	}

	// Populate current values
	res.Bio = listUserProfileResult.Bio
	res.ProfileName = listUserProfileResult.ProfileName
	res.ProfilePicture = listUserProfileResult.ProfilePicture

	if req.NewBio != "" {
		updateProfileParams := db.UpdateProfileBioParams{
			Bio:   req.NewBio,
			Email: req.Email,
		}
		userProfileResult, err := s.store.ModifyProfileBio(ctx, updateProfileParams)
		if err != nil {
			return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to update user profile: "+err.Error(), http.StatusInternalServerError)
		}
		res.Bio = userProfileResult.Bio
	}

	if req.NewProfileName != "" {
		updateProfileParams := db.UpdateProfileNameParams{
			ProfileName: req.NewProfileName,
			Email:       req.Email,
		}
		userProfileResult, err := s.store.ModifyProfileName(ctx, updateProfileParams)
		if err != nil {
			return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to update user profile: "+err.Error(), http.StatusInternalServerError)
		}
		res.ProfileName = userProfileResult.ProfileName
	}

	if req.NewProfilePicture != "" {
		updateProfileParams := db.UpdateProfileProfilePictureParams{
			ProfilePicture: req.NewProfilePicture,
			Email:          req.Email,
		}
		userProfileResult, err := s.store.ModifyProfilePicture(ctx, updateProfileParams)
		if err != nil {
			return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to update user profile: "+err.Error(), http.StatusInternalServerError)
		}
		res.ProfilePicture = userProfileResult.ProfilePicture
	}
	return res, nil
}
