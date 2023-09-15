package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"realtimeForum/db"
	"realtimeForum/utils"
)

// The AddUserHandler function handles POST and GET requests for adding and retrieving
// user entries respectively, including decoding the request body, logging the received
// user, and interacting with the database.
func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	// The code block is handling the POST request for adding a user entry to the database.
	if r.Method == "POST" {
		var user db.UserEntry
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			utils.HandleError("Problem decoding JSON in AddUserHandler", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Println("Received user:", user.Username, user.Age, user.Gender, user.FirstName, user.LastName, user.Email, user.Password)

		err = db.AddUserToDatabase(user.Username, user.Age, user.Gender, user.FirstName, user.LastName, user.Email, user.Password)
		if err != nil {
			utils.HandleError("Problem adding to USERS in AddUserHandler", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}

	// The code block is handling the GET request for retrieving user entries from the database.
	if r.Method == "GET" {
		users, err := db.GetUsersFromDatabase()
		if err != nil {
			utils.HandleError("Problem getting USERS from db in AddUserHandler", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(users) > 0 {
			json.NewEncoder(w).Encode(users)
		} else {
			w.Write([]byte("No users available"))
		}
	}

}
