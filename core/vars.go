package core

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"log"
)

var (
	auth  *Authenticator
	ctx   context.Context
	state string
	//Store *sessions.CookieStore
)

func init() {

	var err error

	// Initialize Context
	ctx = context.Background()

	// Initialize Authenticator
	auth, err = NewAuthenticator()
	if err != nil {
		log.Println("Error initalizaing authenticator:", err.Error())
		return
	}

	// Initialize Random State
	state = generateRandomState()

	// Initialize Cookie Store
	//Store = sessions.NewCookieStore([]byte("secret"))
}

// generateRandomState Creates a random state string
func generateRandomState() string {

	// Generate random state
	b := make([]byte, 32)
	rand.Read(b)

	return base64.StdEncoding.EncodeToString(b)
}
