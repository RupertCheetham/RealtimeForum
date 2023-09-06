package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"realtimeForum/db"
)

// The AddRegistrationHandler function handles POST and GET requests for adding and retrieving
// registration entries respectively, including decoding the request body, logging the received
// registration, and interacting with the database.
func AddRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	if r.Method == "POST" {
		var registration db.RegistrationEntry
		err := json.NewDecoder(r.Body).Decode(&registration)

		fmt.Println("registration:", registration)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// log.Println("Received registration:", registration.Username, registration.Age, registration.Gender, registration.FirstName, registration.LastName, registration.Email, registration.Password)
		err = db.AddRegistrationToDatabase(registration.Username, registration.Age, registration.Gender, registration.FirstName, registration.LastName, registration.Email, registration.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}

	if r.Method == "GET" {
		registrations, err := db.GetRegistrationFromDatabase()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(registrations) > 0 {
			json.NewEncoder(w).Encode(registrations)
		} else {
			w.Write([]byte("No registrations available"))
		}
	}

}
