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

// Handler for reactions (likes/dislikes)
func ReactionHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	if r.Method == "POST" {
		ReactionHandlerPostMethod(w, r)
	}

	if r.Method == "GET" {
		ReactionHandlerGetMethod(w, r)
	}

}

// This deals with posting decoding reaction data and sending it to the relevant functions
func ReactionHandlerPostMethod(w http.ResponseWriter, r *http.Request) {
	var reactionEntry db.ReactionEntry
	err := json.NewDecoder(r.Body).Decode(&reactionEntry)
	if err != nil {
		log.Println("Problem decoding JSON in ReactionHandler", err)
		utils.HandleError("Problem decoding JSON in ReactionHandler", err)
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

// Returns values for a given post/comment
// This function is called by the JS after it has made a reaction post request
func ReactionHandlerGetMethod(w http.ResponseWriter, r *http.Request) {

	// Reads whether the reaction was to a post or a comment from the request URL
	reactionParentClass := r.URL.Query().Get("reactionParentClass")
	// Reads the reaction tables row ID from the request URL
	rowID, err := strconv.Atoi(r.URL.Query().Get("rowID"))

	if err != nil {
		log.Println("There was an issue converting a string to an int in ReactionHandlerGetMethod", err)
		utils.HandleError("There was an issue converting a string to an int in ReactionHandlerGetMethod", err)
	}

	// chooses the correct table name, based on the submitted class
	tableName := ""
	if reactionParentClass == "post" {
		tableName = "POSTREACTIONS"
	} else {
		tableName = "COMMENTREACTIONS"
	}

	// if the submitted row ID was 0 then fear not!
	// the row ID will be the latest entry into the relevant reaction table
	// this function updates the supplied rowID to the new, correct value
	if rowID == 0 {
		rowID = obtainNewRowID(tableName)
	}

	likes, dislikes, err := GetLikesAndDislikes(tableName, rowID)

	if err != nil {
		log.Println("There was a problem with GetLikesAndDislikes in ReactionHandler")
		utils.HandleError("There was a problem with GetLikesAndDislikes in ReactionHandler", err)
	}

	response := struct {
		ReactionID int
		Likes      int
		Dislikes   int
	}{
		ReactionID: rowID,
		Likes:      likes,
		Dislikes:   dislikes,
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

func obtainNewRowID(tableName string) int {

	var rowID int

	query := fmt.Sprintf("SELECT Id FROM %s ORDER BY Id DESC LIMIT 1", tableName)
	err := db.Database.QueryRow(query).Scan(&rowID)

	if err != nil {
		log.Println("Error querying latest ReactionID in ReactionHandlerGetMethod:", err)
		utils.HandleError("Error querying latest ReactionID in ReactionHandlerGetMethod:", err)
	}
	return rowID
}

// Returns the likes and dislikes for a given post/comment, from the relevant table
func GetLikesAndDislikes(tableName string, rowID int) (int, int, error) {

	query := fmt.Sprintf("SELECT Likes, Dislikes FROM %s WHERE Id = ?", tableName)

	var likes, dislikes int
	err := db.Database.QueryRow(query, rowID).Scan(&likes, &dislikes)
	if err != nil {
		log.Println("there was a problem in GetLikesAndDislikes", err)
		utils.HandleError("there was a problem in GetLikesAndDislikes", err)
		return 0, 0, err
	}

	return likes, dislikes, nil
}
