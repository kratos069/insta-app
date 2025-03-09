package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomLike(t *testing.T) Like {
	post := createRandomPost(t)

	arg := CreateLikeParams{
		UserID: post.UserID,
		PostID: post.PostID,
	}

	like, err := testQueries.CreateLike(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, like)

	require.Equal(t, arg.UserID, like.UserID)
	require.Equal(t, arg.PostID, like.PostID)

	return like
}

func TestCreateLike(t *testing.T) {
	createRandomLike(t)
}

func TestCountLikesByPost(t *testing.T) {
	like := createRandomLike(t)

	_, err := testQueries.CountLikesByPost(context.Background(), like.PostID)
	require.NoError(t, err)
}

func TestDeleteLike(t *testing.T) {
	like := createRandomLike(t)

	arg := DeleteLikeParams{
		PostID: like.PostID,
		UserID: like.UserID,
	}

	err := testQueries.DeleteLike(context.Background(), arg)
	require.NoError(t, err)
}