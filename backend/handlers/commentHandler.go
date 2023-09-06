package handlers

import (
	"encoding/json"
	"net/http"
	"realtimeForum/db"
	"realtimeForum/utils"
)

// Handler for comments page
func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	if r.Method == "POST" {
		var comment db.CommentEntry
		err := json.NewDecoder(r.Body).Decode(&comment)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			utils.HandleError("Error whilst dealing with comment post in AddCommentHandler. ", err)
			return
		}

		utils.WriteMessageToLogFile("Received comment: Username - " + comment.Username + "; Comment - " + comment.Body)

		err = db.AddCommentToDatabase(comment.Username, comment.PostID, comment.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			utils.HandleError("Error with AddCommentToDatabase in AddCommentHandler. ", err)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}

	if r.Method == "GET" {
		comments, err := db.GetCommentsFromDatabase()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			utils.HandleError("Error whilst dealing with submitted comment get in AddCommentHandler. ", err)
			return
		}

		if len(comments) > 0 {
			json.NewEncoder(w).Encode(comments)
		} else {
			w.Write([]byte("No comments available"))
		}
	}

}
