package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"realtimeForum/db"
	"realtimeForum/handlers"
	"realtimeForum/utils"
	"time"
)

const timeout = 2 * time.Minute

var sessionExpiration = time.Now().Add(timeout)

// Handler to login page
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers for this handler
	handlers.SetupCORS(&w, r)
	var login db.UserEntry
	if r.Method == "POST" {
		err := json.NewDecoder(r.Body).Decode(&login)

		fmt.Println("login data:", login)

		if err != nil {
			utils.HandleError("Unable to decode json", err)
			http.Error(w, "Internal sever error", http.StatusBadRequest)
			return
		}

		// use the login data to find the user in the database
		msg, id, err := db.GetLoginEntry(login)

		if err != nil {
			utils.HandleError("Unable to get user's id", err)
			http.Error(w, "Internal sever error", http.StatusBadRequest)
			return
		}

		jsonResponse, err := json.Marshal(msg)

		fmt.Println("jsonResponse:", msg)

		if err != nil {
			utils.HandleError("Unable to decode json", err)
			http.Error(w, "Internal sever error", http.StatusBadRequest)
			return
		}

		// after getting username from login check, use the user's id to create a session for the user
		userSession, err := db.CreateSession(id, sessionExpiration)

		if err != nil {
			utils.HandleError("unable to create user session:", err)
			http.Error(w, "Internal sever error", http.StatusBadRequest)
			return
		}

		// Set the session ID as a cookie
		sessionCookie := http.Cookie{
			Name:     "sessionID",
			Value:    userSession.SessionID,
			Expires:  sessionExpiration,
			HttpOnly: true,
			Secure:   true,
			Path:     "/",
			Domain:   "localhost",
			MaxAge:   int(timeout.Seconds()),
			SameSite: http.SameSiteNoneMode,
		}
		http.SetCookie(w, &sessionCookie)

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}

	if r.Method == "GET" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized access"))
	}
}
