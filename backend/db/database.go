package db

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var Database *sql.DB

// initialises database
func InitDatabase() {
	var err error

	wipeDatabaseOnCommand()

	Database, err = sql.Open("sqlite3", "db/realtimeDatabase.db")
	if err != nil {
		log.Fatal(err)
	}

	// Open the schema.sql file
	schemaFile, err := os.Open("db/schema.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer schemaFile.Close()

	// Read the SQL statements from the file
	schemaAsBytes, err := io.ReadAll(schemaFile)
	if err != nil {
		log.Fatal(err)
	}
	schema := string(schemaAsBytes)

	// Execute the schema statements to create the tables
	_, err = Database.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
	addExampleEntriesOnCommand()
}

// deletes database if first arg is 'new'
func wipeDatabaseOnCommand() {
	if len(os.Args) > 1 {
		if os.Args[1] == "new" {
			os.Remove("db/realtimeDatabase.db")
			fmt.Println("Deleted realtimeDatabase.db")
		}
	}
}

// adds example entries if first arg is 'new'
func addExampleEntriesOnCommand() {
	if len(os.Args) > 1 {
		if os.Args[1] == "new" {
			AddRegistrationToDatabase("Ardek", int(35), "male", "Rupert", "Cheetham", "cheethamthing@gmail.com", "password12345")
			AddRegistrationToDatabase("john_doe", 30, "Male", "John", "Doe", "john.doe@example.com", "password123")

			AddPostToDatabase("Ardek", "no-image", "This is the message body", "various, categories")
			AddPostToDatabase("Nikoi", "no-image", "This is the another message body", "various, categories")
			AddPostToDatabase("Martin", "no-image", "This is the third body", "category")
			AddPostToDatabase("Mike", "no-image", "giggle, giggle, giggle", "various, categories")

			AddCommentToDatabase(int(1), "Ardek", "This is a comment")
			AddCommentToDatabase(int(3), "Nikoi", "This is another comment")
			AddCommentToDatabase(int(111), "Martin", "This is the third comment")
			AddCommentToDatabase(int(6), "Mike", "giggle, giggle, giggle")
		}
	}
}

// adds registration to the registrations table in database
func AddRegistrationToDatabase(username string, age int, gender string, firstName string, lastName string, email string, password string) error {
	_, err := Database.Exec("INSERT INTO REGISTRATION (username, age, gender, firstName, lastName, email, password) VALUES (?, ?, ?, ?, ?, ?, ?)", username, age, gender, firstName, lastName, email, password)
	if err != nil {
		log.Println("Error adding registration to REGISTRATION in database:", err)
	}
	return err
}

func AddCommentToDatabase(postID int, username string, body string) error {
	var likes = 0
	var dislikes = 0
	var whoLiked = ""
	var whoDisliked = ""
	_, err := Database.Exec("INSERT INTO COMMENTS (postID, username, body, likes, dislikes, whoLiked, whoDisliked) VALUES (?, ?, ?, ?, ?, ?, ?)", postID, username, body, likes, dislikes, whoLiked, whoDisliked)
	if err != nil {
		log.Println("Error adding comment to COMMENTS in AddCommentToDatabase:", err)
	}
	return err
}

// retrieves all registrations from database and returns them as an array of RegistrationEntry
func GetRegistrationFromDatabase() ([]RegistrationEntry, error) {
	rows, err := Database.Query("SELECT username, age, gender, firstName, lastName, email, password FROM registration ORDER BY id ASC")
	if err != nil {
		log.Println("Error querying registrations from database in GetRegistrationFromDatabase:", err)
		return nil, err
	}
	defer rows.Close()

	var registrations []RegistrationEntry
	for rows.Next() {
		var entry RegistrationEntry
		err := rows.Scan(&entry.Username, &entry.Age, &entry.Gender, &entry.FirstName, &entry.LastName, &entry.Email, &entry.Password)
		if err != nil {
			log.Println("Error scanning row from database:", err)
			return nil, err
		}
		registrations = append(registrations, entry)
	}

	return registrations, nil
}

// adds a post to the database
func AddPostToDatabase(username string, img string, body string, categories string) error {
	var likes = 0
	var dislikes = 0
	var whoLiked = ""
	var whoDisliked = ""
	_, err := Database.Exec("INSERT INTO POSTS (username, img, body, categories, likes, dislikes, whoLiked, whoDisliked) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", username, img, body, categories, likes, dislikes, whoLiked, whoDisliked)
	if err != nil {
		log.Println("Error adding post to database in addPostToDatabase:", err)
	}
	return err
}

// retrieves all posts from database and returns them
func GetPostFromDatabase() ([]PostEntry, error) {
	rows, err := Database.Query("SELECT id, username, img, Body, categories, creationDate, Llikes, dislikes, whoLiked, whoDisliked FROM POSTS ORDER BY Id ASC")
	if err != nil {
		log.Println("Error querying posts from database:", err)
		return nil, err
	}
	defer rows.Close()

	var posts []PostEntry
	for rows.Next() {
		var post PostEntry
		err := rows.Scan(&post.Id, &post.Username, &post.Img, &post.Body, &post.Categories, &post.CreationDate, &post.Likes, &post.Dislikes, &post.WhoLiked, &post.WhoDisliked)
		if err != nil {
			log.Println("Error scanning row from database in GetPostFromDatabase:", err)
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
