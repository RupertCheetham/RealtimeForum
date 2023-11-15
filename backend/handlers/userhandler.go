package handlers

import (
	"encoding/json"
	"net/http"
	"realtimeForum/db"
	"realtimeForum/utils"
)

func GetPostsForSpecificUser(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	cookieValue := GetCookie(w, r)
	userIDFromSession, _ := db.GetSessionByToken(cookieValue)
	userID := userIDFromSession.UserId

	// This code block is handling the logic for retrieving posts from the database when the HTTP request
	// method is GET.
	if r.Method == http.MethodGet {
		posts, err := db.GetAllUserPostsAndCommentsFromDatabase(userID)
		//posts, err := db.GetPostFromDatabase()
		if err != nil {
			utils.HandleError("Problem getting posts from db in GetPostHandler", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(posts)
	}
}
