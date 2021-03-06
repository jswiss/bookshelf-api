// Code generated by sqlc. DO NOT EDIT.
// source: borrow.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createBorrowedBook = `-- name: CreateBorrowedBook :one
INSERT INTO borrowed_books (
  book_id,
  friend_id
) VALUES (
  $1, $2
) RETURNING id, book_id, friend_id, borrowed_date, returned_date, created_at, updated_at
`

type CreateBorrowedBookParams struct {
	BookID   int32 `json:"book_id"`
	FriendID int32 `json:"friend_id"`
}

func (q *Queries) CreateBorrowedBook(ctx context.Context, arg CreateBorrowedBookParams) (BorrowedBook, error) {
	row := q.db.QueryRowContext(ctx, createBorrowedBook, arg.BookID, arg.FriendID)
	var i BorrowedBook
	err := row.Scan(
		&i.ID,
		&i.BookID,
		&i.FriendID,
		&i.BorrowedDate,
		&i.ReturnedDate,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteBorrowedBook = `-- name: DeleteBorrowedBook :exec
DELETE  FROM borrowed_books WHERE id = $1
`

func (q *Queries) DeleteBorrowedBook(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteBorrowedBook, id)
	return err
}

const getBorrowedBook = `-- name: GetBorrowedBook :one
SELECT bb.id, bb.book_id, bb.friend_id, bb.borrowed_date, bb.returned_date, b.title, b.author, f.full_name FROM borrowed_books bb
INNER JOIN books b ON bb.book_id = b.id
INNER JOIN friends f ON bb.friend_id = f.id
WHERE bb.id = $1 LIMIT 1
`

type GetBorrowedBookRow struct {
	ID           int32        `json:"id"`
	BookID       int32        `json:"book_id"`
	FriendID     int32        `json:"friend_id"`
	BorrowedDate time.Time    `json:"borrowed_date"`
	ReturnedDate sql.NullTime `json:"returned_date"`
	Title        string       `json:"title"`
	Author       string       `json:"author"`
	FullName     string       `json:"full_name"`
}

func (q *Queries) GetBorrowedBook(ctx context.Context, id int32) (GetBorrowedBookRow, error) {
	row := q.db.QueryRowContext(ctx, getBorrowedBook, id)
	var i GetBorrowedBookRow
	err := row.Scan(
		&i.ID,
		&i.BookID,
		&i.FriendID,
		&i.BorrowedDate,
		&i.ReturnedDate,
		&i.Title,
		&i.Author,
		&i.FullName,
	)
	return i, err
}

const listBorrowedBooks = `-- name: ListBorrowedBooks :many
SELECT bb.id, bb.book_id, bb.friend_id, bb.borrowed_date, bb.returned_date, b.title, b.author, f.full_name FROM borrowed_books bb
INNER JOIN books b ON bb.book_id = b.id
INNER JOIN friends f ON bb.friend_id = f.id
ORDER BY bb.id
LIMIT $1
OFFSET $2
`

type ListBorrowedBooksParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type ListBorrowedBooksRow struct {
	ID           int32        `json:"id"`
	BookID       int32        `json:"book_id"`
	FriendID     int32        `json:"friend_id"`
	BorrowedDate time.Time    `json:"borrowed_date"`
	ReturnedDate sql.NullTime `json:"returned_date"`
	Title        string       `json:"title"`
	Author       string       `json:"author"`
	FullName     string       `json:"full_name"`
}

func (q *Queries) ListBorrowedBooks(ctx context.Context, arg ListBorrowedBooksParams) ([]ListBorrowedBooksRow, error) {
	rows, err := q.db.QueryContext(ctx, listBorrowedBooks, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListBorrowedBooksRow{}
	for rows.Next() {
		var i ListBorrowedBooksRow
		if err := rows.Scan(
			&i.ID,
			&i.BookID,
			&i.FriendID,
			&i.BorrowedDate,
			&i.ReturnedDate,
			&i.Title,
			&i.Author,
			&i.FullName,
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

const updateBorrowedBook = `-- name: UpdateBorrowedBook :exec
UPDATE borrowed_books
SET returned_date = NOW()
WHERE id = $1
`

func (q *Queries) UpdateBorrowedBook(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, updateBorrowedBook, id)
	return err
}
