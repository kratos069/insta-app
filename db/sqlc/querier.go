// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CountLikesByPost(ctx context.Context, postID int64) ([]int64, error)
	CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error)
	CreateFollow(ctx context.Context, arg CreateFollowParams) error
	CreateLike(ctx context.Context, arg CreateLikeParams) (Like, error)
	CreatePost(ctx context.Context, arg CreatePostParams) (Post, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateVerifyEmail(ctx context.Context, arg CreateVerifyEmailParams) (VerifyEmail, error)
	DeleteComment(ctx context.Context, commentID int64) error
	DeleteFollow(ctx context.Context, arg DeleteFollowParams) error
	DeleteLike(ctx context.Context, arg DeleteLikeParams) error
	DeletePost(ctx context.Context, postID int64) error
	DeleteUser(ctx context.Context, userID int64) error
	GetCommentByID(ctx context.Context, commentID int64) (Comment, error)
	GetPostByID(ctx context.Context, postID int64) (Post, error)
	GetSessionByID(ctx context.Context, id uuid.UUID) (Session, error)
	GetUserByID(ctx context.Context, userID int64) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	ListCommentsByPost(ctx context.Context, postID int64) ([]Comment, error)
	ListFollowers(ctx context.Context, followingID int64) ([]ListFollowersRow, error)
	ListFollowing(ctx context.Context, followerID int64) ([]ListFollowingRow, error)
	ListPostsByUser(ctx context.Context, userID int64) ([]Post, error)
	UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateVerifyEmail(ctx context.Context, arg UpdateVerifyEmailParams) (VerifyEmail, error)
}

var _ Querier = (*Queries)(nil)
