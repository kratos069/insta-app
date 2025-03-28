// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  username, hashed_password, full_name, email, profile_picture, bio
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING user_id, username, hashed_password, full_name, profile_picture, bio, email, password_changed_at, created_at
`

type CreateUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	ProfilePicture string `json:"profile_picture"`
	Bio            string `json:"bio"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
		arg.ProfilePicture,
		arg.Bio,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.ProfilePicture,
		&i.Bio,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, userID int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, userID)
	return err
}

const getUserByID = `-- name: GetUserByID :one
SELECT user_id, username, hashed_password, full_name, profile_picture, bio, email, password_changed_at, created_at FROM users
WHERE user_id = $1 LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, userID int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.ProfilePicture,
		&i.Bio,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT user_id, username, hashed_password, full_name, profile_picture, bio, email, password_changed_at, created_at FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.ProfilePicture,
		&i.Bio,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
hashed_password = COALESCE($1, hashed_password),
password_changed_at = COALESCE($2, password_changed_at),
full_name = COALESCE($3, full_name),
email = COALESCE($4, email)
WHERE
username = $5
RETURNING user_id, username, hashed_password, full_name, profile_picture, bio, email, password_changed_at, created_at
`

type UpdateUserParams struct {
	HashedPassword    sql.NullString `json:"hashed_password"`
	PasswordChangedAt sql.NullTime   `json:"password_changed_at"`
	FullName          sql.NullString `json:"full_name"`
	Email             sql.NullString `json:"email"`
	Username          string         `json:"username"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.HashedPassword,
		arg.PasswordChangedAt,
		arg.FullName,
		arg.Email,
		arg.Username,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.ProfilePicture,
		&i.Bio,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}
