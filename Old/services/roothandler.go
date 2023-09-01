package forum

import (
	"fmt"
	"net/http"
)

func roothandler(w http.ResponseWriter, r *http.Request) {

	// checks if the user is logged in
	isLoggedin, cookieValue, _ := loggedIn(w, r)
	fmt.Println("isLoggedin:", isLoggedin)

	if isLoggedin {
		// checks to see if the user has logged in elsewhere (thus the session has been deleted)
		if UserSessionMissing(cookieValue) {
			// if yes then deletes that sessions cookie
			cookie, _ := r.Cookie("session")
			cookie.Value = ""         //sets cookie value to an empty string
			cookie.MaxAge = -1        //sets the existance of the cookie to -1
			http.SetCookie(w, cookie) //sets the modified values of the cookie
			fmt.Println("Session deleted")
			// and changees the status to logged out
			isLoggedin = false
		}
	}
	var submittedStruct FinalStruct
	// gets the posts from the database
	submittedStruct.Posts = getPostsFromDB()

	//If logged in then adds username to the data submitted to html
	if isLoggedin {

		username, _ := getUsernameFromUserID(cookieUserID(cookieValue))
		submittedStruct.LoggedInUsername = username

		//If *not* logged in then username is "You are not logged in"
	} else {

		submittedStruct.LoggedInUsername = "You are not logged in"

	}

	tmpl.ExecuteTemplate(w, "index.html", submittedStruct)
}
