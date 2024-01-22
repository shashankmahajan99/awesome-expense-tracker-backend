// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	CountExpenses(ctx context.Context) (int64, error)
	CountProfiles(ctx context.Context) (int64, error)
	CountReports(ctx context.Context) (int64, error)
	CountSettings(ctx context.Context) (int64, error)
	CountUsers(ctx context.Context) (int64, error)
	CreateExpense(ctx context.Context, arg CreateExpenseParams) (sql.Result, error)
	CreateProfile(ctx context.Context, arg CreateProfileParams) (sql.Result, error)
	CreateReport(ctx context.Context, arg CreateReportParams) (sql.Result, error)
	CreateSetting(ctx context.Context, arg CreateSettingParams) (sql.Result, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error)
	DeleteExpense(ctx context.Context, id int32) error
	DeleteProfile(ctx context.Context, email string) error
	DeleteReport(ctx context.Context, id int32) error
	DeleteSetting(ctx context.Context, id int32) error
	DeleteUser(ctx context.Context, username string) error
	GetExpense(ctx context.Context, id int32) (Expense, error)
	GetProfileByEmail(ctx context.Context, email string) (GetProfileByEmailRow, error)
	GetProfileByID(ctx context.Context, id int32) (Profile, error)
	GetReport(ctx context.Context, id int32) (Report, error)
	GetSetting(ctx context.Context, id int32) (Setting, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByID(ctx context.Context, id int32) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	ListExpenses(ctx context.Context, arg ListExpensesParams) ([]Expense, error)
	ListProfiles(ctx context.Context, arg ListProfilesParams) ([]Profile, error)
	ListReports(ctx context.Context, arg ListReportsParams) ([]Report, error)
	ListSettings(ctx context.Context, arg ListSettingsParams) ([]Setting, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	UpdateExpense(ctx context.Context, arg UpdateExpenseParams) error
	UpdateProfileBio(ctx context.Context, arg UpdateProfileBioParams) (sql.Result, error)
	UpdateProfileName(ctx context.Context, arg UpdateProfileNameParams) (sql.Result, error)
	UpdateProfileProfilePicture(ctx context.Context, arg UpdateProfileProfilePictureParams) (sql.Result, error)
	UpdateReport(ctx context.Context, arg UpdateReportParams) error
	UpdateSetting(ctx context.Context, arg UpdateSettingParams) error
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error
	UpdateUserUsername(ctx context.Context, arg UpdateUserUsernameParams) error
}

var _ Querier = (*Queries)(nil)
