package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"realtimeForum/db"
	"realtimeForum/utils"
)

// Handler for adding post to DB
func AddPostHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	cookie, _ := r.Cookie("sessionID")

	fmt.Println("get cookie in addposthandler:", cookie)

	// This code block is handling the logic for adding a new post to the database.
	if r.Method == "POST" {
		var post db.PostEntry
		err := json.NewDecoder(r.Body).Decode(&post)

		if err != nil {
			utils.HandleError("Problem decoding JSON in AddPostHandler", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Println("Received post:", post.UserId, post.Img, post.Body, post.Categories)

		err = db.AddPostToDatabase(post.UserId, post.Img, post.Body, post.Categories)
		if err != nil {
			utils.HandleError("Problem adding to POSTS in AddPostHandler", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		msg := map[string]string{
			"message": "post made successfully",
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
}

// Handler for getting post from DB
func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	// This code block is handling the logic for retrieving posts from the database when the HTTP request
	// method is GET.
	if r.Method == "GET" {
		posts, err := db.GetPostFromDatabase()
		if err != nil {
			utils.HandleError("Problem getting posts from db in AddPostHandler", err)
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
