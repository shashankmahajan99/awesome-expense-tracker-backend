
-- name: CreateExpense :execresult
INSERT INTO Expenses (
  user_id,
  amount,
  description,
  category
) VALUES (
  ?, ?, ?, ?
);

-- name: GetExpense :one
SELECT * FROM Expenses
WHERE id = ?;

-- name: UpdateExpense :exec
UPDATE Expenses
SET
  user_id = ?,
  amount = ?,
  description = ?,
  category = ?
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
