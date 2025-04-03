package gapi

import (
	db "github.com/insta-app/db/sqlc"
	"github.com/insta-app/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// convert db User type to pb User type
func convertUser(user db.User) *pb.User {
	return &pb.User{
		UserId:         user.UserID,
		Username:       user.Username,
		FullName:       user.FullName,
		Email:          user.Email,
		ProfilePicture: user.ProfilePicture,
		Bio:            user.Bio,
		CreatedAt:      timestamppb.New(user.CreatedAt),
	}
}

