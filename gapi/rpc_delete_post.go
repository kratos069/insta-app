package gapi

import (
	"context"
	"log"
	"path/filepath"
	"strings"

	"github.com/insta-app/pb"
	"github.com/insta-app/util"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) DeletePost(ctx context.Context,
	req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	// authorization
	authPayload, err := server.authorizeUser(ctx, []string{util.AdminRole, util.CustomerRole})
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	post, err := server.store.GetPostByID(ctx, req.PostId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "post not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get post: %v", err)
	}

	if authPayload.Role != util.AdminRole && post.UserID != authPayload.UserID {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized for deleting: %v", err)
	}

	// Initialize Cloudinary service.
	cloudService, err := util.NewCloudinaryService()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to initialize cloud service: %v", err)
	}

	// If the post has an image URL, attempt to delete the image from Cloudinary.
	if post.ContentUrl != "" {
		publicID := extractPublicID(post.ContentUrl)
		if err := cloudService.DeleteImage(ctx, publicID); err != nil {
			// Log the error but return an internal error.
			log.Printf("Failed to delete image from Cloudinary: %v", err)
			return nil, status.Errorf(codes.Internal, "failed to delete image from cloud: %v", err)
		}
	}

	// Delete the post from the database.
	if err := server.store.DeletePost(ctx, post.PostID); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete post from db: %v", err)
	}

	// Return a success message.
	return &pb.DeletePostResponse{
		Message: "post successfully deleted",
	}, nil
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
