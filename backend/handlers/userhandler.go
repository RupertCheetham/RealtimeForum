package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"realtimeForum/db"
	"realtimeForum/utils"
)

func GetUserInfoForChatHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	cookieValue := GetCookie(w, r)
	userIDFromSession, _ := db.GetSessionByToken(cookieValue)
	userID := userIDFromSession.UserId

	// userID := r.URL.Query().Get("userId")

	fmt.Println("userID:", userID)

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

		if len(posts) > 0 {
			json.NewEncoder(w).Encode(posts)
		} else {
			w.Write([]byte("No posts available"))
		}
	}
}
