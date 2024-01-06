
-- name: CreateReport :execresult
INSERT INTO Reports (
  user_id,
  title
) VALUES (
  ?, ?
);

-- name: GetReport :one
SELECT * FROM Reports
WHERE id = ?;

-- name: UpdateReport :exec
UPDATE Reports
SET
  user_id = ?,
  title = ?
WHERE id = ?;

-- name: DeleteReport :exec
DELETE FROM Reports
WHERE id = ?;

-- name: ListReports :many
SELECT * FROM Reports
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: CountReports :one
SELECT count(*) FROM Reports;
