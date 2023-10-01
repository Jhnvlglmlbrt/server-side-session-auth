package api

import (
	"net/http"

	"github.com/Jhnvlglmlbrt/server-side-session-auth/usecases"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	usecases.DeleteCookie(w, c)
}
