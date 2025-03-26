package gapi

import (
	"fmt"

	db "github.com/insta-app/db/sqlc"
	"github.com/insta-app/pb"
	"github.com/insta-app/token"
	"github.com/insta-app/util"
)

// servers gRPC requests for the insta-app
type Server struct {
	pb.UnimplementedInstaAppServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// Creates gRPC server
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

	return server, nil
}
