package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/insta-app/db/sqlc"
	"github.com/insta-app/token"
	"github.com/insta-app/util"
)

// servers HTTP requests for the insta-app
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// Creates HTTP server and Setup Routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	// Routes
	server.setupRoutes()

	return server, nil
}

func (server *Server) setupRoutes() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/tokens/renew_access", server.renewAccessToken)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker,
		[]string{util.AdminRole, util.CustomerRole}))

	authRoutes.GET("/users/:user_id", server.getUserByID)
	authRoutes.DELETE("/users/:user_id", server.deleteUser)

	authRoutes.POST("/posts", server.createPost)
	authRoutes.GET("/posts/:post_id", server.getPost)
	authRoutes.PUT("/posts/:post_id", server.updatePost)
	authRoutes.DELETE("/posts/:post_id", server.deletePost)

	authRoutes.POST("/comments/:post_id", server.createComment)
	authRoutes.DELETE("/comments/:comment_id", server.deleteComment)
	authRoutes.GET("/comments/:post_id", server.listCommentsByPost)

	authRoutes.POST("/follows/:following_user_id", server.createFollow)
	authRoutes.DELETE("/follows/:following_user_id", server.unFollow)
	authRoutes.GET("/follows/follower/:user_id_followers", server.listFollowers)
	authRoutes.GET("/follows/following/:user_id_following", server.listFollowing)

	authRoutes.POST("/likes/like/:post_id", server.createLike)
	authRoutes.POST("/likes/unlike/:post_id", server.unLike)
	authRoutes.GET("/likes/:post_id", server.countLikesByPost)

	server.router = router

}

// Starts and runs HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
