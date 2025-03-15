package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	// creates a token for specific username and valid duration
	CreateToken(username string, userID int64, duration time.Duration) (string, error)

	// check if input token is valid or not
	VerifyToken(token string) (*Payload, error)
}