// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: follow.sql

package db

import (
	"context"
	"time"
)

const createFollow = `-- name: CreateFollow :exec
INSERT INTO follows (
  follower_id,
  following_id
) VALUES (
  $1, $2
)
ON CONFLICT DO NOTHING
`

type CreateFollowParams struct {
	FollowerID  int64 `json:"follower_id"`
	FollowingID int64 `json:"following_id"`
}

func (q *Queries) CreateFollow(ctx context.Context, arg CreateFollowParams) error {
	_, err := q.db.Exec(ctx, createFollow, arg.FollowerID, arg.FollowingID)
	return err
}

const deleteFollow = `-- name: DeleteFollow :exec
DELETE FROM follows
WHERE follower_id = $1 AND following_id = $2
`

type DeleteFollowParams struct {
	FollowerID  int64 `json:"follower_id"`
	FollowingID int64 `json:"following_id"`
}

func (q *Queries) DeleteFollow(ctx context.Context, arg DeleteFollowParams) error {
	_, err := q.db.Exec(ctx, deleteFollow, arg.FollowerID, arg.FollowingID)
	return err
}

const listFollowers = `-- name: ListFollowers :many
SELECT
  u.user_id,
  u.username,
  u.full_name,
  u.profile_picture,
  u.bio,
  u.email,
  u.password_changed_at,
  u.created_at
FROM follows f
JOIN users u ON f.follower_id = u.user_id
WHERE f.following_id = $1
`

type ListFollowersRow struct {
	UserID            int64     `json:"user_id"`
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	ProfilePicture    string    `json:"profile_picture"`
	Bio               string    `json:"bio"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func (q *Queries) ListFollowers(ctx context.Context, followingID int64) ([]ListFollowersRow, error) {
	rows, err := q.db.Query(ctx, listFollowers, followingID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListFollowersRow{}
	for rows.Next() {
		var i ListFollowersRow
		if err := rows.Scan(
			&i.UserID,
			&i.Username,
			&i.FullName,
			&i.ProfilePicture,
			&i.Bio,
			&i.Email,
			&i.PasswordChangedAt,
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

const listFollowing = `-- name: ListFollowing :many
SELECT
  u.user_id,
  u.username,
  u.full_name,
  u.profile_picture,
  u.bio,
  u.email,
  u.password_changed_at,
  u.created_at
FROM follows f
JOIN users u ON f.following_id = u.user_id
WHERE f.follower_id = $1
`

type ListFollowingRow struct {
	UserID            int64     `json:"user_id"`
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	ProfilePicture    string    `json:"profile_picture"`
	Bio               string    `json:"bio"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func (q *Queries) ListFollowing(ctx context.Context, followerID int64) ([]ListFollowingRow, error) {
	rows, err := q.db.Query(ctx, listFollowing, followerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListFollowingRow{}
	for rows.Next() {
		var i ListFollowingRow
		if err := rows.Scan(
			&i.UserID,
			&i.Username,
			&i.FullName,
			&i.ProfilePicture,
			&i.Bio,
			&i.Email,
			&i.PasswordChangedAt,
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
