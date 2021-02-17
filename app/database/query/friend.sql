-- name: CreateFriend :one
INSERT INTO friends (
  full_name,
  photo
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetFriend :one
SELECT * FROM friends
WHERE id = $1 LIMIT 1;

-- name: ListFriends :many
SELECT * FROM friends
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteFriend :exec
DELETE  FROM friends WHERE id = $1;
