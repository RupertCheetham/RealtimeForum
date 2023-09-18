package db

// import (
// 	"database/sql"
// 	"fmt"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"github.com/gorilla/sessions"
// 	_ "github.com/mattn/go-sqlite3"
// )

// var store *sessions.CookiesStore
// var db *sql.DB

// func main() {
// 	// Initialize the session store with a secret key
// 	store = sessions.NewCookieStore([]byte("your-secret-key"))

// 	// Initialize the SQLite database connection (replace with your database setup)
// 	db, _ = sql.Open("sqlite3", "your-database-file.db")
// 	defer db.Close()

// 	// Initialize the router
// 	r := mux.NewRouter()

// 	// Define routes
// 	r.HandleFunc("/login", loginHandler).Methods("POST")
// 	r.HandleFunc("/dashboard", dashboardHandler).Methods("GET")

// 	// Start the HTTP server
// 	http.Handle("/", r)
// 	http.ListenAndServe(":8080", nil)
// }

// func loginHandler(w http.ResponseWriter, r *http.Request) {
// 	// Simulate user authentication (replace with your actual authentication logic)
// 	username := "john_doe"

// 	// Create a new session
// 	session, _ := store.Get(r, "session-name")

// 	// Set session options (expiration time)
// 	session.Options = &sessions.Options{
// 		Path:     "/",
// 		MaxAge:   3600, // Session expires in 1 hour (adjust as needed)
// 		HttpOnly: true,
// 	}

// 	// Store the user's data in the session (in this case, just the username)
// 	session.Values["username"] = username

// 	// Save the session
// 	session.Save(r, w)

// 	// Redirect the user to the dashboard
// 	http.Redirect(w, r, "/dashboard", http.StatusFound)
// }

// func dashboardHandler(w http.ResponseWriter, r *http.Request) {
// 	// Retrieve the user's session
// 	session, _ := store.Get(r, "session-name")

// 	// Check if the user is authenticated (replace with your authentication logic)
// 	if username, ok := session.Values["username"].(string); ok {
// 		// User is authenticated; you can use 'username' to retrieve associated data from your database
// 		// For demonstration purposes, we'll simply display the username
// 		fmt.Fprintf(w, "Welcome, %s!", username)
// 	} else {
// 		// User is not authenticated; you can redirect them to the login page
// 		http.Redirect(w, r, "/login", http.StatusFound)
// 	}
// }
