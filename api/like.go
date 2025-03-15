package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/insta-app/db/sqlc"
)

type like struct {
	PostID int64 `uri:"post_id" binding:"required,min=1"`
}

func (server *Server) createLike(ctx *gin.Context) {
	var input like

	if err := ctx.ShouldBindUri(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	post, err := server.store.GetPostByID(ctx, input.PostID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	arg := db.CreateLikeParams{
		UserID: post.UserID,
		PostID: post.PostID,
	}

	like, err := server.store.CreateLike(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, like)
}

func (server *Server) unLike(ctx *gin.Context) {
	var input like

	if err := ctx.ShouldBindUri(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	post, err := server.store.GetPostByID(ctx, input.PostID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	arg := db.DeleteLikeParams{
		UserID: post.UserID,
		PostID: post.PostID,
	}

	err = server.store.DeleteLike(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "unliked successfully",
	})
}

func (server *Server) countLikesByPost(ctx *gin.Context) {
	var input like

	if err := ctx.ShouldBindUri(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	post, err := server.store.GetPostByID(ctx, input.PostID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	likes, err := server.store.CountLikesByPost(ctx, post.PostID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, likes)
}
