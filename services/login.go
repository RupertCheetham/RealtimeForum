package forum

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

// HANDLERS FOR USER LOG IN
func LogInUserAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "sign-in.html", http.StatusSeeOther)
		return
	}
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	//get user info
	userName := r.FormValue("user_name")
	userPassword := r.FormValue("user_password")
	//get userID from db
	userID, err := GetUserIDfrom("users", "username", userName)

	if err != nil {
		fmt.Println("there was an error: ", err)
	}
	//get hash password from DB
	hash, err := getUserHash(userName)
	fmt.Println("hash: ", hash)
	if err != nil {
		fmt.Println("there was another error:", err)
	}
	//get username from DB if exists
	userExists, err := UserExists(userName, "")
	fmt.Println("user exists: ", userExists)
	if err != nil {
		fmt.Println(err)
	}
	//if user with given username doesn't exists display error
	if !userExists {
		tmpl.ExecuteTemplate(w, "sign-in.html", ErrorMessage{SignInError: "The username you entered isn't connected to an account."})
		return
	}
	//comapare the hash from DB to the input user password
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(userPassword))
	if err != nil {
		tmpl.ExecuteTemplate(w, "sign-in.html", ErrorMessage{PasswordError: "Incorrect password"})
		return
	}

	//set user session
	SetUserSession(w, r, userID)
	isloggedin, _, _ := loggedIn(w, r)
	fmt.Println("isloggedin in login", isloggedin)
	// Checks that the user only has the 1 session
	extraSessionIDs := UserAlreadyLoggedIn(userID)
	ExtraSessionRemover(extraSessionIDs)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// sets a cookie and creates a session for the user when the user log's in
func SetUserSession(w http.ResponseWriter, r *http.Request, userID string) {
	var sessionID uuid.UUID
	//check if there is a cookie for the user
	cookie, err := r.Cookie("session")
	//if there isn't, create one
	if err == http.ErrNoCookie {
		//generate a UUID for each cookie session
		sessionID, err = uuid.NewV4()
		if err != nil {
			log.Fatalf("failed to generate UUID: %v", err)
		} //create a cookie
		cookie = &http.Cookie{
			Name:     "session",
			Value:    sessionID.String(),
			HttpOnly: true,                    //indicates that the cookie is accessible by the HTTP protocol
			Secure:   true,                    //indicates that the cookie will only be sent over a secure conneection
			MaxAge:   86400 * 7,               //indicates expiration time of the cookie
			Path:     "/",                     //indicates the path that the cookies is valid for
			SameSite: http.SameSiteStrictMode, //restricts the cookie only for this site
		}
		//once the cookie is created insert it to the cookies table with the corresponding userID
		query := "INSERT INTO cookies(sessionid, userid) VALUES (?,?)"
		stmt, err := Database.Prepare(query)
		if err != nil {
			log.Println("Error preparing the query: ", query)
		}
		defer stmt.Close()

		_, err = stmt.Exec(sessionID.String(), userID)
		if err != nil {
			log.Println("Error executing the stmt: ", stmt)
		}
	}

	http.SetCookie(w, cookie)

}

/*
returns if the user is logged in by requesting a cookie with the given name
returns the cookie value and an error if there is one
*/
func loggedIn(w http.ResponseWriter, r *http.Request) (bool, string, error) {
	cookie, err := r.Cookie("session")
	if err != nil {
		//log.Println("Error retrieving session cookie:", err)
		return false, "", err
	}
	return true, cookie.Value, nil
}

// use this function when you need to authorize the user
func Authorize(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loggedIn, _, _ := loggedIn(w, r)
		if !loggedIn {
			http.Redirect(w, r, "/sign-in.html", http.StatusSeeOther)
			return
		}
		h.ServeHTTP(w, r)
	}
}
