package db

import (
	"context"
	"testing"

	"github.com/insta-app/util"
	"github.com/stretchr/testify/require"
)

func createRandomComment(t *testing.T) Comment {
	post := createRandomPost(t)

	arg := CreateCommentParams{
		UserID:  post.UserID,
		PostID:  post.PostID,
		Content: util.RandomString(15),
	}

	comment, err := testStore.CreateComment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, comment)

	require.Equal(t, arg.UserID, comment.UserID)
	require.Equal(t, arg.PostID, comment.PostID)
	require.Equal(t, arg.Content, comment.Content)

	return comment
}

func TestCreateComment(t *testing.T) {
	createRandomComment(t)
}

func TestListCommentsByPost(t *testing.T) {
	comment := createRandomComment(t)

	comments, err := testStore.ListCommentsByPost(context.Background(), comment.PostID)
	require.NoError(t, err)
	require.NotEmpty(t, comments)
}

func TestDeleteComment(t *testing.T) {
	comment := createRandomComment(t)

	err := testStore.DeleteComment(context.Background(), comment.CommentID)
	require.NoError(t, err)
}
