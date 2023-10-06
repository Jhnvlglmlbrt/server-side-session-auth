package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var Db *sql.DB

type Creds struct {
	Password string `json:"password" db:"password"`
	Username string `json:"username" db:"username"`
}

func Signup(w http.ResponseWriter, r *http.Request) {

	creds := &Creds{}

	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)

	if _, err = Db.Query("insert into users values ($1, $2)", creds.Username, string(hashedPassword)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
