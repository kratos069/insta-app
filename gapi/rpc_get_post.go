package gapi

import (
	"context"
	"errors"

	"github.com/insta-app/pb"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetPost(ctx context.Context, req *pb.GetPostRequest) (
	*pb.GetPostResponse, error) {
	// Fetch post from database
	getPost, err := server.store.GetPostByID(ctx, req.PostId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "post not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get post: %v", err)
	}

	return &pb.GetPostResponse{
		Post: convertPost(getPost),
	}, nil
}
