package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Database *sql.DB

func initDatabase() {
	var err error
	Database, err = sql.Open("sqlite3", "../db/realtimeDatabase.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create the post table if it doesn't exist
	_, err = Database.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			Id INTEGER PRIMARY KEY AUTOINCREMENT,
			Nickname TEXT,
			Img TEXT,
			Body TEXT,
			Categories Text,
			CreationDate DATETIME DEFAULT CURRENT_TIMESTAMP,
			Likes INTEGER,
			Dislikes INTEGER,
			WhoLiked TEXT,
			WhoDisliked TEXT
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func addPostToDatabase(nickname string, img string, body string, categories string) error {
	var likes = 0
	var dislikes = 0
	var whoLiked = ""
	var whoDisliked = ""
	_, err := Database.Exec("INSERT INTO posts (Nickname, Img, Body, Categories, Likes, Dislikes, WhoLiked, WhoDisliked) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", nickname, img, body, categories, likes, dislikes, whoLiked, whoDisliked)
	if err != nil {
		log.Println("Error adding post to database in addPostToDatabase:", err)
	}
	return err
}

func getPostFromDatabase() ([]PostEntry, error) {
	rows, err := Database.Query("SELECT Id, Nickname, Img, Body, Categories, CreationDate, Likes, Dislikes, WhoLiked, WhoDisliked FROM posts ORDER BY Id ASC")
	if err != nil {
		log.Println("Error querying posts from database:", err)
		return nil, err
	}
	defer rows.Close()

	var posts []PostEntry
	for rows.Next() {
		var post PostEntry
		err := rows.Scan(&post.Id, &post.Nickname, &post.Img, &post.Body, &post.Categories, &post.CreationDate, &post.Likes, &post.Dislikes, &post.WhoLiked, &post.WhoDisliked)
		if err != nil {
			log.Println("Error scanning row from database:", err)
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
