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
		var comment db.CommentEntry
		err := json.NewDecoder(r.Body).Decode(&comment)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Println("Received comment:", comment.Body)

		err = db.AddCommentToDatabase(comment.Username, comment.Id, comment.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}

	// if r.Method == "GET" {
	// 	comments, err := db.GetCommentFromDatabase()
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}

	// 	if len(comments) > 0 {
	// 		json.NewEncoder(w).Encode(comments)
	// 	} else {
	// 		w.Write([]byte("No posts available"))
	// 	}
	// }

}
