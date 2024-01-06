
-- name: CreateProfile :execresult
INSERT INTO Profiles (
  user_id,
  bio
) VALUES (
  ?, ?
);

-- name: GetProfile :one
SELECT * FROM Profiles
WHERE id = ?;

-- name: UpdateProfile :exec
UPDATE Profiles
SET
  user_id = ?,
  bio = ?
WHERE id = ?;

-- name: DeleteProfile :exec
DELETE FROM Profiles
WHERE id = ?;

-- name: ListProfiles :many
SELECT * FROM Profiles
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: CountProfiles :one
SELECT count(*) FROM Profiles;
