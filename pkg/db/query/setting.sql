
-- name: CreateSetting :execresult
INSERT INTO Settings (
  user_id,
  theme,
  currency
) VALUES (
  ?, ?, ?
);

-- name: GetSetting :one
SELECT * FROM Settings
WHERE id = ?;

-- name: UpdateSetting :exec
UPDATE Settings 
SET
  user_id = ?,
  theme = ?,
  currency = ?
WHERE id = ?;

-- name: DeleteSetting :exec
DELETE FROM Settings
WHERE id = ?;

-- name: ListSettings :many
SELECT * FROM Settings
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: CountSettings :one
SELECT count(*) FROM Settings;
