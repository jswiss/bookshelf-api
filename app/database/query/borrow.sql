-- name: CreateBorrowedBook :one
INSERT INTO borrowed_books (
  book_id,
  friend_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetBorrowedBook :one
SELECT bb.id, bb.book_id, bb.friend_id, bb.borrowed_date, bb.returned_date, b.title, b.author, f.full_name FROM borrowed_books bb
INNER JOIN books b ON bb.book_id = b.id
INNER JOIN friends f ON bb.friend_id = f.id
WHERE bb.id = $1 LIMIT 1;

-- name: ListBorrowedBooks :many
SELECT bb.id, bb.book_id, bb.friend_id, bb.borrowed_date, bb.returned_date, b.title, b.author, f.full_name FROM borrowed_books bb
INNER JOIN books b ON bb.book_id = b.id
INNER JOIN friends f ON bb.friend_id = f.id
ORDER BY bb.id
LIMIT $1
OFFSET $2;

-- name: UpdateBorrowedBook :exec
UPDATE borrowed_books
SET returned_date = NOW()
WHERE id = $1;

-- name: DeleteBorrowedBook :exec
DELETE  FROM borrowed_books WHERE id = $1;
