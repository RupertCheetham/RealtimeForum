package forum

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	maxFileSize = 20 << 20 // 20MB
	dirPath     = "static/uploadFiles/images"
)

var supportedFileTypes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
	"image/gif":  true,
}

// Reads a users submitted post data, inserts it into the table (using InsertToPostTable())
// Then loads up the main page again, with the fresh post
func CreateUserPost(w http.ResponseWriter, r *http.Request) {

	//if user logged in get the session id
	_, CookieValue, err := loggedIn(w, r)
	if err != nil {
		log.Println("Error retrieving sessionid  in createPost: ", err)
	}

	// checks that the user hasn't logged in elsewhere
	if UserSessionMissing(CookieValue) {
		cookie, _ := r.Cookie("session")
		cookie.Value = ""         //sets cookie value to an empty string
		cookie.MaxAge = -1        //sets the existance of the cookie to -1
		http.SetCookie(w, cookie) //sets the modified values of the cookie
		fmt.Println("Session deleted")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	//get the user id from cookies table of the corresponding session id
	userID, err := GetUserIDfrom("cookies", "sessionid", CookieValue)
	fmt.Println("User ID: ", userID)
	if err != nil {
		log.Println("Error retrieving userid in createPost: ", userID)
	}

	//parse the form
	err = r.ParseMultipartForm(maxFileSize)
	if err != nil {
		http.Error(w, "Unable to parse form  in createPost", http.StatusBadRequest)
		return
	}

	userText := r.PostFormValue("usertxt")
	if userText == "" {
		userText = " "
	}

	//get the file from the form
	file, fileHeader, err := r.FormFile("usrfile")
	if err != nil && err != http.ErrMissingFile {
		log.Fatal("file missing in createPost", err)
	}
	//this anonymous func will run after the code below has executed and close the file
	defer func() {
		if file != nil {
			file.Close()
		}
	}() //calling the func

	categories := r.Form["Category"]

	fmt.Println("category:", categories)
	if userText == " " && file == nil {
		fmt.Println("MESSAGE FOR NO CONTENT IN POST")
		// http.Redirect(w, r, "/", http.StatusBadRequest)
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	//if file is given
	if file != nil {
		if fileHeader.Size > maxFileSize {
			log.Println("File is too big!!")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		} else if !supportedFileTypes[fileHeader.Header.Get("Content-Type")] {
			//todo display this error to user instead of printing it to console
			log.Println("File type is not supported!!")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		//create a file in the given directory with the suffix .jpg
		osFile, err := os.CreateTemp(dirPath, "*.jpg")
		if err != nil {
			log.Fatal("Error creating file: ", err)
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer osFile.Close()
		//copy the contents of the file to the file created above
		_, err = io.Copy(osFile, file)
		if err != nil {
			log.Fatal("Error copying file: ", err)
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		fmt.Println("file stored!")

		//get relative path of the file
		relativePath := osFile.Name()
		// Insertion to db with file

		InsertToPostTable(userID, relativePath, userText, categories)
		fmt.Println("POST INSERTED WITH FILE")
	} else {
		//insert to db without file
		// `InsertToPostTable(userID, "no file given", userText)` is inserting a post into the database
		// without a file. It is passing the userID, "no file given" as the img, and the userText as the
		// body to the `InsertToPostTable` function. The function then inserts the post into the database
		// with the given parameters and prints "POST INSERTED WITHOUT FILE" to the console.
		InsertToPostTable(userID, "static/uploadFiles/images/Connectify.jpg", userText, categories)
		fmt.Println("POST INSERTED WITHOUT FILE")
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Insers Post to Table
func InsertToPostTable(UserID, img, body string, categories []string) {
	var categoriesString string
	if len(categories) > 0 {
		for _, category := range categories {

			categoriesString = categoriesString + "," + category
		}
		categoriesString = categoriesString[1:]
	} else {
		categoriesString = "None"
	}
	likes := 0
	dislikes := 0
	whoLiked := ""
	whoDisliked := ""

	//default behaviour
	query := "INSERT INTO posts(UserID, img, body, categories, likes, dislikes, whoLiked, whoDisliked) VALUES(?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := Database.Prepare(query)
	if err != nil {
		log.Println("error is: ", err)
	}
	_, err = stmt.Exec(UserID, img, body, categoriesString, likes, dislikes, whoLiked, whoDisliked)
	if err != nil {
		log.Fatal(err)
	}
}
