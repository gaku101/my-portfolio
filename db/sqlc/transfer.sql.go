// Code generated by sqlc. DO NOT EDIT.
// source: transfer.sql

package db

import (
	"context"
	"time"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO transfers (from_user_id, to_user_id, amount)
VALUES ($1, $2, $3)
RETURNING id, from_user_id, to_user_id, amount, created_at
`

type CreateTransferParams struct {
	FromUserID int64 `json:"from_user_id"`
	ToUserID   int64 `json:"to_user_id"`
	Amount     int64 `json:"amount"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, createTransfer, arg.FromUserID, arg.ToUserID, arg.Amount)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromUserID,
		&i.ToUserID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getTransfer = `-- name: GetTransfer :one
SELECT id, from_user_id, to_user_id, amount, created_at
FROM transfers
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetTransfer(ctx context.Context, id int64) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTransfer, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromUserID,
		&i.ToUserID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listTransfers = `-- name: ListTransfers :many
SELECT transfers.id,
  from_user_id,
  to_user_id,
  amount,
  transfers.created_at,
  users.username
FROM transfers
  JOIN users ON from_user_id = users.id
  AND to_user_id = $1
ORDER BY id DESC
LIMIT $2 OFFSET $3
`

type ListTransfersParams struct {
	ToUserID int64 `json:"to_user_id"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

type ListTransfersRow struct {
	ID         int64     `json:"id"`
	FromUserID int64     `json:"from_user_id"`
	ToUserID   int64     `json:"to_user_id"`
	Amount     int64     `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
	Username   string    `json:"username"`
}

func (q *Queries) ListTransfers(ctx context.Context, arg ListTransfersParams) ([]ListTransfersRow, error) {
	rows, err := q.db.QueryContext(ctx, listTransfers, arg.ToUserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListTransfersRow{}
	for rows.Next() {
		var i ListTransfersRow
		if err := rows.Scan(
			&i.ID,
			&i.FromUserID,
			&i.ToUserID,
			&i.Amount,
			&i.CreatedAt,
			&i.Username,
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
