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
SELECT sqlc.embed(Expenses) FROM Expenses
JOIN Users ON Expenses.user_id = Users.id
WHERE Users.email = ? and Expenses.id = ?;

-- name: GetExpenseByIdPvt :one
SELECT * FROM Expenses
WHERE id = ?;

-- name: GetExpensesByUserId :many
SELECT sqlc.embed(Expenses)
  FROM Expenses
JOIN Users ON Expenses.user_id = Users.id
WHERE Users.email = ?;

-- name: UpdateExpense :execresult
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
WHERE user_id IN (
  SELECT id FROM Users WHERE email = ?
) AND Expenses.id = ?;

-- name: DeleteExpense :execrows
DELETE FROM Expenses
WHERE user_id IN (
  SELECT id FROM Users WHERE email = ?
) AND Expenses.id = ?;


-- name: ListExpenses :many
SELECT * FROM Expenses
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: CountExpenses :one
SELECT count(*) FROM Expenses;
