package handlers

import "net/http"

// Handler for homepage
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Real Time Forum API"))
}
