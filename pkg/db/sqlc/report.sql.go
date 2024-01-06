// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: report.sql

package db

import (
	"context"
	"database/sql"
)

const countReports = `-- name: CountReports :one
SELECT count(*) FROM Reports
`

func (q *Queries) CountReports(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countReports)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createReport = `-- name: CreateReport :execresult
INSERT INTO Reports (
  user_id,
  title
) VALUES (
  ?, ?
)
`

type CreateReportParams struct {
	UserID sql.NullInt32  `json:"user_id"`
	Title  sql.NullString `json:"title"`
}

func (q *Queries) CreateReport(ctx context.Context, arg CreateReportParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createReport, arg.UserID, arg.Title)
}

const deleteReport = `-- name: DeleteReport :exec
DELETE FROM Reports
WHERE id = ?
`

func (q *Queries) DeleteReport(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteReport, id)
	return err
}

const getReport = `-- name: GetReport :one
SELECT id, user_id, title FROM Reports
WHERE id = ?
`

func (q *Queries) GetReport(ctx context.Context, id int32) (Report, error) {
	row := q.db.QueryRowContext(ctx, getReport, id)
	var i Report
	err := row.Scan(&i.ID, &i.UserID, &i.Title)
	return i, err
}

const listReports = `-- name: ListReports :many
SELECT id, user_id, title FROM Reports
ORDER BY id
LIMIT ?
OFFSET ?
`

type ListReportsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListReports(ctx context.Context, arg ListReportsParams) ([]Report, error) {
	rows, err := q.db.QueryContext(ctx, listReports, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Report{}
	for rows.Next() {
		var i Report
		if err := rows.Scan(&i.ID, &i.UserID, &i.Title); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateReport = `-- name: UpdateReport :exec
UPDATE Reports
SET
  user_id = ?,
  title = ?
WHERE id = ?
`

type UpdateReportParams struct {
	UserID sql.NullInt32  `json:"user_id"`
	Title  sql.NullString `json:"title"`
	ID     int32          `json:"id"`
}

func (q *Queries) UpdateReport(ctx context.Context, arg UpdateReportParams) error {
	_, err := q.db.ExecContext(ctx, updateReport, arg.UserID, arg.Title, arg.ID)
	return err
}
