package main

import (
	"fmt"
	"log"
	"net/http"
	"realtimeForum/db"
	"realtimeForum/handlers"
	"realtimeForum/utils"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db.InitDatabase()
	// log.Println("Database initialized successfully")
	utils.WriteMessageToLogFile("Database initialized successfully")

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/posts", handlers.AddPostHandler)
	http.HandleFunc("/comments", handlers.AddCommentHandler)
	// http.HandleFunc("/getComments", handlers.GetCommentsHandler)
	http.HandleFunc("/registrations", handlers.AddUserHandler)

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
