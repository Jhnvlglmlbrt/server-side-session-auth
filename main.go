package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Jhnvlglmlbrt/server-side-session-auth/api"
	_ "github.com/lib/pq"
)

func main() {
	initDB()

	http.HandleFunc("/signup", api.Signup)
	http.HandleFunc("/signin", api.Signin)
	http.HandleFunc("/welcome", api.Welcome)
	http.HandleFunc("/refresh", api.Refresh)
	http.HandleFunc("/logout", api.Logout)

	log.Fatal(http.ListenAndServe(":4000", nil))

}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Rachel"
	dbname   = "mydb"
)

func initDB() {
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	api.Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = api.Db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
