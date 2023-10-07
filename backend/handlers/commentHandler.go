package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"realtimeForum/db"
	"realtimeForum/utils"
	"strconv"
)

// Handler for comments
func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	// This code block is handling the POST request for adding a comment.
	if r.Method == "POST" {
		fmt.Println("Oh boy, here I am in AddCommentHandler")
		var comment db.CommentEntry
		err := json.NewDecoder(r.Body).Decode(&comment)
		log.Println("comment:", comment)
		log.Println("comment.ParentPostID:", comment.ParentPostID)
		if err != nil {
			log.Println("Error in AddCommentHandler")
			utils.HandleError("Error in AddCommentHandler", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Println("Received comment:", comment.Body)

		err = db.AddCommentToDatabase(comment.ParentPostID, comment.Id, comment.Body)
		if err != nil {
			utils.HandleError("Problem adding comment to db in AddCommentHandler", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}

	// This code block is handling the GET request for retrieving comments from the database.
	if r.Method == "GET" {
		parentPostID, err := strconv.Atoi(r.URL.Query().Get("postID"))
		if err != nil {
			log.Println("There was a problem converting parentPostID to int in AddCommentHandler.", err)
			utils.HandleError("There was a problem converting parentPostID to int in AddCommentHandler.", err)
		}

		comments, err := db.GetCommentsFromDatabase(parentPostID)
		if err != nil {
			log.Println("Problem getting comment from db in AddCommentHandler", err)
			utils.HandleError("Problem getting comment from db in AddCommentHandler", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// returns the comments, if any
		json.NewEncoder(w).Encode(comments)

	}

}
