package forum

import (
	"fmt"
	"net/http"
)

// Boots up a apge that displays the current user's posts and liked posts
func MyActivity(w http.ResponseWriter, r *http.Request) {

	// checks that the user session is still running (hasn't logged in elsewhere)
	cookie, _ := r.Cookie("session")
	if UserSessionMissing(cookie.Value) {

		cookie.Value = ""         //sets cookie value to an empty string
		cookie.MaxAge = -1        //sets the existance of the cookie to -1
		http.SetCookie(w, cookie) //sets the modified values of the cookie
		fmt.Println("Session deleted")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var submittedStruct Activity
	submittedStruct.NoPostsMessage = ""
	submittedStruct.NoLikedPostsMessage = ""
	errorMessage := "There are no posts here"
	_, cookieValue, _ := loggedIn(w, r)
	UserID := cookieUserID(cookieValue)
	Username, _ := getUsernameFromUserID(UserID)

	submittedStruct.LoggedInUsername = Username
	submittedStruct.MyPosts = FilterPostsByUserID(UserID)
	if len(submittedStruct.MyPosts) == 0 {
		submittedStruct.NoPostsMessage = errorMessage
	}
	submittedStruct.MyLikedPosts = FilterLikedPostsByUserID(UserID)
	if len(submittedStruct.MyLikedPosts) == 0 {
		submittedStruct.NoLikedPostsMessage = errorMessage
	}
	tmpl.ExecuteTemplate(w, "myActivity.html", submittedStruct)
}
