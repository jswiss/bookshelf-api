// Code generated by sqlc. DO NOT EDIT.
// source: friend.sql

package database

import (
	"context"
)

const createFriend = `-- name: CreateFriend :one

INSERT INTO friends (full_name, photo)
VALUES ($1,
        $2) RETURNING id, full_name, photo, created_at, updated_at
`

type CreateFriendParams struct {
	FullName string `json:"full_name"`
	Photo    string `json:"photo"`
}

func (q *Queries) CreateFriend(ctx context.Context, arg CreateFriendParams) (Friend, error) {
	row := q.db.QueryRowContext(ctx, createFriend, arg.FullName, arg.Photo)
	var i Friend
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Photo,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteFriend = `-- name: DeleteFriend :exec

DELETE
FROM friends
WHERE id = $1
`

func (q *Queries) DeleteFriend(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteFriend, id)
	return err
}

const getFriend = `-- name: GetFriend :one

SELECT id, full_name, photo, created_at, updated_at
FROM friends
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetFriend(ctx context.Context, id int32) (Friend, error) {
	row := q.db.QueryRowContext(ctx, getFriend, id)
	var i Friend
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Photo,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listFriends = `-- name: ListFriends :many

SELECT id, full_name, photo, created_at, updated_at
FROM friends
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListFriendsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListFriends(ctx context.Context, arg ListFriendsParams) ([]Friend, error) {
	rows, err := q.db.QueryContext(ctx, listFriends, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Friend{}
	for rows.Next() {
		var i Friend
		if err := rows.Scan(
			&i.ID,
			&i.FullName,
			&i.Photo,
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

const updateFriend = `-- name: UpdateFriend :exec

UPDATE friends
SET full_name = $2,
    photo = $3
WHERE id = $1 RETURNING friends.id, friends.full_name, friends.photo, friends.created_at, friends.updated_at
`

type UpdateFriendParams struct {
	ID       int32  `json:"id"`
	FullName string `json:"full_name"`
	Photo    string `json:"photo"`
}

func (q *Queries) UpdateFriend(ctx context.Context, arg UpdateFriendParams) error {
	_, err := q.db.ExecContext(ctx, updateFriend, arg.ID, arg.FullName, arg.Photo)
	return err
}
