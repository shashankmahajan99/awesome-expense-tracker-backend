
-- name: CreateSetting :execresult
INSERT INTO settings (
  user_id,
  theme,
  currency
) VALUES (
  ?, ?, ?
);

-- name: GetSetting :one
SELECT * FROM settings
WHERE id = ?;

-- name: UpdateSetting :exec
UPDATE settings 
SET
  user_id = ?,
  theme = ?,
  currency = ?
WHERE id = ?;

-- name: DeleteSetting :exec
DELETE FROM settings
WHERE id = ?;

-- name: ListSettings :many
SELECT * FROM settings
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: CountSettings :one
SELECT count(*) FROM settings;
