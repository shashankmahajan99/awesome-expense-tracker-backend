-- name: CreateUser :execresult
INSERT INTO users (
  username,
  password,
  email
) VALUES (
  ?, ?, ?
);

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = ?;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = ?;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = ?;

-- name: DeleteUser :execrows
DELETE FROM users
WHERE username = ?;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: CountUsers :one
SELECT count(*) FROM users;

-- name: UpdateUserPassword :execresult
UPDATE users
SET
  password = ?
WHERE email = ?;

-- name: UpdateUserUsername :execresult
UPDATE users
SET
  username = ?
WHERE email = ?;