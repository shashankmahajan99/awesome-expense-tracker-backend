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
SELECT count(*) FROM expenses
`

func (q *Queries) CountExpenses(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countExpenses)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createExpense = `-- name: CreateExpense :execresult
INSERT INTO expenses (
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
  (SELECT id from users where email = ?), ?, ?, ?, ?, ?, ?, ?, ?
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

const deleteExpense = `-- name: DeleteExpense :execrows
DELETE FROM expenses
WHERE user_id IN (
  SELECT id FROM users WHERE email = ?
) AND expenses.id = ?
`

type DeleteExpenseParams struct {
	Email string `json:"email"`
	ID    int32  `json:"id"`
}

func (q *Queries) DeleteExpense(ctx context.Context, arg DeleteExpenseParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, deleteExpense, arg.Email, arg.ID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const getExpenseById = `-- name: GetExpenseById :one
SELECT expenses.id, expenses.user_id, expenses.amount, expenses.description, expenses.category, expenses.tx_date, expenses.tag, expenses.paid_to, expenses.paid_by, expenses.flow, expenses.created_at, expenses.updated_at FROM expenses
JOIN users ON expenses.user_id = users.id
WHERE users.email = ? and expenses.id = ?
`

type GetExpenseByIdParams struct {
	Email string `json:"email"`
	ID    int32  `json:"id"`
}

type GetExpenseByIdRow struct {
	Expense Expense `json:"expense"`
}

func (q *Queries) GetExpenseById(ctx context.Context, arg GetExpenseByIdParams) (GetExpenseByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getExpenseById, arg.Email, arg.ID)
	var i GetExpenseByIdRow
	err := row.Scan(
		&i.Expense.ID,
		&i.Expense.UserID,
		&i.Expense.Amount,
		&i.Expense.Description,
		&i.Expense.Category,
		&i.Expense.TxDate,
		&i.Expense.Tag,
		&i.Expense.PaidTo,
		&i.Expense.PaidBy,
		&i.Expense.Flow,
		&i.Expense.CreatedAt,
		&i.Expense.UpdatedAt,
	)
	return i, err
}

const getExpenseByIdPvt = `-- name: GetExpenseByIdPvt :one
SELECT id, user_id, amount, description, category, tx_date, tag, paid_to, paid_by, flow, created_at, updated_at FROM expenses
WHERE id = ?
`

func (q *Queries) GetExpenseByIdPvt(ctx context.Context, id int32) (Expense, error) {
	row := q.db.QueryRowContext(ctx, getExpenseByIdPvt, id)
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
SELECT expenses.id, expenses.user_id, expenses.amount, expenses.description, expenses.category, expenses.tx_date, expenses.tag, expenses.paid_to, expenses.paid_by, expenses.flow, expenses.created_at, expenses.updated_at
  FROM expenses
JOIN users ON expenses.user_id = users.id
WHERE users.email = ?
`

type GetExpensesByUserIdRow struct {
	Expense Expense `json:"expense"`
}

func (q *Queries) GetExpensesByUserId(ctx context.Context, email string) ([]GetExpensesByUserIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getExpensesByUserId, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetExpensesByUserIdRow{}
	for rows.Next() {
		var i GetExpensesByUserIdRow
		if err := rows.Scan(
			&i.Expense.ID,
			&i.Expense.UserID,
			&i.Expense.Amount,
			&i.Expense.Description,
			&i.Expense.Category,
			&i.Expense.TxDate,
			&i.Expense.Tag,
			&i.Expense.PaidTo,
			&i.Expense.PaidBy,
			&i.Expense.Flow,
			&i.Expense.CreatedAt,
			&i.Expense.UpdatedAt,
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

const listExpenses = `-- name: ListExpenses :many
SELECT id, user_id, amount, description, category, tx_date, tag, paid_to, paid_by, flow, created_at, updated_at FROM expenses
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

const updateExpense = `-- name: UpdateExpense :execresult
UPDATE expenses
SET
  amount = ?,
  description = ?,
  category = ?,
  tx_date = ?,
  tag = ?,
  paid_to = ?,
  paid_by = ?,
  flow = ?
WHERE user_id IN (
  SELECT id FROM users WHERE email = ?
) AND expenses.id = ?
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
	Email       string    `json:"email"`
	ID          int32     `json:"id"`
}

func (q *Queries) UpdateExpense(ctx context.Context, arg UpdateExpenseParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateExpense,
		arg.Amount,
		arg.Description,
		arg.Category,
		arg.TxDate,
		arg.Tag,
		arg.PaidTo,
		arg.PaidBy,
		arg.Flow,
		arg.Email,
		arg.ID,
	)
}
