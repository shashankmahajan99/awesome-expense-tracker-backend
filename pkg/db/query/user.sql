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

-- name: GetUserByEmail :one
SELECT * FROM Users
WHERE email = ?;

-- name: GetUserByID :one
SELECT * FROM Users
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

-- name: UpdateUserPassword :exec
UPDATE Users
SET
  password = ?
WHERE email = ?;

-- name: UpdateUserUsername :exec
UPDATE Users
SET
  username = ?
WHERE email = ?;