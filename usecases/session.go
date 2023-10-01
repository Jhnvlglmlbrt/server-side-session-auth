package usecases

func CheckSession(sessionToken string) (*Session, bool) {
	userSession, exists := Sessions[sessionToken]
	if !exists {
		return &Session{}, false
	}

	if userSession.isExpired() {
		delete(Sessions, sessionToken)
		return &Session{}, false
	}

	return &userSession, true
}
