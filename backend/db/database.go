package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"realtimeForum/utils"

	_ "github.com/mattn/go-sqlite3"
)

var Database *sql.DB

// initialises database
func InitDatabase() {
	var err error

	Database, err = sql.Open("sqlite3", "./db/realtimeDatabase.db")
	if err != nil {
		utils.HandleError("Unable to open database", err)
		log.Println("Unable to open database", err)
	}

	utils.WriteMessageToLogFile("Connected to SQLite database")

	// Apply "up" migrations from SQL files
	RunMigrations(Database, "./db/migrations", "up")
	if err != nil {
		utils.HandleError("Error applying 'up' migrations: ", err)
		log.Println("Error applying 'up' migrations: ", err)
	}

	// fmt.Println("Migrations applied successfully")

	AddExampleEntries()
	DeleteUserTest()
	DeleteAllUsersTest()
	WipeDatabaseOnCommand()
}

// Applies "up" migrations from SQL files
func RunMigrations(Database *sql.DB, migrationDir, direction string) {
	files, err := os.ReadDir(migrationDir)
	if err != nil {
		utils.HandleError("Error reading migration directory", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		if direction == "up" && !isUpMigration(fileName) {
			continue
		}
		if direction == "down" && !isDownMigration(fileName) {
			continue
		}

		migrationPath := migrationDir + "/" + fileName

		sqlBytes, err := os.ReadFile(migrationPath)
		if err != nil {
			message := fmt.Sprintf("error reading migration file %s", migrationPath)
			utils.HandleError(message, err)
		}

		_, err = Database.Exec(string(sqlBytes))
		if err != nil {
			message := fmt.Sprintf("error executing migration %s:", migrationPath)
			utils.HandleError(message, err)
		}

	}
}

func isUpMigration(fileName string) bool {
	return len(fileName) > 3 && fileName[len(fileName)-7:] == "_up.sql"
}

func isDownMigration(fileName string) bool {
	return len(fileName) > 5 && fileName[len(fileName)-9:] == "_down.sql"
}
