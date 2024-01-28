-- name: CreateExpense :execresult
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
);

-- name: GetExpenseById :one
SELECT * FROM Expenses
WHERE id = ?;

-- name: GetExpensesByUserId :many
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
WHERE Users.email = ?;

-- name: UpdateExpense :exec
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
WHERE id = ?;

-- name: DeleteExpense :exec
DELETE FROM Expenses
WHERE id = ?;

-- name: ListExpenses :many
SELECT * FROM Expenses
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: CountExpenses :one
SELECT count(*) FROM Expenses;
