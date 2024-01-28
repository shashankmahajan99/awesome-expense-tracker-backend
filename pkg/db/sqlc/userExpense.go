package db

import (
	"context"
)

// AddExpense creates a new expense.
func (store *MySQLStore) AddExpense(ctx context.Context, arg CreateExpenseParams) (Expense, error) {
	var result Expense
	var expenseID int32
	err := store.execTx(ctx, func(q *Queries) error {
		res, err := q.CreateExpense(ctx, arg)
		if err != nil {
			return err
		}
		lastInsertID, err := res.LastInsertId()
		if err != nil {
			return err
		}
		expenseID = int32(lastInsertID)
		return nil
	})
	if err != nil {
		return result, err
	}
	result, err = store.GetExpenseByIdPvt(ctx, expenseID)
	if err != nil {
		return result, err
	}

	return result, nil
}

// ListExpenseByID lists all expenses for a user.
func (store *MySQLStore) ListExpenseByID(ctx context.Context, arg GetExpenseByIdParams) (Expense, error) {
	var result Expense
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		res, err := q.GetExpenseById(ctx, arg)
		if err != nil {
			return err
		}
		result = res.Expense
		return err
	})
	if err != nil {
		return result, err
	}
	return result, nil
}

// ListExpenseByEmail lists all expenses for a user.
func (store *MySQLStore) ListExpenseByEmail(ctx context.Context, arg string) ([]Expense, error) {
	var result []Expense
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		res, err := q.GetExpensesByUserId(ctx, arg)
		if err != nil {
			return err
		}
		result = make([]Expense, len(res))
		for i, expense := range res {
			result[i] = expense.Expense
		}
		return nil
	})
	if err != nil {
		return result, err
	}
	return result, nil
}

// RemoveExpense removes an expense.
func (store *MySQLStore) RemoveExpense(ctx context.Context, arg DeleteExpenseParams) (int64, error) {
	numberOfRows := int64(0)
	err := store.execTx(ctx, func(q *Queries) error {
		numRows, err := q.DeleteExpense(ctx, arg)
		numberOfRows = numRows
		return err
	})

	return numberOfRows, err
}

// ModifyExpense updates an expense.
func (store *MySQLStore) ModifyExpense(ctx context.Context, arg UpdateExpenseParams) (Expense, error) {
	var result Expense
	err := store.execTx(ctx, func(q *Queries) error {
		res, err := q.UpdateExpense(ctx, arg)
		if err != nil {
			return err
		}
		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return nil
		}
		return nil
	})
	if err != nil {
		return result, err
	}
	res, err := store.GetExpenseByIdPvt(ctx, arg.ID)
	if err != nil {
		return result, err
	}
	result = Expense{
		ID:          res.ID,
		UserID:      res.UserID,
		Amount:      res.Amount,
		Description: res.Description,
		Category:    res.Category,
		TxDate:      res.TxDate,
		Tag:         res.Tag,
		PaidTo:      res.PaidTo,
		PaidBy:      res.PaidBy,
		Flow:        res.Flow,
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
	}
	return result, nil
}
