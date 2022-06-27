package token

import "time"

//Maker is an interface for managing tokens

type Maker interface {
	//CreateToken creates a new token
	CreateToken(username string, duration time.Duration) (string, error)

	//VerifyToken verifies a token
	VerifyToken(token string) (*Payload, error)
}
