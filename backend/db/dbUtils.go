package db

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
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
			hashPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
			if err != nil {
				log.Fatalf("Error bcrypting in AddExampleEntries: %v", err)
			}
			err = AddUserToDatabase("Ardek", int(35), "male", "Rupert", "Cheetham", "cheethamthing@gmail.com", string(hashPassword))
			if err != nil {
				log.Fatalf("Error adding entry to USERS table in AddExampleEntries: %v", err)
			}

			hashPassword, _ = bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
			err = AddUserToDatabase("Knikoi", 40, "Male", "Kwashie", "Nikoi", "nikoi.doe@example.com", string(hashPassword))
			if err != nil {
				log.Fatalf("Error adding entry to USERS table in AddExampleEntries: %v", err)
			}

			hashPassword, _ = bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
			err = AddUserToDatabase("MFenton", 35, "Male", "Martin", "Fenton", "martin.doe@example.com", string(hashPassword))
			if err != nil {
				log.Fatalf("Error adding entry to USERS table in AddExampleEntries: %v", err)
			}

			hashPassword, _ = bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
			err = AddUserToDatabase("Madeleke", 20, "Male", "Mike", "A", "mike.doe@example.com", string(hashPassword))
			if err != nil {
				log.Fatalf("Error adding entry to USERS table in AddExampleEntries: %v", err)
			}

			hashPassword, _ = bcrypt.GenerateFromPassword([]byte("t"), bcrypt.DefaultCost)
			err = AddUserToDatabase("t", 255, "Male", "Mr", "E", "john.doe@example.com", string(hashPassword))
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
			AddReactionToDatabase("post", 1, 1, "dislike")
			AddReactionToDatabase("post", 2, 3, "like")
			AddReactionToDatabase("comment", 1, 1, "dislike")
			AddReactionToDatabase("comment", 2, 2, "like")
		}
		log.Println("Example Database entries added successfully")

	}
}

func DeleteUserTest() {
	if len(os.Args) == 3 {
		if os.Args[1] == "delete" && os.Args[2] == "user" {
			// delete user
			err := DeleteUserFromDatabase("b")
			if err != nil {
				log.Fatalf("Error adding entry to USERS table in DeleteUserFromDatabase: %v", err)
			}
		}
	}
}

func DeleteAllUsersTest() {
	if len(os.Args) == 4 {
		if os.Args[1] == "delete" && os.Args[2] == "all" && os.Args[3] == "users" {
			// delete user
			err := DeleteAllUsersFromDatabase()
			if err != nil {
				log.Fatalf("Error adding entry to USERS table in DeleteAllUsersFromDatabase: %v", err)
			}
		}
	}
}
