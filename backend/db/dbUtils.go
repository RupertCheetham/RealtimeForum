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
			// Adds example users to USERS
			err := AddUserToDatabase("Ardek", int(35), "male", "Rupert", "Cheetham", "cheethamthing@gmail.com", "password12345")
			if err != nil {
				log.Fatalf("Error adding entry to USERS table in AddExampleEntries: %v", err)
			}

			err = AddUserToDatabase("john_doe", 30, "Male", "John", "Doe", "john.doe@example.com", "password123")
			if err != nil {
				log.Fatalf("Error adding entry to USERS table in AddExampleEntries: %v", err)
			}

			// Adds example posts to POSTS
			AddPostToDatabase(1, "no-image", "This is the message body", "various, categories")
			if err != nil {
				log.Fatalf("Error adding entry to POST table in AddExampleEntries: %v", err)
			}
			AddPostToDatabase(2, "no-image", "This is the another message body", "various, categories")
			if err != nil {
				log.Fatalf("Error adding entry to POST table in AddExampleEntries: %v", err)
			}
			AddPostToDatabase(3, "no-image", "This is the third body", "category")
			if err != nil {
				log.Fatalf("Error adding entry to POST table in AddExampleEntries: %v", err)
			}
			AddPostToDatabase(4, "no-image", "giggle, giggle, giggle", "various, categories")
			if err != nil {
				log.Fatalf("Error adding entry to POST table in AddExampleEntries: %v", err)
			}

			// Adds example Comments to COMMENTS
			AddCommentToDatabase(1, 1, "This is an example comment")
			if err != nil {
				log.Fatalf("Error adding entry to COMMENT table in AddExampleEntries: %v", err)
			}
			AddCommentToDatabase(2, 3, "This is another example comment")
			if err != nil {
				log.Fatalf("Error adding entry to COMMENT table in AddExampleEntries: %v", err)
			}
			AddCommentToDatabase(3, 4, "This is the third comment")
			if err != nil {
				log.Fatalf("Error adding entry to COMMENT table in AddExampleEntries: %v", err)
			}
			AddCommentToDatabase(4, 1, "chuckle, chuckle, chuckle, chuckle")
			if err != nil {
				log.Fatalf("Error adding entry to COMMENT table in AddExampleEntries: %v", err)
			}
			AddReactionToDatabase("POSTREACTIONS", 1, "dislike")
			AddReactionToDatabase("POSTREACTIONS", 3, "like")
			AddReactionToDatabase("COMMENTREACTIONS", 1, "dislike")
			AddReactionToDatabase("COMMENTREACTIONS", 2, "like")
		}
		log.Println("Example Database entries added successfully")

	}
}
