package main

import (
	"log"
	"net/http"

	"github.com/Jhnvlglmlbrt/server-side-session-auth/api"
)

func main() {
	http.HandleFunc("/signin", api.Signin)
	http.HandleFunc("/welcome", api.Welcome)
	http.HandleFunc("/refresh", api.Refresh)
	http.HandleFunc("/logout", api.Logout)

	log.Fatal(http.ListenAndServe(":4000", nil))
}
