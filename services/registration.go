package forum

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

//HANDLERS FOR USER REGISTERATION

// make a handler function to authenticate registretion of users
func RegisterUserAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "register.html", http.StatusSeeOther)
		return
	}
	//parse the form
	err := r.ParseForm()
	HandleError(err)

	//get user info from the form
	userName := r.FormValue("user_name")
	userEmail := r.FormValue("user_email")
	userPassword := r.FormValue("user_password")
	pass_confirm := r.FormValue("confirm_password")
	fmt.Println("username: ", userName)
	fmt.Println("useremai: ", userEmail)
	fmt.Println("userpass: ", userPassword)
	fmt.Println("pass confirm: ", pass_confirm)

	//check if the username/email exists in our database
	emailExists, err := UserExists("", userEmail)
	if err != nil {
		fmt.Println(err)
	}
	usernameExists, err := UserExists(userName, "")
	if err != nil {
		fmt.Println(err)
	}

	if !ValidUserName(userName) {
		tmpl.ExecuteTemplate(w, "register.html", ErrorMessage{UsernameError: "Invalid username"})
		return
	} else if !ValidEmail(userEmail) {
		tmpl.ExecuteTemplate(w, "register.html", ErrorMessage{EmailError: "Please enter a valid email address"})
		return
	} else if !ValidPassword(userPassword) {
		tmpl.ExecuteTemplate(w, "register.html", ErrorMessage{PasswordError: "Invalid Password"})
		return
		//only checking if the email exists in our database
	} else if emailExists {
		tmpl.ExecuteTemplate(w, "register.html", ErrorMessage{UsernameError: "Email already taken"})
		return
		//only checking if the username exists in our database
	} else if usernameExists {
		tmpl.ExecuteTemplate(w, "register.html", ErrorMessage{UsernameError: "Username is already taken"})
		return
	} else if userPassword != pass_confirm {
		tmpl.ExecuteTemplate(w, "register.html", ErrorMessage{ConfirmError: "Passwords do not match"})
		return
	}
	RegisterUser(userName, userEmail, userPassword)
	userID, err := GetUserIDfrom("users", "username", userName)
	fmt.Println("userid: ", userID)
	if err != nil {
		fmt.Println("there was an error: ", err)
	}
	SetUserSession(w, r, userID)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// insert the user in to database
func RegisterUser(username, useremail, userpass string) {
	var query = "INSERT INTO users(username, email, password) VALUES(?,?,?)"
	var hash []byte
	hash, err = bcrypt.GenerateFromPassword([]byte(userpass), bcrypt.DefaultCost) //generates hash for user password
	stmt, err := Database.Prepare(query)
	HandleError(err)

	defer stmt.Close()
	stmt.Exec(username, useremail, hash)
}
