package apipkg

import (
	"context"
	"net/http"

	"github.com/bufbuild/protovalidate-go"
	AwesomeExpenseTrackerApi "github.com/shashankmahajan99/awesome-expense-tracker-backend/api"
	db "github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/db/sqlc"
	failuremanagement "github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/failure_management"
	"github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/utils"
)

// GetUserProfile gets a user profile.
func (s *Server) GetUserProfile(_ context.Context, _ *AwesomeExpenseTrackerApi.GetUserProfileRequest) (*AwesomeExpenseTrackerApi.GetUserProfileResponse, error) {
	return nil, nil
}

// UpdateUserProfile updates a user profile.
func (s *Server) UpdateUserProfile(ctx context.Context, req *AwesomeExpenseTrackerApi.UpdateUserProfileRequest) (res *AwesomeExpenseTrackerApi.UpdateUserProfileResponse, err error) {
	res = &AwesomeExpenseTrackerApi.UpdateUserProfileResponse{}
	v, err := protovalidate.New()
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNAL_ERROR, "failed to initialize validator: "+err.Error(), http.StatusInternalServerError)
	}

	err = v.Validate(req)
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INVALID_REQUEST, "failed to validate request: "+err.Error(), http.StatusBadRequest)
	}

	if req.NewBio != "" {
		updateProfileParams := db.UpdateProfileBioParams{
			Bio: req.NewBio,
		}
		userProfileResult, err := s.store.ModifyProfileBio(ctx, updateProfileParams)
		if err != nil {
			return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNAL_ERROR, "failed to update user profile: "+err.Error(), http.StatusInternalServerError)
		}
		res.Bio = userProfileResult.Bio
	}

	if req.NewName != "" {
		updateProfileParams := db.UpdateProfileNameParams{
			Name: req.NewName,
		}
		userProfileResult, err := s.store.ModifyProfileName(ctx, updateProfileParams)
		if err != nil {
			return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNAL_ERROR, "failed to update user profile: "+err.Error(), http.StatusInternalServerError)
		}
		res.Name = userProfileResult.Name
	}

	if req.NewProfilePicture != "" {
		updateProfileParams := db.UpdateProfileProfilePictureParams{
			ProfilePicture: req.NewProfilePicture,
		}
		userProfileResult, err := s.store.ModifyProfilePicture(ctx, updateProfileParams)
		if err != nil {
			return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNAL_ERROR, "failed to update user profile: "+err.Error(), http.StatusInternalServerError)
		}
		res.Bio = userProfileResult.Bio
	}
	return nil, nil
}
