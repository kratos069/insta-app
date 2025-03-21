package api

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/insta-app/db/sqlc"
	"github.com/insta-app/token"
	"github.com/insta-app/util"
)

type createUserRequest struct {
	Username       string `json:"username" binding:"required,alphanum"`
	Password       string `json:"password" binding:"required,min=6"`
	FullName       string `json:"full_name" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	ProfilePicture string `json:"profile_picture" binding:"required"`
	Bio            string `json:"bio" binding:"required"`
}

type userResponse struct {
	UserID         int64     `json:"user_id"`
	Username       string    `json:"username"`
	FullName       string    `json:"full_name"`
	Email          string    `json:"email"`
	ProfilePicture string    `json:"profile_picture"`
	Bio            string    `json:"bio"`
	CreatedAt      time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		UserID:         user.UserID,
		Username:       user.Username,
		FullName:       user.FullName,
		Email:          user.Email,
		ProfilePicture: user.ProfilePicture,
		Bio:            user.Bio,
		CreatedAt:      user.CreatedAt,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
		ProfilePicture: req.ProfilePicture,
		Bio:            req.Bio,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	resp := newUserResponse(user)

	ctx.JSON(http.StatusOK, resp)
}

type inputUserID struct {
	UserID int64 `uri:"user_id" binding:"required,min=1"`
}

func (server *Server) getUserByID(ctx *gin.Context) {
	var input inputUserID

	if err := ctx.ShouldBindUri(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	user, err := server.store.GetUserByID(ctx, input.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	resp := newUserResponse(user)

	ctx.JSON(http.StatusOK, resp)
}

func (server *Server) deleteUser(ctx *gin.Context) {
	var input inputUserID

	if err := ctx.ShouldBindUri(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	user, err := server.store.GetUserByID(ctx, input.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if user.UserID != authPayload.UserID {
		err := errors.New("account doesnot belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errResponse(err))
		return
	}

	err = server.store.DeleteUser(ctx, user.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
	})
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	SessionID             uuid.UUID    `json:"session_id"`
	AccessToken           string       `json:"access_token"`
	AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
	RefreshToken          string       `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
	User                  userResponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var input loginUserRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	user, err := server.store.GetUserByUsername(ctx, input.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	err = util.CheckPassword(input.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errResponse(err))
		return
	}

	// creating access token
	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		user.UserID,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	// creating refresh token
	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		user.UserID,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	// create session
	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     accessPayload.Username,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiredAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	resp := loginUserResponse{
		SessionID:             session.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  newUserResponse(user),
	}

	ctx.JSON(http.StatusOK, resp)
}
