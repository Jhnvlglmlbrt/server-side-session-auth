package usecases

import (
	"time"
)

// var users = map[string]string{
// 	"user1": "password1",
// 	"user2": "password2",
// }

var Sessions = map[string]Session{}

var expiresAt = time.Now().Add(120 * time.Second)

type Session struct {
	Username string
	Expiry   time.Time
}

// type Credentials struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

func (s Session) isExpired() bool {
	return s.Expiry.Before(time.Now())
}

// func VerifyPassword(username, password string) bool {
// 	expectedPassword, ok := users[username]
// 	return ok && expectedPassword == password
// }
