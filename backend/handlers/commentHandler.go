package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"realtimeForum/db"
)

// Handler for posts page
func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
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
