package forum

import (
	"log"
)

var Count *int

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// // A function to delete the current cookie session if the cookie table (in database) is empty
// // Shouldn't really happen normally, but pops up a lot because of how we're testing
// func isDBEmpty(w http.ResponseWriter, r *http.Request) {

// 	err = Database.QueryRow("SELECT COUNT() FROM Cookies").Scan(&Count)
// 	if err != nil {
// 		fmt.Println("Failed to execute query:", err)
// 	}

// 	// sets cookie max age to a negative, killing the cookie
// 	if *Count == 0 {
// 		cookie, err := r.Cookie("session")
// 		if err != nil {
// 			return
// 		}
// 		cookie.Value = ""         //sets cookie value to an empty string
// 		cookie.MaxAge = -1        //sets the existance of the cookie to -1
// 		http.SetCookie(w, cookie) //sets the modified values of the cookie
// 		fmt.Println("The table is empty.")
// 		http.Redirect(w, r, "/", http.StatusSeeOther)
// 	}
// }
