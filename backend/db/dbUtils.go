package db

import (
	"fmt"
	"log"
	"os"
)

// files, err := os.ReadDir("./db/migrations")
// fmt.Println(files)
// if err != nil {
// 	log.Fatal("unable to read migrations:", err)
// }

// deletes database if first arg is 'new'
func WipeDatabaseOnCommand() {
	if len(os.Args) > 1 {
		if os.Args[1] == "new" {
			// Rollback the last migration (uncomment if needed)
			err := RunMigrations(Database, "./db/migrations", "down")
			if err != nil {
				log.Fatalf("Error applying 'down' migrations: %v", err)
			}
			fmt.Println("Dropped all tables")
		}
	}
}

// adds example posts if first arg is 'new'
func AddExampleEntries() {
	if len(os.Args) > 1 {
		if os.Args[1] == "test" {
			err := AddUserToDatabase("Ardek", int(35), "male", "Rupert", "Cheetham", "cheethamthing@gmail.com", "password12345")
			if err != nil {
				log.Fatalf("Error adding entry to USERS table in AddExampleEntries: %v", err)
			}

			err = AddUserToDatabase("john_doe", 30, "Male", "John", "Doe", "john.doe@example.com", "password123")
			if err != nil {
				log.Fatalf("Error adding entry to USERS table in AddExampleEntries: %v", err)
			}

			AddPostToDatabase("Ardek", "no-image", "This is the message body", "various, categories")
			if err != nil {
				log.Fatalf("Error adding entry to POST table in AddExampleEntries: %v", err)
			}
			_, _ = Database.Exec("UPDATE POSTS SET Likes = ?, Dislikes = ?, WhoLiked = ?, WhoDisliked = ? WHERE Id = ?", 1, 2, "hello, there, ghghg", "dgf, sdfgsdfg, ertret", 1)
			AddPostToDatabase("Nikoi", "no-image", "This is the another message body", "various, categories")
			if err != nil {
				log.Fatalf("Error adding entry to POST table in AddExampleEntries: %v", err)
			}
			AddPostToDatabase("Martin", "no-image", "This is the third body", "category")
			if err != nil {
				log.Fatalf("Error adding entry to POST table in AddExampleEntries: %v", err)
			}
			AddPostToDatabase("Mike", "no-image", "giggle, giggle, giggle", "various, categories")
			if err != nil {
				log.Fatalf("Error adding entry to POST table in AddExampleEntries: %v", err)
			}

			AddCommentToDatabase("Ardek", 111, "This is an example comment")
			if err != nil {
				log.Fatalf("Error adding entry to COMMENT table in AddExampleEntries: %v", err)
			}
			AddCommentToDatabase("Nikoi", 111, "This is another example comment")
			if err != nil {
				log.Fatalf("Error adding entry to COMMENT table in AddExampleEntries: %v", err)
			}
			AddCommentToDatabase("Martin", 111, "This is the third comment")
			if err != nil {
				log.Fatalf("Error adding entry to COMMENT table in AddExampleEntries: %v", err)
			}
			AddCommentToDatabase("Mike", 111, "chuckle, chuckle, chuckle, chuckle")
			if err != nil {
				log.Fatalf("Error adding entry to COMMENT table in AddExampleEntries: %v", err)
			}
		}
		log.Println("Example Database entries added successfully")
	}
}
