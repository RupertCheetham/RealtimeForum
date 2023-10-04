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
	if r.Method == "GET" {

		// parentID, err := strconv.Atoi(r.URL.Query().Get("parentID"))
		// if err != nil {
		// 	log.Println("There was an issue with parentID in ReactionHandler")
		// 	utils.HandleError("There was an isusue with parentID in ReactionHandler", err)
		// }
		// Reads the rowID from the get request URL (http://localhost:8080/reaction?rowID=${parseInt(ReactionID)}&reactionTable=${Type})
		rowID, err := strconv.Atoi(r.URL.Query().Get("rowID"))
		if err != nil {
			log.Println("There was an issue converting a string to an int in ReactionHandler")
			utils.HandleError("There was an issue converting a string to an int in ReactionHandler", err)
		}
		log.Println("rowID is:", rowID)
		// Reads the reactionTable from the request URL
		reactionTable := r.URL.Query().Get("reactionTable")
		log.Println("type is:", reactionTable)
		if rowID == 0 {
			// postOrCommentTable := ""
			// if reactionTable == "POSTREACTIONS" {
			// 	postOrCommentTable = "POSTS"
			// } else {
			// 	postOrCommentTable = "COMMENTS"
			// }

			// updateQuery := fmt.Sprintf("UPDATE %s SET ReactionID = (SELECT Id FROM %s ORDER BY Id DESC LIMIT 1) WHERE Id = ?", postOrCommentTable, reactionTable)
			// _, err = db.Database.Exec(updateQuery, parentID)
			// if err != nil {
			// 	log.Println("There was a problem with updating a post/comments reactionID in ReactionHandler", err)
			// 	utils.HandleError("There was a problem with updating a post/comments reactionID in ReactionHandler", err)
			// }
			// Query for the latest ReactionID after the update
			var latestReactionID int
			query := fmt.Sprintf("SELECT Id FROM %s ORDER BY Id DESC LIMIT 1", reactionTable)
			err = db.Database.QueryRow(query).Scan(&latestReactionID)

			if err != nil {
				// Handle the error
				log.Println("Error querying latest ReactionID:", err)
				// You may return an error response or handle it as needed
			}
			rowID = latestReactionID
		}

		likes, dislikes, err := GetLikesAndDislikes(reactionTable, rowID)
		log.Println("likes are", likes, "and dislikes are", dislikes)
		if err != nil {
			log.Println("There was a problem with GetLikesAndDislikes in ReactionHandler")
			utils.HandleError("There was a problem with GetLikesAndDislikes in ReactionHandler", err)
		}

		response := struct {
			Likes    int
			Dislikes int
		}{
			Likes:    likes,
			Dislikes: dislikes,
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Println("There was a problem marshalling in ReactionHandler.", err)
			utils.HandleError("There was a problem marshalling in ReactionHandler.", err)

		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonResponse)
	}

}

// Returns the likes and dislikes for a given post/comment, from the relevant table
func GetLikesAndDislikes(reactionTable string, rowID int) (int, int, error) {
	query := fmt.Sprintf("SELECT Likes, Dislikes FROM %s WHERE Id = ?", reactionTable)

	var likes, dislikes int
	err := db.Database.QueryRow(query, rowID).Scan(&likes, &dislikes)
	if err != nil {
		utils.HandleError("there was a problem in GetLikesAndDislikes", err)
		return 0, 0, err
	}

	return likes, dislikes, nil
}
