package forum

import (
	"fmt"
	"log"
	"net/http"
)

// this function will get triggered when the URL path == /logout
func LogOut(w http.ResponseWriter, r *http.Request) {
	//request the cookie with the name session
	cookie, err := r.Cookie("session")

	query := "DELETE From Cookies WHERE SessionID = ?"

	stmt, err := Database.Prepare(query)
	if err != nil {
		log.Println("error preparing the query", query)
	}
	defer stmt.Close()

	result, err := stmt.Exec(cookie.Value)
	if err != nil {
		panic(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	if rowsAffected == 0 {
		fmt.Println("No rows affected")
	} else {
		//if cookie is found, delete the cookie and log the user out by stopping the session
		cookie.Value = ""         //sets cookie value to an empty string
		cookie.MaxAge = -1        //sets the existance of the cookie to -1
		http.SetCookie(w, cookie) //sets the modified values of the cookie
		fmt.Println("Session deleted")
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
