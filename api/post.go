package api

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	db "github.com/insta-app/db/sqlc"
	"github.com/insta-app/token"
	"github.com/insta-app/util"
)

func (server *Server) createPost(ctx *gin.Context) {
	caption := ctx.PostForm("caption")

	if caption == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "caption missing...",
		})
		return
	}

	imageUrl, err := uploadToCloud(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.CreatePostParams{
		UserID:     authPayload.UserID,
		ContentUrl: imageUrl,
		Caption:    caption,
	}

	post, err := server.store.CreatePost(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, post)
}

type inputPostID struct {
	PostID int64 `uri:"post_id" binding:"required,min=1"`
}

func (server *Server) getPost(ctx *gin.Context) {
	var postIdReq inputPostID

	if err := ctx.ShouldBindUri(&postIdReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	getPost, err := server.store.GetPostByID(ctx, postIdReq.PostID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, getPost)
}

func (server *Server) updatePost(ctx *gin.Context) {

	var postIdReq inputPostID

	if err := ctx.ShouldBindUri(&postIdReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	post, err := server.store.GetPostByID(ctx, postIdReq.PostID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	caption := ctx.PostForm("caption")

	arg := db.UpdatePostParams{
		PostID:  post.PostID,
		Caption: caption,
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if post.UserID != authPayload.UserID {
		err := errors.New("post doesnot belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errResponse(err))
		return
	}

	updatedPost, err := server.store.UpdatePost(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, updatedPost)
}

func (server *Server) deletePost(ctx *gin.Context) {
	var postIdReq inputPostID

	if err := ctx.ShouldBindUri(&postIdReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	post, err := server.store.GetPostByID(ctx, postIdReq.PostID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if post.UserID != authPayload.UserID {
		err := errors.New("post doesnot belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errResponse(err))
		return
	}

	cloudService, err := util.NewCloudinaryService()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	// Delete image from Cloudinary
	// extract cloud public ID from URL
	if post.ContentUrl != "" {
		publicID := extractPublicID(post.ContentUrl)
		// fmt.Printf("Extracted Public ID: %s\n", publicID)

		if err := cloudService.DeleteImage(ctx, publicID); err != nil {
			log.Printf("Failed to delete image from Cloudinary: %v\n", err)
			ctx.JSON(http.StatusInternalServerError, errResponse(err))
			return
		}
	}

	err = server.store.DeletePost(ctx, post.PostID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "post successfully deleted",
	})
}

// ------------------------------------------------------------------//
// ------------------------------Helper Funcs------------------------//
// ------------------------------------------------------------------//

func uploadToCloud(ctx *gin.Context) (string, error) {
	file, err := ctx.FormFile("content_url")
	if err != nil {
		// checking file size (less than 5 mb)
		if file.Size > 5*1024*1024 {
			ctx.JSON(http.StatusBadRequest, errResponse(err))
			return "", err
		}

		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return "", err
	}

	// check format validity
	if !containsValidFormat(file.Header.Get("Content-Type")) {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return "", err
	}

	// Upload the image locally
	err = ctx.SaveUploadedFile(file, "uploads/"+file.Filename)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return "", err
	}

	cloudService, err := util.NewCloudinaryService()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return "", err
	}

	imageUrl, err := cloudService.UploadImage(ctx, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return "", err
	}

	return imageUrl, nil
}

// Helper function to check if a string exists in a slice
func containsValidFormat(item string) bool {
	slice := []string{"image/png", "image/jpeg", "image/jpg", "image/gif"}

	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// Helper function to extract public ID from a Cloudinary URL
func extractPublicID(url string) string {
	parts := strings.Split(url, "/")
	lastPart := parts[len(parts)-1]
	publicID := strings.TrimSuffix(lastPart, filepath.Ext(lastPart)) // Remove file extension

	// Extract folder path if it exists
	if len(parts) > 7 { // Cloudinary path structure
		folderPath := strings.Join(parts[7:len(parts)-1], "/") // Preserve folder structure
		publicID = folderPath + "/" + publicID
	}

	return publicID
}
