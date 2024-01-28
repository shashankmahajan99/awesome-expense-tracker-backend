// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: expense.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const countExpenses = `-- name: CountExpenses :one
SELECT count(*) FROM Expenses
`

func (q *Queries) CountExpenses(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countExpenses)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createExpense = `-- name: CreateExpense :execresult
INSERT INTO Expenses (
  user_id,
  amount,
  description,
  category,
  tx_date,
  tag,
  paid_to,
  paid_by,
  flow
) VALUES (
  (SELECT id from Users where email = ?), ?, ?, ?, ?, ?, ?, ?, ?
)
`

type CreateExpenseParams struct {
	Email       string    `json:"email"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	TxDate      time.Time `json:"tx_date"`
	Tag         string    `json:"tag"`
	PaidTo      string    `json:"paid_to"`
	PaidBy      string    `json:"paid_by"`
	Flow        string    `json:"flow"`
}

func (q *Queries) CreateExpense(ctx context.Context, arg CreateExpenseParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createExpense,
		arg.Email,
		arg.Amount,
		arg.Description,
		arg.Category,
		arg.TxDate,
		arg.Tag,
		arg.PaidTo,
		arg.PaidBy,
		arg.Flow,
	)
}

const deleteExpense = `-- name: DeleteExpense :exec
DELETE FROM Expenses
WHERE id = ?
`

func (q *Queries) DeleteExpense(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteExpense, id)
	return err
}

const getExpenseById = `-- name: GetExpenseById :one
SELECT id, user_id, amount, description, category, tx_date, tag, paid_to, paid_by, flow, created_at, updated_at FROM Expenses
WHERE id = ?
`

func (q *Queries) GetExpenseById(ctx context.Context, id int32) (Expense, error) {
	row := q.db.QueryRowContext(ctx, getExpenseById, id)
	var i Expense
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Amount,
		&i.Description,
		&i.Category,
		&i.TxDate,
		&i.Tag,
		&i.PaidTo,
		&i.PaidBy,
		&i.Flow,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getExpensesByUserId = `-- name: GetExpensesByUserId :many
SELECT (
  Expenses.id,
  user_id,
  amount,
  description,
  category,
  tx_date,
  tag,
  paid_to,
  paid_by,
  flow,
  Expesnes.created_at,
  Expenses.updated_at
  )
  FROM Expenses
JOIN Users ON Expenses.user_id = Users.id
WHERE Users.email = ?
`

func (q *Queries) GetExpensesByUserId(ctx context.Context, email string) ([]interface{}, error) {
	rows, err := q.db.QueryContext(ctx, getExpensesByUserId, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []interface{}{}
	for rows.Next() {
		var column_1 interface{}
		if err := rows.Scan(&column_1); err != nil {
			return nil, err
		}
		items = append(items, column_1)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listExpenses = `-- name: ListExpenses :many
SELECT id, user_id, amount, description, category, tx_date, tag, paid_to, paid_by, flow, created_at, updated_at FROM Expenses
ORDER BY id
LIMIT ?
OFFSET ?
`

type ListExpensesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListExpenses(ctx context.Context, arg ListExpensesParams) ([]Expense, error) {
	rows, err := q.db.QueryContext(ctx, listExpenses, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Expense{}
	for rows.Next() {
		var i Expense
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Amount,
			&i.Description,
			&i.Category,
			&i.TxDate,
			&i.Tag,
			&i.PaidTo,
			&i.PaidBy,
			&i.Flow,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateExpense = `-- name: UpdateExpense :exec
UPDATE Expenses
SET
  amount = ?,
  description = ?,
  category = ?,
  tx_date = ?,
  tag = ?,
  paid_to = ?,
  paid_by = ?,
  flow = ?
WHERE id = ?
`

type UpdateExpenseParams struct {
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	TxDate      time.Time `json:"tx_date"`
	Tag         string    `json:"tag"`
	PaidTo      string    `json:"paid_to"`
	PaidBy      string    `json:"paid_by"`
	Flow        string    `json:"flow"`
	ID          int32     `json:"id"`
}

func (q *Queries) UpdateExpense(ctx context.Context, arg UpdateExpenseParams) error {
	_, err := q.db.ExecContext(ctx, updateExpense,
		arg.Amount,
		arg.Description,
		arg.Category,
		arg.TxDate,
		arg.Tag,
		arg.PaidTo,
		arg.PaidBy,
		arg.Flow,
		arg.ID,
	)
	return err
}
