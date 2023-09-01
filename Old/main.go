package main

import (
	f "forum/services"
	"net/http"
)

func main() {

	http.HandleFunc("/", f.ServePage)
	http.HandleFunc("/registerauth", f.RegisterUserAuth)
	http.HandleFunc("/loginauth", f.LogInUserAuth)
	http.HandleFunc("/google-login", f.InitiateGoogleAuth)
	http.HandleFunc("/github-login", f.InitiateGithubAuth)
	http.HandleFunc("/google-callback", f.GoogleHandler)
	http.HandleFunc("/github-callback", f.GithubHandler)
	http.HandleFunc("/logout", f.LogOut)
	http.HandleFunc("/upload", f.Authorize(f.CreateUserPost))
	http.HandleFunc("/uploadComment", f.Authorize(f.CreateUserComment))
	http.HandleFunc("/likeOrDislike", f.HandleLikesOrDislikes)
	http.HandleFunc("/myActivity", f.Authorize(f.MyActivity))
	http.HandleFunc("/category", f.CategoryPage)
	f.StartServer()

}
