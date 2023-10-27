package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const CookieName = "sessionID"

// The function sets up Cross-Origin Resource Sharing (CORS) headers for an HTTP response.
func SetupCORS(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "https://localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Methods", "HEAD, POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Requested-With")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")

	// Respond to preflight request with 200 OK status
	if r.Method == "OPTIONS" {
		(*w).WriteHeader(http.StatusOK)
		return
	}
}

func GetCookie(w http.ResponseWriter, r *http.Request) string {
	SetupCORS(&w, r)
	sessionCookie, _ := r.Cookie(CookieName)
	fmt.Println("sessioncookie in getCookie", sessionCookie)
	splitCookieValue := strings.Split(sessionCookie.String(), "=")
	var cookieValue string

	if len(splitCookieValue) == 2 {
		cookieValue = splitCookieValue[1]
	}

	return cookieValue
}

func ActionSuccessMessage(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)
	msg := map[string]string{
		"message": "successful request",
	}

	jsonResponse, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

func ActionFailedMessage(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)
	msg := map[string]string{
		"message": "unsuccessful request",
	}

	jsonResponse, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusRequestTimeout)
	w.Write(jsonResponse)
}
