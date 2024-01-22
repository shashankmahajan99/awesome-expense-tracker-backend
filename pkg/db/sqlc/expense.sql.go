// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: expense.sql

package db

import (
	"context"
	"database/sql"
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
  category
) VALUES (
  ?, ?, ?, ?
)
`

type CreateExpenseParams struct {
	UserID      sql.NullInt32 `json:"user_id"`
	Amount      string        `json:"amount"`
	Description string        `json:"description"`
	Category    string        `json:"category"`
}

func (q *Queries) CreateExpense(ctx context.Context, arg CreateExpenseParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createExpense,
		arg.UserID,
		arg.Amount,
		arg.Description,
		arg.Category,
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

const getExpense = `-- name: GetExpense :one
SELECT id, user_id, amount, description, category, created_at, updated_at FROM Expenses
WHERE id = ?
`

func (q *Queries) GetExpense(ctx context.Context, id int32) (Expense, error) {
	row := q.db.QueryRowContext(ctx, getExpense, id)
	var i Expense
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Amount,
		&i.Description,
		&i.Category,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listExpenses = `-- name: ListExpenses :many
SELECT id, user_id, amount, description, category, created_at, updated_at FROM Expenses
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
  user_id = ?,
  amount = ?,
  description = ?,
  category = ?
WHERE id = ?
`

type UpdateExpenseParams struct {
	UserID      sql.NullInt32 `json:"user_id"`
	Amount      string        `json:"amount"`
	Description string        `json:"description"`
	Category    string        `json:"category"`
	ID          int32         `json:"id"`
}

func (q *Queries) UpdateExpense(ctx context.Context, arg UpdateExpenseParams) error {
	_, err := q.db.ExecContext(ctx, updateExpense,
		arg.UserID,
		arg.Amount,
		arg.Description,
		arg.Category,
		arg.ID,
	)
	return err
}
