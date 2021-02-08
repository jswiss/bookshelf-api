-- name: CreateBook :one
INSERT INTO books (
  title,
  author,
  cover_image
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetBook :one
SELECT * FROM books
WHERE id = $1 LIMIT 1;

-- name: ListBooks :many
SELECT * FROM books
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteBook :exec
DELETE  FROM books WHERE id = $1;
