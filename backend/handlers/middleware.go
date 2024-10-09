package handlers

import (
	"net/http"
)

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
