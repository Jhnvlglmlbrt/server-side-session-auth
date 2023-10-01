package usecases

import (
	"time"

	"github.com/google/uuid"
)

var SessionToken = uuid.NewString()

func SetToken(username string) {
	expiresAt := time.Now().Add(120 * time.Second)

	Sessions[SessionToken] = Session{
		Username: username,
		Expiry:   expiresAt,
	}
}
