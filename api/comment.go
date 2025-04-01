package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/insta-app/db/sqlc"
	"github.com/insta-app/token"
	"github.com/jackc/pgx/v5"
)

type comment struct {
	CommentContent string `json:"comment_content" binding:"required"`
}

type post struct {
	CommentPostID int64 `uri:"post_id" binding:"required,min=1"`
}

// will create a comment on a post
func (server *Server) createComment(ctx *gin.Context) {
	var inputComment comment
	var inputPostID post

	if err := ctx.ShouldBindJSON(&inputComment); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	if err := ctx.ShouldBindUri(&inputPostID); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	getPost, err := server.store.GetPostByID(ctx, inputPostID.CommentPostID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	arg := db.CreateCommentParams{
		UserID:  getPost.UserID,
		PostID:  getPost.PostID,
		Content: inputComment.CommentContent,
	}

	comment, err := server.store.CreateComment(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

type commentID struct {
	CommentID int64 `uri:"comment_id" binding:"required,min=1"`
}

// will delete a comment from a post
func (server *Server) deleteComment(ctx *gin.Context) {
	var input commentID

	if err := ctx.ShouldBindUri(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	comment, err := server.store.GetCommentByID(ctx, input.CommentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	fmt.Printf("Auth UserID: %v, Comment UserID: %v\n", authPayload.UserID, comment.UserID)

	if authPayload.UserID != comment.UserID {
		err := errors.New("comment doesnot belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errResponse(err))
		return
	}

	err = server.store.DeleteComment(ctx, comment.CommentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "comment successfully deleted",
	})
}

// list all comments on a specific post
func (server *Server) listCommentsByPost(ctx *gin.Context) {
	var input post

	if err := ctx.ShouldBindUri(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	post, err := server.store.GetPostByID(ctx, input.CommentPostID)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	comments, err := server.store.ListCommentsByPost(ctx, post.PostID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, comments)
}
