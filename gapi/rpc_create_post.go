package gapi

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	db "github.com/insta-app/db/sqlc"
	"github.com/insta-app/pb"
	"github.com/insta-app/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreatePost(stream grpc.ClientStreamingServer[pb.CreatePostRequest,
	pb.CreatePostResponse]) error {
	// authorization
	authPayload, err := server.authorizeUser(stream.Context())
	if err != nil {
		return unauthenticatedError(err)
	}

	// Create a temporary file to store the uploaded file chunks.
	tmpFile, err := os.CreateTemp("", "upload-*")
	if err != nil {
		return status.Errorf(codes.Internal, "failed to create temp file: %v", err)
	}
	defer func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}()

	// Receive the first message, which must contain the metadata.
	req, err := stream.Recv()
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "failed to receive metadata: %v", err)
	}

	meta := req.GetMetadata()
	if meta == nil {
		return status.Error(codes.InvalidArgument, "metadata missing")
	}

	userID := meta.GetUserId()
	caption := meta.GetCaption()
	fileName := meta.GetFileName()
	contentType := meta.GetContentType()

	if authPayload.UserID != userID {
		return status.Error(codes.Unauthenticated, "userID issue")
	}

	if caption == "" {
		return status.Error(codes.InvalidArgument, "caption missing")
	}
	if fileName == "" || contentType == "" {
		return status.Error(codes.InvalidArgument, "file name or content type missing")
	}

	// Read and write file chunks.
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// End of stream.
			break
		}
		if err != nil {
			return status.Errorf(codes.Internal, "error receiving file chunk: %v", err)
		}

		chunk := req.GetChunk()
		if chunk == nil {
			continue // Skip if no chunk data.
		}

		if _, err := tmpFile.Write(chunk.GetContent()); err != nil {
			return status.Errorf(codes.Internal, "failed to write file chunk: %v", err)
		}
	}

	// Validate file size (limit to 5MB).
	stat, err := tmpFile.Stat()
	if err != nil {
		return status.Errorf(codes.Internal, "failed to stat temp file: %v", err)
	}
	if stat.Size() > 5*1024*1024 {
		return status.Error(codes.InvalidArgument, "file size exceeds 5MB limit")
	}

	// Validate content type.
	if !containsValidFormat(contentType) {
		return status.Error(codes.InvalidArgument, "invalid file format")
	}

	// Ensure the uploads directory exists.
	uploadsDir := "uploads"
	if err := os.MkdirAll(uploadsDir, 0755); err != nil {
		return status.Errorf(codes.Internal, "failed to create uploads directory: %v", err)
	}

	// Save the temporary file to a permanent location.
	destPath := filepath.Join(uploadsDir, fileName)
	if _, err := tmpFile.Seek(0, io.SeekStart); err != nil {
		return status.Errorf(codes.Internal, "failed to seek temp file: %v", err)
	}
	destFile, err := os.Create(destPath)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to create destination file: %v", err)
	}
	defer destFile.Close()
	if _, err = io.Copy(destFile, tmpFile); err != nil {
		return status.Errorf(codes.Internal, "failed to save file locally: %v", err)
	}

	fileStat, err := destFile.Stat()
	if err != nil {
		return status.Errorf(codes.Internal, "failed to get file stat: %v", err)
	}
	header := &multipart.FileHeader{
		Filename: fileStat.Name(),
		Size:     fileStat.Size(),
	}

	// Upload the image to the cloud.
	cloudService, err := util.NewCloudinaryService()
	if err != nil {
		return status.Errorf(codes.Internal, "failed to initialize cloud service: %v", err)
	}
	imageUrl, err := cloudService.UploadImage(stream.Context(), header)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to upload image: %v", err)
	}

	// Create post parameters.
	arg := db.CreatePostParams{
		UserID:     authPayload.UserID,
		ContentUrl: imageUrl,
		Caption:    caption,
	}

	post, err := server.store.CreatePost(stream.Context(), arg)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to create post: %v", err)
	}

	// Build and send the response.
	res := &pb.CreatePostResponse{
		Post: convertPost(post),
	}
	return stream.SendAndClose(res)
}

// containsValidFormat checks if the provided content type is allowed.
// You can adjust this implementation based on your requirements.
func containsValidFormat(item string) bool {
	slice := []string{"image/png", "image/jpeg", "image/jpg", "image/gif"}

	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
