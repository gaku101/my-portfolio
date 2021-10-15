// Code generated by sqlc. DO NOT EDIT.
// source: post.sql

package db

import (
	"context"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (
    author,
    title,
    book_author,
    book_image,
    book_page,
    book_page_read
  )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, author, title, created_at, updated_at, book_author, book_image, book_page, book_page_read
`

type CreatePostParams struct {
	Author       string `json:"author"`
	Title        string `json:"title"`
	BookAuthor   string `json:"book_author"`
	BookImage    string `json:"book_image"`
	BookPage     int16  `json:"book_page"`
	BookPageRead int16  `json:"book_page_read"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.Author,
		arg.Title,
		arg.BookAuthor,
		arg.BookImage,
		arg.BookPage,
		arg.BookPageRead,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Author,
		&i.Title,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.BookAuthor,
		&i.BookImage,
		&i.BookPage,
		&i.BookPageRead,
	)
	return i, err
}

const deletePost = `-- name: DeletePost :one
DELETE FROM posts
WHERE id = $1
RETURNING id, author, title, created_at, updated_at, book_author, book_image, book_page, book_page_read
`

func (q *Queries) DeletePost(ctx context.Context, id int64) (Post, error) {
	row := q.db.QueryRowContext(ctx, deletePost, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Author,
		&i.Title,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.BookAuthor,
		&i.BookImage,
		&i.BookPage,
		&i.BookPageRead,
	)
	return i, err
}

const getPost = `-- name: GetPost :one
SELECT id, author, title, created_at, updated_at, book_author, book_image, book_page, book_page_read
FROM posts
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetPost(ctx context.Context, id int64) (Post, error) {
	row := q.db.QueryRowContext(ctx, getPost, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Author,
		&i.Title,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.BookAuthor,
		&i.BookImage,
		&i.BookPage,
		&i.BookPageRead,
	)
	return i, err
}

const listMyPosts = `-- name: ListMyPosts :many
SELECT id, author, title, created_at, updated_at, book_author, book_image, book_page, book_page_read
FROM posts
WHERE author = $1
ORDER BY id
LIMIT $2 OFFSET $3
`

type ListMyPostsParams struct {
	Author string `json:"author"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (q *Queries) ListMyPosts(ctx context.Context, arg ListMyPostsParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, listMyPosts, arg.Author, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Author,
			&i.Title,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.BookAuthor,
			&i.BookImage,
			&i.BookPage,
			&i.BookPageRead,
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

const listPosts = `-- name: ListPosts :many
SELECT id, author, title, created_at, updated_at, book_author, book_image, book_page, book_page_read
FROM posts
WHERE NOT author = $1
ORDER BY id DESC
LIMIT $2 OFFSET $3
`

type ListPostsParams struct {
	Author string `json:"author"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (q *Queries) ListPosts(ctx context.Context, arg ListPostsParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, listPosts, arg.Author, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Author,
			&i.Title,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.BookAuthor,
			&i.BookImage,
			&i.BookPage,
			&i.BookPageRead,
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

const updatePost = `-- name: UpdatePost :one
UPDATE posts
SET title = $2
WHERE id = $1
RETURNING id, author, title, created_at, updated_at, book_author, book_image, book_page, book_page_read
`

type UpdatePostParams struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, updatePost, arg.ID, arg.Title)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Author,
		&i.Title,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.BookAuthor,
		&i.BookImage,
		&i.BookPage,
		&i.BookPageRead,
	)
	return i, err
}
