// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: post.sql

package db

import (
	"context"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (
  user_id, content_url, caption
) VALUES (
  $1, $2, $3
)
RETURNING post_id, user_id, content_url, caption, created_at
`

type CreatePostParams struct {
	UserID     int64  `json:"user_id"`
	ContentUrl string `json:"content_url"`
	Caption    string `json:"caption"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRow(ctx, createPost, arg.UserID, arg.ContentUrl, arg.Caption)
	var i Post
	err := row.Scan(
		&i.PostID,
		&i.UserID,
		&i.ContentUrl,
		&i.Caption,
		&i.CreatedAt,
	)
	return i, err
}

const deletePost = `-- name: DeletePost :exec
DELETE FROM posts
WHERE post_id = $1
`

func (q *Queries) DeletePost(ctx context.Context, postID int64) error {
	_, err := q.db.Exec(ctx, deletePost, postID)
	return err
}

const getPostByID = `-- name: GetPostByID :one
SELECT post_id, user_id, content_url, caption, created_at FROM posts
WHERE post_id = $1 LIMIT 1
`

func (q *Queries) GetPostByID(ctx context.Context, postID int64) (Post, error) {
	row := q.db.QueryRow(ctx, getPostByID, postID)
	var i Post
	err := row.Scan(
		&i.PostID,
		&i.UserID,
		&i.ContentUrl,
		&i.Caption,
		&i.CreatedAt,
	)
	return i, err
}

const listPostsByUser = `-- name: ListPostsByUser :many
SELECT post_id, user_id, content_url, caption, created_at FROM posts
WHERE user_id = $1
ORDER BY post_id
`

func (q *Queries) ListPostsByUser(ctx context.Context, userID int64) ([]Post, error) {
	rows, err := q.db.Query(ctx, listPostsByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.PostID,
			&i.UserID,
			&i.ContentUrl,
			&i.Caption,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePost = `-- name: UpdatePost :one
UPDATE posts
SET caption = $2
WHERE post_id = $1
RETURNING post_id, user_id, content_url, caption, created_at
`

type UpdatePostParams struct {
	PostID  int64  `json:"post_id"`
	Caption string `json:"caption"`
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error) {
	row := q.db.QueryRow(ctx, updatePost, arg.PostID, arg.Caption)
	var i Post
	err := row.Scan(
		&i.PostID,
		&i.UserID,
		&i.ContentUrl,
		&i.Caption,
		&i.CreatedAt,
	)
	return i, err
}
