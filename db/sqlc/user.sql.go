// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (username, hashed_password, email, profile, image)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, username, hashed_password, email, profile, image, password_changed_at, created_at
`

type CreateUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	Email          string `json:"email"`
	Profile        string `json:"profile"`
	Image          string `json:"image"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.HashedPassword,
		arg.Email,
		arg.Profile,
		arg.Image,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.Email,
		&i.Profile,
		&i.Image,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, username, hashed_password, email, profile, image, password_changed_at, created_at
FROM users
WHERE username = $1
LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.Email,
		&i.Profile,
		&i.Image,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, username, hashed_password, email, profile, image, password_changed_at, created_at
FROM users
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetUserById(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.Email,
		&i.Profile,
		&i.Image,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUserImage = `-- name: GetUserImage :one
SELECT image
FROM users
WHERE username = $1
LIMIT 1
`

func (q *Queries) GetUserImage(ctx context.Context, username string) (string, error) {
	row := q.db.QueryRowContext(ctx, getUserImage, username)
	var image string
	err := row.Scan(&image)
	return image, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET profile = $2
WHERE id = $1
RETURNING id, username, hashed_password, email, profile, image, password_changed_at, created_at
`

type UpdateUserParams struct {
	ID      int64  `json:"id"`
	Profile string `json:"profile"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser, arg.ID, arg.Profile)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.Email,
		&i.Profile,
		&i.Image,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateUserImage = `-- name: UpdateUserImage :one
UPDATE users
SET image = $2
WHERE username = $1
RETURNING id, username, hashed_password, email, profile, image, password_changed_at, created_at
`

type UpdateUserImageParams struct {
	Username string `json:"username"`
	Image    string `json:"image"`
}

func (q *Queries) UpdateUserImage(ctx context.Context, arg UpdateUserImageParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserImage, arg.Username, arg.Image)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.Email,
		&i.Profile,
		&i.Image,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}
