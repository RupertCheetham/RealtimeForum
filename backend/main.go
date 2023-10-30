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
	http.HandleFunc("/api/registrations", auth.RegistrationUserHandler)
	http.HandleFunc("/api/login", auth.LoginHandler)
	http.HandleFunc("/api/logout", auth.LogoutHandler)
	http.HandleFunc("/api/getUsername", handlers.CookieCheck(handlers.GetUsernameHandler, handlers.RequestTimeoutFailedMessage))
	http.HandleFunc("/api/getposts", handlers.CookieCheck(handlers.GetPostHandler, handlers.RequestTimeoutFailedMessage))
	http.HandleFunc("/api/addposts", handlers.CookieCheck(handlers.AddPostHandler, handlers.RequestTimeoutFailedMessage))
	http.HandleFunc("/api/addcomments", handlers.CookieCheck(handlers.AddCommentHandler, handlers.RequestTimeoutFailedMessage))
	// http.HandleFunc("/api/getcomments", handlers.CookieCheck(handlers.GetCommentHandler, handlers.RequestTimeoutFailedMessage))
	http.HandleFunc("/api/getcomments", handlers.GetCommentHandler)
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
