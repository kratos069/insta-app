package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/insta-app/db/sqlc"
	"github.com/insta-app/token"
)

type follow struct {
	FollowingUserID int64 `uri:"following_user_id" binding:"required,min=1"`
}

func (server *Server) createFollow(ctx *gin.Context) {
	var input follow

	if err := ctx.ShouldBindUri(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	user, err := server.store.GetUserByID(ctx, input.FollowingUserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.CreateFollowParams{
		FollowerID:  authPayload.UserID, // userID
		FollowingID: user.UserID,        // to be followed user
	}

	// prevent self-follow
	if authPayload.UserID == user.UserID {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot follow yourself",
		})
		return
	}

	err = server.store.CreateFollow(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error": "followed successfully",
	})
}

func (server *Server) unFollow(ctx *gin.Context) {
	var input follow

	if err := ctx.ShouldBindUri(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	user, err := server.store.GetUserByID(ctx, input.FollowingUserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.DeleteFollowParams{
		FollowerID:  authPayload.UserID, // userID
		FollowingID: user.UserID,        // to be un-followed user
	}

	// prevent self-unfollow
	if authPayload.UserID == user.UserID {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot un-follow yourself",
		})
		return
	}

	err = server.store.DeleteFollow(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "unfollowed successfully",
	})
}

type listFollowFollowers struct {
	UserIDList int64 `uri:"user_id_followers" binding:"required,min=1"`
}

func (server *Server) listFollowers(ctx *gin.Context) {
	var input listFollowFollowers

	if err := ctx.ShouldBindUri(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	followers, err := server.store.ListFollowers(ctx, input.UserIDList)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, followers)
}

type listFollowFollowing struct {
	UserIDList int64 `uri:"user_id_following" binding:"required,min=1"`
}

func (server *Server) listFollowing(ctx *gin.Context) {
	var input listFollowFollowing

	if err := ctx.ShouldBindUri(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	following, err := server.store.ListFollowing(ctx, input.UserIDList)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, following)
}
