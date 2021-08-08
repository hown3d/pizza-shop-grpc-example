// Code generated by sqlc. DO NOT EDIT.
// source: queries.sql

package storage

import (
	"context"

	"github.com/jackc/pgtype"
)

const createBake = `-- name: CreateBake :one
INSERT INTO bake (
  name, bake_status 
) VALUES (
  $1, $2
)
RETURNING id, create_time, name, bake_status
`

type CreateBakeParams struct {
	Name       string
	BakeStatus StatusEnum
}

func (q *Queries) CreateBake(ctx context.Context, arg CreateBakeParams) (Bake, error) {
	row := q.db.QueryRowContext(ctx, createBake, arg.Name, arg.BakeStatus)
	var i Bake
	err := row.Scan(
		&i.ID,
		&i.CreateTime,
		&i.Name,
		&i.BakeStatus,
	)
	return i, err
}

const deleteBake = `-- name: DeleteBake :exec
DELETE FROM bake
WHERE id = $1
`

func (q *Queries) DeleteBake(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteBake, id)
	return err
}

const getAllBake = `-- name: GetAllBake :many
SELECT id, create_time, name, bake_status FROM bake
`

func (q *Queries) GetAllBake(ctx context.Context) ([]Bake, error) {
	rows, err := q.db.QueryContext(ctx, getAllBake)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Bake
	for rows.Next() {
		var i Bake
		if err := rows.Scan(
			&i.ID,
			&i.CreateTime,
			&i.Name,
			&i.BakeStatus,
		); err != nil {
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

const getBake = `-- name: GetBake :one
SELECT id, create_time, name, bake_status FROM bake
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetBake(ctx context.Context, id pgtype.UUID) (Bake, error) {
	row := q.db.QueryRowContext(ctx, getBake, id)
	var i Bake
	err := row.Scan(
		&i.ID,
		&i.CreateTime,
		&i.Name,
		&i.BakeStatus,
	)
	return i, err
}

const updateBake = `-- name: UpdateBake :exec
UPDATE bake SET bake_status = $2
WHERE id = $1
`

type UpdateBakeParams struct {
	ID         pgtype.UUID
	BakeStatus StatusEnum
}

func (q *Queries) UpdateBake(ctx context.Context, arg UpdateBakeParams) error {
	_, err := q.db.ExecContext(ctx, updateBake, arg.ID, arg.BakeStatus)
	return err
}