package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"realtimeForum/db"
)

// Handler for comments
func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	// This code block is handling the POST request for adding a comment.
	if r.Method == "POST" {
		var comment db.CommentEntry
		err := json.NewDecoder(r.Body).Decode(&comment)
		log.Println(comment)
		log.Println(comment.ParentPostID)
		if err != nil {
			log.Println("Error in AddCommentHandler")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Println("Received comment:", comment.Body)

		err = db.AddCommentToDatabase(comment.Username, comment.ParentPostID, comment.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}

	// This code block is handling the GET request for retrieving comments from the database.
	if r.Method == "GET" {
		comments, err := db.GetCommentsFromDatabase()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(comments) > 0 {
			json.NewEncoder(w).Encode(comments)
		} else {
			w.Write([]byte("No posts available"))
		}
	}

}
