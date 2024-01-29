package apipkg

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	AwesomeExpenseTrackerApi "github.com/shashankmahajan99/awesome-expense-tracker-backend/api"
	db "github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/db/sqlc"
	failuremanagement "github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/failure_management"
	"github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/utils"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ExpenseMgmtServer is the interface that provides expense management methods.
type ExpenseMgmtServer interface {
	CreateExpense(ctx context.Context, req *AwesomeExpenseTrackerApi.CreateExpenseRequest) (*AwesomeExpenseTrackerApi.CreateExpenseResponse, error)
	GetExpense(ctx context.Context, req *AwesomeExpenseTrackerApi.GetExpenseRequest) (*AwesomeExpenseTrackerApi.GetExpenseResponse, error)
	ListExpenses(ctx context.Context, req *AwesomeExpenseTrackerApi.ListExpensesRequest) (*AwesomeExpenseTrackerApi.ListExpensesResponse, error)
	DeleteExpense(ctx context.Context, req *AwesomeExpenseTrackerApi.DeleteExpenseRequest) (*AwesomeExpenseTrackerApi.DeleteExpenseResponse, error)
	UpdateExpense(ctx context.Context, req *AwesomeExpenseTrackerApi.UpdateExpenseRequest) (*AwesomeExpenseTrackerApi.UpdateExpenseResponse, error)
}

// CreateExpense creates a new expense.
func (s *Server) CreateExpense(ctx context.Context, req *AwesomeExpenseTrackerApi.CreateExpenseRequest) (res *AwesomeExpenseTrackerApi.CreateExpenseResponse, err error) {
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
	expenseResult, err := s.store.AddExpense(ctx, createExpenseParams)
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to create expense: "+err.Error(), http.StatusInternalServerError)
	}

	res.Id = expenseResult.ID

	return res, nil
}

// GetExpense gets an expense.
func (s *Server) GetExpense(ctx context.Context, req *AwesomeExpenseTrackerApi.GetExpenseRequest) (res *AwesomeExpenseTrackerApi.GetExpenseResponse, err error) {
	res = &AwesomeExpenseTrackerApi.GetExpenseResponse{}
	userEmail, ok := ctx.Value(utils.EmailKey).(string)
	if !ok {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to get user email: "+err.Error(), http.StatusInternalServerError)
	}

	// get expense
	getExpenseByIDParams := db.GetExpenseByIdParams{
		Email: userEmail,
		ID:    int32(req.Id),
	}
	expenseResult, err := s.store.ListExpenseByID(ctx, getExpenseByIDParams)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "expense doesn't exist", http.StatusBadRequest)
		}
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to get expense: "+err.Error(), http.StatusInternalServerError)
	}

	res.Expense = &AwesomeExpenseTrackerApi.ExpenseObject{
		Id:          int64(expenseResult.ID),
		Amount:      expenseResult.Amount,
		Description: expenseResult.Description,
		Date:        timestamppb.New(expenseResult.TxDate),
		Tag:         expenseResult.Tag,
		Category:    expenseResult.Category,
		PaidTo:      expenseResult.PaidTo,
		PaidBy:      expenseResult.PaidBy,
		Flow:        expenseResult.Flow,
	}

	return res, nil
}

// ListExpenses gets all expenses.
func (s *Server) ListExpenses(ctx context.Context, _ *AwesomeExpenseTrackerApi.ListExpensesRequest) (res *AwesomeExpenseTrackerApi.ListExpensesResponse, err error) {
	res = &AwesomeExpenseTrackerApi.ListExpensesResponse{}
	userEmail, ok := ctx.Value(utils.EmailKey).(string)
	if !ok {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to get user email: "+err.Error(), http.StatusInternalServerError)
	}

	// get expense
	expenseResult, err := s.store.ListExpenseByEmail(ctx, userEmail)
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to get expense: "+err.Error(), http.StatusInternalServerError)
	}

	expenseList := make([]*AwesomeExpenseTrackerApi.ExpenseObject, 0, len(expenseResult))
	for _, expense := range expenseResult {
		expenseList = append(expenseList, &AwesomeExpenseTrackerApi.ExpenseObject{
			Id:          int64(expense.ID),
			Amount:      expense.Amount,
			Description: expense.Description,
			Date:        timestamppb.New(expense.TxDate),
			Tag:         expense.Tag,
			Category:    expense.Category,
			PaidTo:      expense.PaidTo,
			PaidBy:      expense.PaidBy,
			Flow:        expense.Flow,
		})
	}

	res.Expenses = expenseList

	return res, nil
}

// DeleteExpense deletes an expense.
func (s *Server) DeleteExpense(ctx context.Context, req *AwesomeExpenseTrackerApi.DeleteExpenseRequest) (res *AwesomeExpenseTrackerApi.DeleteExpenseResponse, err error) {
	res = &AwesomeExpenseTrackerApi.DeleteExpenseResponse{}
	userEmail, ok := ctx.Value(utils.EmailKey).(string)
	if !ok {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to get user email: "+err.Error(), http.StatusInternalServerError)
	}

	// delete expense
	numOfRowsAffected, err := s.store.RemoveExpense(ctx, db.DeleteExpenseParams{
		Email: userEmail,
		ID:    int32(req.Id),
	})
	if err != nil {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to delete expense: "+err.Error(), http.StatusInternalServerError)
	}

	if numOfRowsAffected == 0 {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, fmt.Sprintf("expense with id %d doesn't exist", req.Id), http.StatusBadRequest)
	}
	res.Id = int32(req.Id)
	return res, nil
}

// UpdateExpense updates an expense.
func (s *Server) UpdateExpense(ctx context.Context, req *AwesomeExpenseTrackerApi.UpdateExpenseRequest) (res *AwesomeExpenseTrackerApi.UpdateExpenseResponse, err error) {
	res = &AwesomeExpenseTrackerApi.UpdateExpenseResponse{}
	userEmail, ok := ctx.Value(utils.EmailKey).(string)
	if !ok {
		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to get user email: "+err.Error(), http.StatusInternalServerError)
	}

	// update expense
	updateExpenseParams := db.UpdateExpenseParams{
		Amount:      req.Expense.Amount,
		Description: req.Expense.Description,
		TxDate:      req.Expense.Date.AsTime(),
		Tag:         req.Expense.Tag,
		Email:       userEmail,
		Category:    req.Expense.Category,
		PaidTo:      req.Expense.PaidTo,
		PaidBy:      req.Expense.PaidBy,
		Flow:        req.Expense.Flow,
		ID:          int32(req.Id),
	}
	expenseResult, err := s.store.ModifyExpense(ctx, updateExpenseParams)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "expense doesn't exist", http.StatusBadRequest)
		}

		return nil, failuremanagement.NewCustomErrorResponse(utils.INTERNALERROR, "failed to update expense: "+err.Error(), http.StatusInternalServerError)
	}

	res.Id = expenseResult.ID

	return res, nil
}
