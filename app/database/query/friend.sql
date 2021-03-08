-- name: CreateFriend :one

INSERT INTO friends (full_name, photo)
VALUES ($1,
        $2) RETURNING *;

-- name: GetFriend :one

SELECT *
FROM friends
WHERE id = $1
LIMIT 1;

-- name: UpdateFriend :exec

UPDATE friends
SET full_name = $2,
    photo = $3
WHERE id = $1 RETURNING friends.*;

-- name: ListFriends :many

SELECT *
FROM friends
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteFriend :exec

DELETE
FROM friends
WHERE id = $1;
