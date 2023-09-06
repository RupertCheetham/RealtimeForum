package db

import (
	"fmt"
	"log"
	"os"
	"realtimeForum/utils"
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
				utils.HandleError("Error applying 'down' migrations: %v", err)
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
			AddRegistrationToDatabase("Ardek", int(35), "male", "Rupert", "Cheetham", "cheethamthing@gmail.com", "password12345")
			AddRegistrationToDatabase("john_doe", 30, "Male", "John", "Doe", "john.doe@example.com", "password123")

			AddPostToDatabase("Ardek", "no-image", "This is the message body", "various, categories")
			AddPostToDatabase("Nikoi", "no-image", "This is the another message body", "various, categories")
			AddPostToDatabase("Martin", "no-image", "This is the third body", "category")
			AddPostToDatabase("Mike", "no-image", "giggle, giggle, giggle", "various, categories")

			AddCommentToDatabase("Ardek", 111, "This is an example comment")
			AddCommentToDatabase("Nikoi", 111, "This is another example comment")
			AddCommentToDatabase("Martin", 111, "This is the third comment")
			AddCommentToDatabase("Mike", 111, "chuckle, chuckle, chuckle, chuckle")

		}
		utils.WriteMessageToLogFile("Example Database entries added successfully")
	}
}
