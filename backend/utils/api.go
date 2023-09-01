package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type PostEntry struct {
	Id           int    `json:"Id"`
	Nickname     string `json:"Nickname"`
	Img          string `json:"Img"`
	Body         string `json:"Body"`
	Categories   string `json:"Categories"`
	CreationDate string `json:"CreationDate"`
	Likes        int    `json:"Likes"`
	Dislikes     int    `json:"Dislikes"`
	WhoLiked     string `json:"WhoLiked"`
	WhoDisliked  string `json:"WhoDisliked"`
}

func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func AddPostHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers for this handler
	setupCORS(&w, r)

	if r.Method == "POST" {
		var post PostEntry
		err := json.NewDecoder(r.Body).Decode(&post)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Println("Received post:", post.Nickname, post.Img, post.Body, post.Categories)

		err = addPostToDatabase(post.Nickname, post.Img, post.Body, post.Categories)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}

	if r.Method == "GET" {
		posts, err := getPostFromDatabase()
		if err != nil {
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

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Real Time Forum API"))
}

func main() {
	initDatabase()
	log.Println("Database initialized successfully")

	addPostToDatabase("Ardek", "no-image", "This is the message body", "various, categories")
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/posts", AddPostHandler)

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
