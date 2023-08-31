package forum

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
)

func configureGoogleOAuth2Client() *oauth2.Config {
	config := &oauth2.Config{
		ClientID:     "623060904480-iudd85gmtab7k4sr1utdo92e5lpfsuqd.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "https://localhost:7000/google-callback",
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}
	return config
}

// takes user to google login page then redirects to permitted redirect
func InitiateGoogleAuth(w http.ResponseWriter, r *http.Request) {
	config := configureGoogleOAuth2Client()
	url := config.AuthCodeURL("state-token", oauth2.AccessTypeOnline)
	http.Redirect(w, r, url, http.StatusFound)
}

func configureGithubOAuth2Client() *oauth2.Config {
	githubEndpoint := oauth2.Endpoint{
		AuthURL:  "https://github.com/login/oauth/authorize",
		TokenURL: "https://github.com/login/oauth/access_token",
	}

	config := &oauth2.Config{
		ClientID:     "4764bca26a3fa9ab0e32",
		ClientSecret: "4b97506ea58726f63c065ce8ed60b35c1fcbf532",
		RedirectURL:  "https://localhost:7000/github-callback",
		Scopes:       []string{"user:email"}, // Specify desired scopes, including user:email
		Endpoint:     githubEndpoint,
	}
	return config
}

// takes user to github login page then redirects to permitted redirect
func InitiateGithubAuth(w http.ResponseWriter, r *http.Request) {
	config := configureGithubOAuth2Client()
	url := config.AuthCodeURL("state-token")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// gets userID (big consistent number), name and email from google
func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) string {
	config := configureGoogleOAuth2Client()
	code := r.URL.Query().Get("code")
	token, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Println("There's a problem with token in HandleGoogleCallback")
		// Handle error
	}

	idToken, err := idtoken.Validate(context.Background(), token.Extra("id_token").(string), "623060904480-iudd85gmtab7k4sr1utdo92e5lpfsuqd.apps.googleusercontent.com")
	if err != nil {
		log.Println("There's a problem with idToken in HandleGoogleCallback")
		// Handle error
	}

	// Access the claims and fields of the ID token

	email := idToken.Claims["email"].(string)
	username := idToken.Claims["name"].(string)

	userStruct := Authentication{
		Email: email,
		Name:  username,
	}

	userID := dealWithUser(w, r, userStruct)
	return userID

}

func handleGitHubCallback(w http.ResponseWriter, r *http.Request) string {
	config := configureGithubOAuth2Client()
	// state := r.FormValue("state")
	// if state != oauthStateString {
	// 	// Invalid state parameter, handle the error
	// 	return
	// }
	code := r.FormValue("code")
	token, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Println("There was a problem with token exchange in handleGitHubCallback:", err)
	}

	// Create a GitHub client using the access token
	client := github.NewClient(oauth2.NewClient(context.Background(), config.TokenSource(context.Background(), token)))
	// Get the authenticated user's information
	user, _, err := client.Users.Get(context.Background(), "")
	if err != nil {
		log.Println("There was a problem with client.Users in handleGitHubCallback:", err)
	}

	if user == nil {
		log.Println("The user is nil in handleGitHubCallback")
	}

	// Use the user information as needed
	username := *user.Login
	var email string
	if user.Email != nil {
		email = *user.Email
	} else {
		email = username
	}

	userStruct := Authentication{
		Email: email,
		Name:  username,
	}
	userID := dealWithUser(w, r, userStruct)
	return userID

}

func ExternalAuthHandler(w http.ResponseWriter, r *http.Request, authoriser string) {
	var userID string

	if authoriser == "Google" {
		// gets google UserID, name and email from google
		userID = HandleGoogleCallback(w, r)
	}
	if authoriser == "Github" {
		// gets google UserID, name and email from google
		userID = handleGitHubCallback(w, r)
	}
	extraSessionIDs := UserAlreadyLoggedIn(userID)
	ExtraSessionRemover(extraSessionIDs)
	var submittedStruct FinalStruct
	// gets the posts from the database
	submittedStruct.Posts = getPostsFromDB()
	username, _ := getUsernameFromUserID(userID)
	submittedStruct.LoggedInUsername = username
	tmpl.ExecuteTemplate(w, "index.html", submittedStruct)
	// http.Redirect(w, r, "/", http.StatusSeeOther)
}

func dealWithUser(w http.ResponseWriter, r *http.Request, UserStruct Authentication) string {
	var userID string
	// check if the username/email exists in our database
	emailExists, err := UserExists("", UserStruct.Email)
	if err != nil {
		fmt.Println(err)
	}
	usernameExists, err := UserExists(UserStruct.Name, "")
	if err != nil {
		fmt.Println(err)
	}

	if !emailExists && !usernameExists {
		RegisterUser(UserStruct.Name, UserStruct.Email, "")
		userID, err = GetUserIDfrom("users", "username", UserStruct.Name)
		if err != nil {
			log.Println("there was a problem with finding UserID from googlename in dealWithUser: ", err)
		}
	} else if emailExists && !usernameExists {
		userID, err = GetUserIDfrom("users", "email", UserStruct.Email)
		if err != nil {
			log.Println("there was a problem with finding UserID from gmail in dealWithUser: ", err)
		}
	} else if emailExists && usernameExists {
		userID, err = GetUserIDfrom("users", "email", UserStruct.Email)
		if err != nil {
			log.Println("there was a problem with finding UserID from gmail in dealWithUser: ", err)
		}
	}

	fmt.Println("userid: ", userID)
	if err != nil {
		log.Println("there was an error in dealWithUser: ", err)
	}
	SetUserSession(w, r, userID)
	isloggedin, _, _ := loggedIn(w, r)
	fmt.Println("isloggedin in external", isloggedin)
	return userID
}

func GoogleHandler(w http.ResponseWriter, r *http.Request) {
	ExternalAuthHandler(w, r, "Google")
}

func GithubHandler(w http.ResponseWriter, r *http.Request) {
	ExternalAuthHandler(w, r, "Github")
}
