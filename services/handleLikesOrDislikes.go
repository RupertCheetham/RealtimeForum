package forum

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Does the likes and dislikes, then navigates back to the correct page
// We had to get the html to submit variables related to what page they're currently on,
// We could use this to create if statements in the backend to load up the correct page depending upon where the data came from
func HandleLikesOrDislikes(w http.ResponseWriter, r *http.Request) {

	var updateAccountSQL string

	isLoggedin, cookievalue, _ := loggedIn(w, r)

	if !isLoggedin {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	userID := cookieUserID(cookievalue)

	postNum := r.FormValue("postNum")
	likedOrDislike := r.FormValue("likeOrDislike")
	fmt.Println("post number", postNum, "received a", likedOrDislike)

	var currentLikes int
	var currentDislikes int
	var whoLiked string
	var whoDisliked string

	if (r.FormValue("type") == "post") || (r.FormValue("type") == "myActivity") {
		err = Database.QueryRow("SELECT Likes, Dislikes, WhoLiked, WhoDisliked FROM Posts WHERE PostID = ?", postNum).Scan(&currentLikes, &currentDislikes, &whoLiked, &whoDisliked)
	} else {
		err = Database.QueryRow("SELECT Likes, Dislikes, WhoLiked, WhoDisliked FROM Comments WHERE CommentID = ?", postNum).Scan(&currentLikes, &currentDislikes, &whoLiked, &whoDisliked)
	}

	if err != nil {
		log.Fatalln(err.Error())
	}

	// This is a bit of a journey... Check inside value tweaker for something that may help or make it worse
	if likedOrDislike == "like" {
		whoLiked, whoDisliked, currentLikes, currentDislikes = valueTweaker(userID, whoLiked, whoDisliked, currentLikes, currentDislikes)
	}

	if likedOrDislike == "dislike" {
		whoDisliked, whoLiked, currentDislikes, currentLikes = valueTweaker(userID, whoDisliked, whoLiked, currentDislikes, currentLikes)
	}

	if (r.FormValue("type") == "post") || (r.FormValue("type") == "myActivity") {
		updateAccountSQL = `UPDATE Posts SET Likes = ?, Dislikes = ?, WhoLiked = ?, WhoDisliked = ? WHERE PostID = ?`
	} else {
		updateAccountSQL = `UPDATE Comments SET Likes = ?, Dislikes = ?, WhoLiked = ?, WhoDisliked = ? WHERE CommentID = ?`
	}

	statement, err := Database.Prepare(updateAccountSQL)

	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(currentLikes, currentDislikes, whoLiked, whoDisliked, postNum)

	if err != nil {
		log.Fatalln(err.Error())
	}
	if r.FormValue("type") == "myActivity" {
		http.Redirect(w, r, "myActivity", http.StatusSeeOther)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Tweaks the values of likes/dislikes as needed
func valueTweaker(userID string, who1 string, who2 string, value1 int, value2 int) (string, string, int, int) {

	// checks to see if the like (or dislike) is a repeat action, if it is then returns values unchanged
	splitA := strings.Split(who1, ",")
	for _, idAccount := range splitA {
		if idAccount == userID {
			return who1, who2, value1, value2
		}
	}

	// checks to see if the oppsite action has already taken place, if it has then returns removes the action from the db
	splitB := strings.Split(who2, ",")
	for i, idAccount := range splitB {
		if idAccount == userID {
			value2 = value2 - 1
			// removes userID so they're no longer on the list as having performed the opposite action
			splitB = append(splitB[:i], splitB[i+1:]...)
		}
	}

	// performs the action (like/dislike)
	who2 = strings.Join(splitB, ",")
	// adds userID to whoLiked or whoDisliked
	who1 = who1 + "," + userID
	value1 = value1 + 1

	return who1, who2, value1, value2

}

// Not entirely sure why this is here
// Returns the comment's parent post ID
func getPostIDFromCommentID(CommentID string) string {
	var PostID string
	query := fmt.Sprintf("SELECT PostID FROM Comments WHERE CommentID = ?")
	row := Database.QueryRow(query, CommentID)
	if err := row.Scan(&PostID); err != nil {
		return "There is an error"
	}
	return PostID
}
