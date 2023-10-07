package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"realtimeForum/db"
	"realtimeForum/utils"
)

func GetNewestPostHandler(w http.ResponseWriter, r *http.Request) {

	SetupCORS(&w, r)

	post, err := db.GetNewestPost()
	if err != nil {
		utils.HandleError("Problem getting post from db in GetNewestPostHandler", err)
		log.Println("Problem getting post from db in GetNewestPostHandler", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(post)
	json.NewEncoder(w).Encode(post)

}
