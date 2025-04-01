package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateFollow(t *testing.T) {
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)

	arg := CreateFollowParams{
		FollowerID:  user1.UserID,
		FollowingID: user2.UserID,
	}

	err := testStore.CreateFollow(context.Background(), arg)
	require.NoError(t, err)
}

func TestDeleteFollow(t *testing.T) {
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)

	arg := DeleteFollowParams{
		FollowerID:  user1.UserID,
		FollowingID: user2.UserID,
	}

	err := testStore.DeleteFollow(context.Background(), arg)
	require.NoError(t, err)
}

func TestListFollowers(t *testing.T) {
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)

	arg := CreateFollowParams{
		FollowerID:  user1.UserID,
		FollowingID: user2.UserID,
	}

	err := testStore.CreateFollow(context.Background(), arg)
	require.NoError(t, err)

	followers, err := testStore.ListFollowers(context.Background(), user2.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, followers)
}

func TestListFollowing(t *testing.T) {
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)

	arg := CreateFollowParams{
		FollowerID:  user1.UserID,
		FollowingID: user2.UserID,
	}

	err := testStore.CreateFollow(context.Background(), arg)
	require.NoError(t, err)

	following, err := testStore.ListFollowing(context.Background(), user1.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, following)
}