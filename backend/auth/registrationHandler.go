package auth

import (
	"encoding/json"
	"net/http"
	"realtimeForum/db"
	"realtimeForum/handlers"
	"realtimeForum/utils"

	"golang.org/x/crypto/bcrypt"
)

// The AddUserHandler function handles POST and GET requests for adding and retrieving
// user entries respectively, including decoding the request body, logging the received
// user, and interacting with the database.
func AddUserHandler(w http.ResponseWriter, r *http.Request) {
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
