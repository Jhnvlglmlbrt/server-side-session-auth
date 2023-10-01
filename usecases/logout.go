package usecases

import (
	"net/http"
	"time"
)

func DeleteCookie(w http.ResponseWriter, c *http.Cookie) {
	sessionToken := c.Value

	delete(Sessions, sessionToken)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now(),
	})
}
