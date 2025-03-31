package gapi

import (
	"fmt"

	db "github.com/insta-app/db/sqlc"
	"github.com/insta-app/pb"
	"github.com/insta-app/token"
	"github.com/insta-app/util"
	"github.com/insta-app/worker"
)

// servers gRPC requests for the insta-app
type Server struct {
	pb.UnimplementedInstaAppServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// Creates gRPC server
func NewServer(config util.Config, store db.Store,
	taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
