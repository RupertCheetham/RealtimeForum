package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"realtimeForum/db"
)

func SetupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// The AddRegistrationHandler function handles POST and GET requests for adding and retrieving
// registration entries respectively, including decoding the request body, logging the received
// registration, and interacting with the database.
func AddRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	if r.Method == "POST" {
		var registration db.RegistrationEntry
		err := json.NewDecoder(r.Body).Decode(&registration)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Println("Received registration:", registration.Username, registration.Age, registration.Gender, registration.FirstName, registration.LastName, registration.Email, registration.Password)

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

// Handler for posts page
func AddPostHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	if r.Method == "POST" {
		var post db.PostEntry
		err := json.NewDecoder(r.Body).Decode(&post)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Println("Received post:", post.Username, post.Img, post.Body, post.Categories)

		err = db.AddPostToDatabase(post.Username, post.Img, post.Body, post.Categories)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}

	if r.Method == "GET" {
		posts, err := db.GetPostFromDatabase()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(posts) > 0 {
			json.NewEncoder(w).Encode(posts)
		} else {
			w.Write([]byte("No posts available"))
		}
	}

}

// Handler for homepage
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Real Time Forum API"))
}
