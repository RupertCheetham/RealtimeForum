package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var Database *sql.DB

// initialises database
func InitDatabase() {
	var err error

	Database, err = sql.Open("sqlite3", "./db/realtimeDatabase.db")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to SQLite database")

	// Apply "up" migrations from SQL files
	err = RunMigrations(Database, "./db/migrations", "up")
	if err != nil {
		log.Fatalf("Error applying 'up' migrations: %v", err)
	}
	// schema := string(schemaAsBytes)

	fmt.Println("Migrations applied successfully")
	AddExampleEntries()
	WipeDatabaseOnCommand()
}

func RunMigrations(Database *sql.DB, migrationDir, direction string) error {
	files, err := os.ReadDir(migrationDir)
	if err != nil {
		return err
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
		fmt.Println("migrationPath:", migrationPath)
		sqlBytes, err := os.ReadFile(migrationPath)
		if err != nil {
			return fmt.Errorf("error reading migration file %s: %v", migrationPath, err)
		}

		// fmt.Println(string(sqlBytes))

		_, err = Database.Exec(string(sqlBytes))
		if err != nil {
			return fmt.Errorf("error executing migration %s: %v", migrationPath, err)
		}

		// fmt.Printf("Applied migration: %s\n", migrationPath)
	}

	return nil
}

func isUpMigration(fileName string) bool {
	return len(fileName) > 3 && fileName[len(fileName)-7:] == "_up.sql"
}

func isDownMigration(fileName string) bool {
	return len(fileName) > 5 && fileName[len(fileName)-9:] == "_down.sql"
}
