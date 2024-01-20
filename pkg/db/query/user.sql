-- name: CreateUser :execresult
INSERT INTO Users (
  username,
  password,
  email
) VALUES (
  ?, ?, ?
);

-- name: GetUserByUsername :one
SELECT * FROM Users
WHERE username = ?;

-- name: GetUserByID :one
SELECT * FROM Users
WHERE id = ?;

-- name: UpdateUser :exec
UPDATE Users
SET
  username = ?,
  password = ?,
  email = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM Users
WHERE username = ?;

-- name: ListUsers :many
SELECT * FROM Users
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: CountUsers :one
SELECT count(*) FROM Users;
