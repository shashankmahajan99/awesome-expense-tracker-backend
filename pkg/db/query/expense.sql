-- name: CreateExpense :execresult
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
);

-- name: GetExpenseById :one
SELECT sqlc.embed(expenses) FROM expenses
JOIN users ON expenses.user_id = users.id
WHERE users.email = ? and expenses.id = ?;

-- name: GetExpenseByIdPvt :one
SELECT * FROM expenses
WHERE id = ?;

-- name: GetExpensesByUserId :many
SELECT sqlc.embed(expenses)
  FROM expenses
JOIN users ON expenses.user_id = users.id
WHERE users.email = ?;

-- name: UpdateExpense :execresult
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
) AND expenses.id = ?;

-- name: DeleteExpense :execrows
DELETE FROM expenses
WHERE user_id IN (
  SELECT id FROM users WHERE email = ?
) AND expenses.id = ?;


-- name: ListExpenses :many
SELECT * FROM expenses
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: CountExpenses :one
SELECT count(*) FROM expenses;
