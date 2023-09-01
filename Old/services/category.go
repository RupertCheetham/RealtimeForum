package forum

import (
	"fmt"
	"log"
	"net/http"
)

// Boots up a apge that displays the current user's posts and liked posts
func CategoryPage(w http.ResponseWriter, r *http.Request) {

	var submittedStruct Category
	submittedStruct.NoPostsMessage = ""
	errorMessage := "There are no posts here"
	isLoggedIn, cookieValue, _ := loggedIn(w, r)
	if isLoggedIn {
		UserID := cookieUserID(cookieValue)
		Username, _ := getUsernameFromUserID(UserID)

		submittedStruct.LoggedInUsername = Username
	} else {
		submittedStruct.LoggedInUsername = "You are not logged in"
	}
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	//get user info
	category := r.FormValue("category")
	fmt.Println("Category:", category)
	submittedStruct.Category = category
	submittedStruct.Posts = FilterPostsByCategory(category)
	if len(submittedStruct.Posts) == 0 {
		submittedStruct.NoPostsMessage = errorMessage
	}

	tmpl.ExecuteTemplate(w, "categoryPage.html", submittedStruct)
}
