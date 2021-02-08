-- name: CreateBook :one
INSERT INTO books (
  title,
  author
  ) VALUES (
  $1, $2
) RETURNING *;

-- name: GetBook :one
SELECT * FROM books
WHERE id = $1 LIMIT 1;

-- name: ListBooks :many
SELECT * FROM books
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateBook :exec
UPDATE books
SET
  title = $2,
  author = $3
WHERE id = $1;

-- name: DeleteBook :exec
DELETE FROM books WHERE id = $1;
