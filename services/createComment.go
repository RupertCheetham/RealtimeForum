package forum

import (
	"log"
	"net/http"
)

// Reads a users submitted comment data, inserts it into the table (using InsertToCommentTable())
func CreateUserComment(w http.ResponseWriter, r *http.Request) {

	_, sessionID, err := loggedIn(w, r)
	if err != nil {
		log.Println("Error retrieving sessionid in CreateUserComment(): ", err)
	}

	UserID, err := GetUserIDfrom("cookies", "sessionid", sessionID)
	if err != nil {
		log.Println("Error retrieving Username in CreateUserComment(): ", UserID)
	}

	err = r.ParseMultipartForm(maxFileSize)
	if err != nil {
		http.Error(w, "Unable to parse form in CreateUserComment()", http.StatusBadRequest)
		return
	}

	userText := r.PostFormValue("usertxt")
	parentPostID := r.PostFormValue("postID")

	//insertion to db
	InsertToCommentTable(parentPostID, UserID, userText)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

// Insers comment to the table
func InsertToCommentTable(parentPostID string, Username string, body string) {

	likes := 0
	dislikes := 0
	whoLiked := ""
	whoDisliked := ""

	query := "INSERT INTO comments(postID, UserID, body, likes, dislikes, whoLiked, whoDisliked) VALUES(?, ?, ?, ?, ?, ?, ?)"
	stmt, err := Database.Prepare(query)
	if err != nil {
		log.Println("error is: ", err)
	}
	_, err = stmt.Exec(parentPostID, Username, body, likes, dislikes, whoLiked, whoDisliked)
	if err != nil {
		log.Fatal(err)
	}
}
