package auth

import (
	"encoding/json"
	"net/http"
	"realtimeForum/db"
	"realtimeForum/handlers"
	"realtimeForum/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const timeout = 1 * time.Minute

var sessionExpiration = time.Now().Add(timeout)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers for this handler
	handlers.SetupCORS(&w, r)
	var login db.UserEntry
	if r.Method == "POST" {
		err := json.NewDecoder(r.Body).Decode(&login)

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
			Name:     handlers.CookieName,
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

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	handlers.SetupCORS(&w, r)
	cookie := http.Cookie{
		Name:     handlers.CookieName,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   true,
		Domain:   "localhost",
	}

	http.SetCookie(w, &cookie)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// The AddUserHandler function handles POST and GET requests for adding and retrieving
// user entries respectively, including decoding the request body, logging the received
// user, and interacting with the database.
func RegistrationUserHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers for this handler
	handlers.SetupCORS(&w, r)

	// The code block is handling the POST request for adding a user entry to the database.
	if r.Method == "POST" {
		var user db.UserEntry
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			utils.HandleError("Problem decoding JSON in AddUserHandler", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			utils.HandleError("error in encryption", err)
		}

		// log.Println("Received user:", user.Username, user.Age, user.Gender, user.FirstName, user.LastName, user.Email, string(hashPassword))
		err = db.AddUserToDatabase(user.Username, user.Age, user.Gender, user.FirstName, user.LastName, user.Email, string(hashPassword))
		if err != nil {
			utils.HandleError("Unable to register a new user in AddUserHandler", err)
			// fmt.Println("Unable to register a new user in AddUserHandler", err)
			http.Error(w, "Unable to register a new user", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}
