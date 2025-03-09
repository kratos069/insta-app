package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/insta-app/util"
	"github.com/stretchr/testify/require"
)

func createRandomPost(t *testing.T) Post {
	user := createRandomUser(t)

	arg := CreatePostParams{
		UserID:     user.UserID,
		ContentUrl: util.RandomString(10),
		Caption:    util.RandomString(15),
	}

	post, err := testQueries.CreatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post)

	require.Equal(t, user.UserID, post.UserID)

	return post
}

func TestCreatePost(t *testing.T) {
	createRandomPost(t)
}

func TestGetPostByID(t *testing.T) {
	post1 := createRandomPost(t)

	post2, err := testQueries.GetPostByID(context.Background(), post1.PostID)
	require.NoError(t, err)
	require.NotEmpty(t, post2)

	require.Equal(t, post1.PostID, post2.PostID)
	require.Equal(t, post1.UserID, post2.UserID)
	require.Equal(t, post1.Caption, post2.Caption)
	require.Equal(t, post1.ContentUrl, post2.ContentUrl)
	require.WithinDuration(t, post1.CreatedAt, post2.CreatedAt, time.Second)
}

func TestListPostsByUser(t *testing.T) {
	post := createRandomPost(t)

	posts, err := testQueries.ListPostsByUser(context.Background(), post.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, posts)
}

func TestUpdatePost(t *testing.T) {
	post := createRandomPost(t)

	arg := UpdatePostParams{
		PostID:  post.PostID,
		Caption: util.RandomString(9),
	}

	updatedPost, err := testQueries.UpdatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedPost)

	require.Equal(t, arg.Caption, updatedPost.Caption)
}

func TestDeletePost(t *testing.T) {
	post := createRandomPost(t)

	err := testQueries.DeletePost(context.Background(), post.PostID)
	require.NoError(t, err)

	post2, err := testQueries.GetPostByID(context.Background(), post.PostID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, post2)
}
