
-- name: CreateReport :execresult
INSERT INTO reports (
  user_id,
  title
) VALUES (
  ?, ?
);

-- name: GetReport :one
SELECT * FROM reports
WHERE id = ?;

-- name: UpdateReport :exec
UPDATE reports
SET
  user_id = ?,
  title = ?
WHERE id = ?;

-- name: DeleteReport :exec
DELETE FROM reports
WHERE id = ?;

-- name: Listreports :many
SELECT * FROM reports
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: Countreports :one
SELECT count(*) FROM reports;
