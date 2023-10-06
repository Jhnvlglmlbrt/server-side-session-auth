package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Jhnvlglmlbrt/server-side-session-auth/usecases"
	"golang.org/x/crypto/bcrypt"
)

func Signin(w http.ResponseWriter, r *http.Request) {

	creds := &usecases.Creds{}

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result := usecases.Db.QueryRow("select password from users where usename=$1", creds.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	storeCreds := &usecases.Creds{}

	err = result.Scan(&storeCreds.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(storeCreds.Password), []byte(creds.Password)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}

	// if !usecases.VerifyPassword(creds.Username, creds.Password) {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	usecases.SetToken(creds.Username)

	usecases.SetTokenCookie(w, usecases.SessionToken)
}
