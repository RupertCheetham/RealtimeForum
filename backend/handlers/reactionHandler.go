package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"realtimeForum/db"
	"realtimeForum/utils"
)

// Handler for posts page
func ReactionHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	// This code block is handling the logic for adding a new post to the database.
	if r.Method == "POST" {
		var reactionEntry db.ReactionEntry
		err := json.NewDecoder(r.Body).Decode(&reactionEntry)

		if err != nil {
			log.Println("Problem decoding JSON in ReactionHandler", err)
			utils.HandleError("Problem decoding JSON in ReactionHandler", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Println("Received", reactionEntry.Action, "for", reactionEntry.Type, reactionEntry.ParentID, "from UserID:", reactionEntry.UserID, ", reactionID:", reactionEntry.ReactionID)

		if reactionEntry.ReactionID == 0 {
			db.AddReactionToDatabase(reactionEntry.Type, reactionEntry.ParentID, reactionEntry.UserID, reactionEntry.Action)
		} else {
			db.UpdateReactionInDatabase(reactionEntry.Type, reactionEntry.ReactionID, reactionEntry.UserID, reactionEntry.Action)
		}

		w.WriteHeader(http.StatusCreated)

	}

	// This code block is handling the logic for retrieving posts from the database when the HTTP request
	// method is GET.
	// if r.Method == "GET" {
	// 	posts, err := db.GetPostFromDatabase()
	// 	if err != nil {
	// 		utils.HandleError("Problem getting posts from db in AddPostHandler", err)
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}

	// 	if len(posts) > 0 {
	// 		json.NewEncoder(w).Encode(posts)
	// 	} else {
	// 		w.Write([]byte("No posts available"))
	// 	}
	// }

}
