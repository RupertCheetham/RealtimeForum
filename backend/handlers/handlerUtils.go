package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"realtimeForum/db"
	"strings"
	"time"
)

const CookieName = "sessionID"

var currentTime = time.Now()

// The function sets up Cross-Origin Resource Sharing (CORS) headers for an HTTP response.
func SetupCORS(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "https://localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Methods", "HEAD, POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Requested-With")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")

	// Respond to preflight request with 200 OK status
	if r.Method == "OPTIONS" {
		(*w).WriteHeader(http.StatusOK)
		return
	}
}

func CookieCheck(successHandler, failureHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		SetupCORS(&w, r)
		cookieValue := GetCookie(w, r)
		if len(cookieValue) == 0 {
			failureHandler(w, r)
		} else {
			successHandler(w, r)
		}
	}
}

func GetCookie(w http.ResponseWriter, r *http.Request) string {
	SetupCORS(&w, r)
	sessionCookie, _ := r.Cookie(CookieName)
	fmt.Println("sessioncookie in getCookie", sessionCookie)
	splitCookieValue := strings.Split(sessionCookie.String(), "=")
	var cookieValue string

	if len(splitCookieValue) == 2 {
		cookieValue = splitCookieValue[1]
	}

	return cookieValue
}

func ActionSuccessMessage(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)
	msg := map[string]string{
		"message": "successful request",
	}

	jsonResponse, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

func ActionFailedMessage(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)
	msg := map[string]string{
		"message": "unsuccessful request",
	}

	jsonResponse, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusRequestTimeout)
	w.Write(jsonResponse)
}

func CheckCookieExpiration(w http.ResponseWriter, r *http.Request, sessionToken *db.Session) {
	if currentTime.Before(sessionToken.ExpirationTime) {
		log.Println("Received post:")

		// stuff should happen here
		// err = db.AddPostToDatabase(post.UserId, post.Img, post.Body, post.Categories)
		// if err != nil {
		// 	utils.HandleError("Problem adding to POSTS in AddPostHandler", err)
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		ActionSuccessMessage(w, r)
	} else {

		ActionFailedMessage(w, r)
	}
}

// func CheckCookieTimeout(w http.ResponseWriter, r *http.Request, cookieValue string) {
// 	if r.Method == "POST" {
// 		CheckCookieExists(w, r, cookieValue)

// 		var post db.PostEntry

// 		err := json.NewDecoder(r.Body).Decode(&post)
// 		if err != nil {
// 			utils.HandleError("Problem decoding JSON in AddPostHandler", err)
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}

// 		sessionToken, _ := db.GetSessionByToken(cookieValue)
// 		if currentTime.Before(sessionToken.ExpirationTime) {
// 			log.Println("Received post:", post.UserId, post.Img, post.Body, post.Categories)

// 			err = db.AddPostToDatabase(post.UserId, post.Img, post.Body, post.Categories)
// 			if err != nil {
// 				utils.HandleError("Problem adding to POSTS in AddPostHandler", err)
// 				http.Error(w, err.Error(), http.StatusInternalServerError)
// 				return
// 			}

// 			ActionSuccessMessage(w, r)
// 		} else {

// 			ActionFailedMessage(w, r)
// 		}
// 		fmt.Println("chilling outside of 408 status response")
// 	}
// }
