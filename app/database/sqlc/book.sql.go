// Code generated by sqlc. DO NOT EDIT.
// source: book.sql

package database

import (
	"context"
)

const createBook = `-- name: CreateBook :one

INSERT INTO books (title, author, cover_image)
VALUES ($1,
        $2,
        $3) RETURNING id, title, author, cover_image, in_stock, created_at, updated_at
`

type CreateBookParams struct {
	Title      string `json:"title"`
	Author     string `json:"author"`
	CoverImage string `json:"cover_image"`
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) (Book, error) {
	row := q.db.QueryRowContext(ctx, createBook, arg.Title, arg.Author, arg.CoverImage)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Author,
		&i.CoverImage,
		&i.InStock,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteBook = `-- name: DeleteBook :exec

DELETE
FROM books
WHERE id = $1
`

func (q *Queries) DeleteBook(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteBook, id)
	return err
}

const getBook = `-- name: GetBook :one

SELECT id, title, author, cover_image, in_stock, created_at, updated_at
FROM books
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetBook(ctx context.Context, id int32) (Book, error) {
	row := q.db.QueryRowContext(ctx, getBook, id)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Author,
		&i.CoverImage,
		&i.InStock,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listBooks = `-- name: ListBooks :many

SELECT id, title, author, cover_image, in_stock, created_at, updated_at
FROM books
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListBooksParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListBooks(ctx context.Context, arg ListBooksParams) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, listBooks, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Book{}
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Author,
			&i.CoverImage,
			&i.InStock,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateBook = `-- name: UpdateBook :exec

UPDATE books
SET title = $2,
    author = $3,
    cover_image = $4
WHERE id = $1 RETURNING books.id, books.title, books.author, books.cover_image, books.in_stock, books.created_at, books.updated_at
`

type UpdateBookParams struct {
	ID         int32  `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	CoverImage string `json:"cover_image"`
}

func (q *Queries) UpdateBook(ctx context.Context, arg UpdateBookParams) error {
	_, err := q.db.ExecContext(ctx, updateBook,
		arg.ID,
		arg.Title,
		arg.Author,
		arg.CoverImage,
	)
	return err
}
