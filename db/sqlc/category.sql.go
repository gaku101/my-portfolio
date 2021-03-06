// Code generated by sqlc. DO NOT EDIT.
// source: category.sql

package db

import (
	"context"
)

const createCategory = `-- name: CreateCategory :one
INSERT INTO category (name)
VALUES ($1)
RETURNING id, name
`

func (q *Queries) CreateCategory(ctx context.Context, name string) (Category, error) {
	row := q.db.QueryRowContext(ctx, createCategory, name)
	var i Category
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getCategory = `-- name: GetCategory :one
SELECT id, name
FROM category
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetCategory(ctx context.Context, id int64) (Category, error) {
	row := q.db.QueryRowContext(ctx, getCategory, id)
	var i Category
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const listCategories = `-- name: ListCategories :many
SELECT id, name
FROM category
ORDER BY id
`

func (q *Queries) ListCategories(ctx context.Context) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, listCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Category{}
	for rows.Next() {
		var i Category
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
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
