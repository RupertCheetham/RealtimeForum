package main

import (
	"fmt"
	"log"
	"net/http"
	"realtimeForum/db"
	"realtimeForum/handlers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db.InitDatabase()
	log.Println("Database initialized successfully")

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/posts", handlers.AddPostHandler)
	http.HandleFunc("/registrations", handlers.AddRegistrationHandler)

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
