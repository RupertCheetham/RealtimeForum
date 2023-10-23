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

const timeout = 30 * time.Minute

var sessionExpiration = time.Now().Add(timeout)

// Handler to login page
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers for this handler
	handlers.SetupCORS(&w, r)
	var login db.UserEntry
	if r.Method == "POST" {
		err := json.NewDecoder(r.Body).Decode(&login)

		if err != nil {
			utils.HandleError("Unable to decode json", err)
			http.Error(w, "Unable to decode json", http.StatusBadRequest)
			return
		}

		msg, err := db.GetLoginEntry(login)
		if err != nil {
			utils.HandleError("Unable to sign in user", err)
			http.Error(w, "Unable to sign in user", http.StatusBadRequest)
			return
		}

		jsonResponse, err := json.Marshal(msg)

		fmt.Println("jsonResponse:", msg)

		if err != nil {
			utils.HandleError("Unable to decode json", err)
			http.Error(w, "Unable to decode json", http.StatusBadRequest)
			return
		}

		dbLoginCheck, err := db.FindUserFromDatabase(login.Username)

		if err != nil {
			fmt.Println("unable to create user session:", err)
		}

		fmt.Println("sessionExpiration:", sessionExpiration)

		userSession, err := db.CreateSession(dbLoginCheck[0].Id, sessionExpiration)

		if err != nil {
			fmt.Println("unable to create user session:", err)
			utils.HandleError("unable to create user session:", err)
			http.Error(w, "Unable to get user id", http.StatusBadRequest)
			return
		}

		// Set the session ID as a cookie
		sessionCookie := http.Cookie{
			Name:    "sessionID",
			Value:   userSession.SessionID,
			Expires: sessionExpiration,
			// HttpOnly: true,
			Secure: true,
			Path:   "/",
			Domain: "localhost",
			MaxAge: int(timeout.Seconds()),
			// SameSite: http.SameSiteNoneMode,
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
