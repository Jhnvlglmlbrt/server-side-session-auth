package api

import (
	"fmt"
	"net/http"

	"github.com/Jhnvlglmlbrt/server-side-session-auth/usecases"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sessionToken := c.Value

	userSession, valid := usecases.CheckSession(sessionToken)
	if !valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Welcome %s!", userSession.Username)))
}
