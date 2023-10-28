package main

import (
	"fmt"
	"log"
	"net/http"
	"realtimeForum/auth"
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
	http.HandleFunc("/api/auth", auth.LoginHandler)
	http.HandleFunc("/api/registrations", auth.AddUserHandler)
	http.HandleFunc("/api/getUsername", handlers.CookieCheck(handlers.GetUsernameHandler, handlers.ActionFailedMessage))
	http.HandleFunc("/api/getposts", handlers.CookieCheck(handlers.GetPostHandler, handlers.ActionFailedMessage))
	http.HandleFunc("/api/addposts", handlers.CookieCheck(handlers.AddPostHandler, handlers.ActionFailedMessage))
	http.HandleFunc("/api/addcomments", handlers.CookieCheck(handlers.AddCommentHandler, handlers.ActionFailedMessage))
	http.HandleFunc("/api/getcomments", handlers.CookieCheck(handlers.GetCommentHandler, handlers.ActionFailedMessage))
	http.HandleFunc("/reaction", handlers.ReactionHandler)
	http.HandleFunc("/chat", handlers.ChatHandler)
	http.HandleFunc("/getChatHistory", handlers.GetChatHistoryHandler)

	// Specify the paths to your TLS certificate and private key files
	certFile := "server.crt"
	keyFile := "server.key"
	fmt.Printf("Starting server at port 8080\n")
	err := http.ListenAndServeTLS(":8080", certFile, keyFile, nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}
