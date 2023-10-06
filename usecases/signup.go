package usecases

import "database/sql"

var Db *sql.DB

type Creds struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}
