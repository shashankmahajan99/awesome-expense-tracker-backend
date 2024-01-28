package apipkg

import (
	"context"
	"net/http"
	"strconv"

	AwesomeExpenseTrackerApi "github.com/shashankmahajan99/awesome-expense-tracker-backend/api"
	db "github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/db/sqlc"
	failuremanagement "github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/failure_management"
	"github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/utils"
)

// ExpenseMgmtServer is the interface that provides expense management methods.
type ExpenseMgmtServer interface {
	CreateExpenseAPI(ctx context.Context, req *AwesomeExpenseTrackerApi.CreateExpenseRequest) (*AwesomeExpenseTrackerApi.CreateExpenseResponse, error)
	GetExpenseAPI(ctx context.Context, req *AwesomeExpenseTrackerApi.GetExpenseRequest) (*AwesomeExpenseTrackerApi.GetExpenseResponse, error)
	ListExpensesAPI(ctx context.Context, req *AwesomeExpenseTrackerApi.ListExpensesRequest) (*AwesomeExpenseTrackerApi.ListExpensesResponse, error)
	DeleteExpenseAPI(ctx context.Context, req *AwesomeExpenseTrackerApi.DeleteExpenseRequest) (*AwesomeExpenseTrackerApi.DeleteExpenseResponse, error)
	UpdateExpenseAPI(ctx context.Context, req *AwesomeExpenseTrackerApi.UpdateExpenseRequest) (*AwesomeExpenseTrackerApi.UpdateExpenseResponse, error)
}

// CreateExpenseAPI creates a new expense.
func (s *Server) CreateExpenseAPI(ctx context.Context, req *AwesomeExpenseTrackerApi.CreateExpenseRequest) (res *AwesomeExpenseTrackerApi.CreateExpenseResponse, err error) {
	res = &AwesomeExpenseTrackerApi.CreateExpenseResponse{}
	userEmail, ok := ctx.Value(utils.EmailKey).(string)
	if !ok {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to get user email: "+err.Error(), http.StatusInternalServerError)
	}
	// create expense
	createExpenseParams := db.CreateExpenseParams{
		Amount:      req.Expense.Amount,
		Description: req.Expense.Description,
		TxDate:      req.Expense.Date.AsTime(),
		Tag:         req.Expense.Tag,
		Email:       userEmail,
		Category:    req.Expense.Category,
		PaidTo:      req.Expense.PaidTo,
		PaidBy:      req.Expense.PaidBy,
		Flow:        req.Expense.Flow,
	}
	expenseResult, err := s.store.CreateExpense(ctx, createExpenseParams)
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to create expense: "+err.Error(), http.StatusInternalServerError)
	}

	expenseID, err := expenseResult.LastInsertId()
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to get expense id: "+err.Error(), http.StatusInternalServerError)
	}

	res.Id = strconv.FormatInt(expenseID, 10)

	return res, nil
}

func (s *Server) GetExpenseAPI(ctx context.Context, req *AwesomeExpenseTrackerApi.GetExpenseRequest) (res *AwesomeExpenseTrackerApi.GetExpenseResponse, err error) {

	return nil, nil
}

func (s *Server) ListExpensesAPI(ctx context.Context, req *AwesomeExpenseTrackerApi.ListExpensesRequest) (res *AwesomeExpenseTrackerApi.ListExpensesResponse, err error) {

	return nil, nil
}

func (s *Server) DeleteExpenseAPI(ctx context.Context, req *AwesomeExpenseTrackerApi.DeleteExpenseRequest) (res *AwesomeExpenseTrackerApi.DeleteExpenseResponse, err error) {

	return nil, nil
}

func (s *Server) UpdateExpenseAPI(ctx context.Context, req *AwesomeExpenseTrackerApi.UpdateExpenseRequest) (res *AwesomeExpenseTrackerApi.UpdateExpenseResponse, err error) {

	return nil, nil
}
