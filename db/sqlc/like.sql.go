// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: like.sql

package db

import (
	"context"
)

const countLikesByPost = `-- name: CountLikesByPost :many
SELECT COUNT(*) AS count
FROM likes
WHERE post_id = $1
`

func (q *Queries) CountLikesByPost(ctx context.Context, postID int64) ([]int64, error) {
	rows, err := q.db.QueryContext(ctx, countLikesByPost, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var count int64
		if err := rows.Scan(&count); err != nil {
			return nil, err
		}
		items = append(items, count)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const createLike = `-- name: CreateLike :one
INSERT INTO likes (
  user_id, post_id
) VALUES (
  $1, $2
)
RETURNING like_id, post_id, user_id, created_at
`

type CreateLikeParams struct {
	UserID int64 `json:"user_id"`
	PostID int64 `json:"post_id"`
}

func (q *Queries) CreateLike(ctx context.Context, arg CreateLikeParams) (Like, error) {
	row := q.db.QueryRowContext(ctx, createLike, arg.UserID, arg.PostID)
	var i Like
	err := row.Scan(
		&i.LikeID,
		&i.PostID,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteLike = `-- name: DeleteLike :exec
DELETE FROM likes
WHERE post_id = $1 AND user_id = $2
`

type DeleteLikeParams struct {
	PostID int64 `json:"post_id"`
	UserID int64 `json:"user_id"`
}

func (q *Queries) DeleteLike(ctx context.Context, arg DeleteLikeParams) error {
	_, err := q.db.ExecContext(ctx, deleteLike, arg.PostID, arg.UserID)
	return err
}
