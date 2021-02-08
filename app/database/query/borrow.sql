-- name: CreateBorrowedBook :one
INSERT INTO borrowed_books (
  book,
  friend
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetBorrowedBook :one
SELECT * FROM borrowed_books
WHERE id = $1 LIMIT 1;

-- name: ListBorrowedBooks :many
SELECT * FROM borrowed_books
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateBorrowedBook :exec
UPDATE borrowed_books
SET returned_date = NOW()
WHERE id = $1;

-- name: DeleteBorrowedBook :exec
DELETE  FROM borrowed_books WHERE id = $1;
