package forum

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var tmpl *template.Template
var Database *sql.DB
var err error

// var w http.ResponseWriter
// var r *http.Request

func init() {

	// if run with the argument "new" then will delete the DB; a new DB is created a few lines down
	// Will also remove all uploaded pictures, except for the logo 'Connectify.jpg'
	if len(os.Args) > 1 {
		if os.Args[1] == "new" {
			// LogOut(w, r)
			os.Remove("./dal/forum.db")
			fmt.Println("Deleted forum.db")

			dirPath := "static/uploadFiles/images"

			// Get a list of all files in the directory
			files, err := ioutil.ReadDir(dirPath)
			if err != nil {
				fmt.Println("Failed to read directory:", err)
				return
			}

			// Iterate over the files and remove them, except for "Connectify.jpg"
			for _, file := range files {
				filePath := filepath.Join(dirPath, file.Name())
				if file.Name() != "Connectify.jpg" {
					err := os.Remove(filePath)
					if err != nil {
						fmt.Println("Failed to remove file:", err)
					} else {
						fmt.Println("Removed file:", filePath)
					}
				}
			}

		}
	}
	//Open the SQlite database, create it if doesn't exist.
	Database, err = sql.Open("sqlite3", "./dal/forum.db")
	if err != nil {
		log.Fatal(err)
	}

	// Open the schema.sql file
	file, err := os.Open("./dal/schema.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the SQL statements from the file
	sqlBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	sql := string(sqlBytes)

	// Execute the SQL statements to create the tables
	_, err = Database.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
	//parse all html files
	tmpl = template.Must(template.ParseGlob("templates/*.html"))

}
