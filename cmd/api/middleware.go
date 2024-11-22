package main

import "net/http"

// SessionLoad loads and saved the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
