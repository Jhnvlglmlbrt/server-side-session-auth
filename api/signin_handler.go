package api

import (
	"encoding/json"
	"net/http"

	"github.com/Jhnvlglmlbrt/server-side-session-auth/usecases"
)

func Signin(w http.ResponseWriter, r *http.Request) {
	var creds usecases.Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !usecases.VerifyPassword(creds.Username, creds.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	usecases.SetToken(creds.Username)

	usecases.SetTokenCookie(w, usecases.SessionToken)
}
