package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var Database *sql.DB

// initialises database
func initDatabase() {
	var err error

	wipeDatabaseOnCommand()

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

	// Create the registration table if it doesn't exist
	_, err = Database.Exec(`
	CREATE TABLE IF NOT EXISTS registration (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nickname TEXT,
		age INTEGER,
		gender TEXT,
		first_name Text,
		last_name TEXT,
		email TEXT,
		password TEXT
	);
	`)
	if err != nil {
		log.Fatal(err)
	}
	addExampleEntries()
}

// deletes database if first arg is 'new'
func wipeDatabaseOnCommand() {
	if len(os.Args) > 1 {
		if os.Args[1] == "new" {
			os.Remove("../db/realtimeDatabase.db")
			fmt.Println("Deleted realtimeDatabase.db")
		}
	}
}

// adds example posts if first arg is 'new'
func addExampleEntries() {
	if len(os.Args) > 1 {
		if os.Args[1] == "new" {
			addRegistrationToDatabase("Ardek", int(35), "male", "Rupert", "Cheetham", "cheethamthing@gmail.com", "password12345")
			addRegistrationToDatabase("john_doe", 30, "Male", "John", "Doe", "john.doe@example.com", "password123")

			addPostToDatabase("Ardek", "no-image", "This is the message body", "various, categories")
			addPostToDatabase("Nikoi", "no-image", "This is the another message body", "various, categories")
			addPostToDatabase("Martin", "no-image", "This is the third body", "category")
			addPostToDatabase("Mike", "no-image", "giggle, giggle, giggle", "various, categories")
		}
	}
}

func addRegistrationToDatabase(nickname string, age int, gender string, firstName string, lastName string, email string, password string) error {
	_, err := Database.Exec("INSERT INTO registration (nickname, age, gender, first_name, last_name, email, password) VALUES (?, ?, ?, ?, ?, ?, ?)", nickname, age, gender, firstName, lastName, email, password)
	if err != nil {
		log.Println("Error adding registration to database:", err)
	}
	return err
}

func getRegistrationFromDatabase() ([]RegistrationEntry, error) {
	rows, err := Database.Query("SELECT nickname, age, gender, first_name, last_name, email, password FROM registration ORDER BY id ASC")
	if err != nil {
		log.Println("Error querying registrations from database:", err)
		return nil, err
	}
	defer rows.Close()

	var registrations []RegistrationEntry
	for rows.Next() {
		var entry RegistrationEntry
		err := rows.Scan(&entry.Nickname, &entry.Age, &entry.Gender, &entry.FirstName, &entry.LastName, &entry.Email, &entry.Password)
		if err != nil {
			log.Println("Error scanning row from database:", err)
			return nil, err
		}
		registrations = append(registrations, entry)
	}

	return registrations, nil
}

// adds a post to the database
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

// retrieves all posts from database and returns them
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
