package forum

import (
	"net/http"
)

// Decides which page to display to user based on the URL Path
func ServePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		roothandler(w, r)
	} else if r.URL.Path == "/register.html" {
		tmpl.ExecuteTemplate(w, "register.html", nil)
	} else if r.URL.Path == "/sign-in.html" {
		tmpl.ExecuteTemplate(w, "sign-in.html", nil)
	} else {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}
}
