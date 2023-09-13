package db

import "log"

func AddUserToDatabase(username string, age int, gender string, firstName string, lastName string, email string, password string) error {
	_, err := Database.Exec("INSERT INTO USERS (Username, Age, Gender, First_name, Last_name, Email, Password) VALUES (?, ?, ?, ?, ?, ?, ?)", username, age, gender, firstName, lastName, email, password)
	if err != nil {
		log.Println("Error adding USER to database:", err)
	}
	return err
}

func GetUsersFromDatabase() ([]UserEntry, error) {
	rows, err := Database.Query("SELECT Username, Age, Gender, First_name, Last_name, Email, Password FROM User ORDER BY Id ASC")
	if err != nil {
		log.Println("Error querying USERS from database:", err)
		return nil, err
	}
	defer rows.Close()

	var users []UserEntry
	for rows.Next() {
		var entry UserEntry
		err := rows.Scan(&entry.Username, &entry.Age, &entry.Gender, &entry.FirstName, &entry.LastName, &entry.Email, &entry.Password)
		if err != nil {
			log.Println("Error scanning row from database:", err)
			return nil, err
		}
		users = append(users, entry)
	}

	return users, nil
}

// // make a handler function to authenticate registration of users
// func RegisterUserAuth(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Redirect(w, r, "register.html", http.StatusSeeOther)
// 		return
// 	}
// 	//parse the form
// 	err := r.ParseForm()
// 	HandleError(err)

// 	//get user info from the form
// 	userName := r.FormValue("user_name")
// 	userEmail := r.FormValue("user_email")
// 	userPassword := r.FormValue("user_password")
// 	pass_confirm := r.FormValue("confirm_password")
// 	fmt.Println("username: ", userName)
// 	fmt.Println("useremai: ", userEmail)
// 	fmt.Println("userpass: ", userPassword)
// 	fmt.Println("pass confirm: ", pass_confirm)

// 	//check if the username/email exists in our database
// 	emailExists, err := UserExists("", userEmail)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	usernameExists, err := UserExists(userName, "")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	if !ValidUserName(userName) {
// 		tmpl.ExecuteTemplate(w, "register.html", ErrorMessage{UsernameError: "Invalid username"})
// 		return
// 	} else if !ValidEmail(userEmail) {
// 		tmpl.ExecuteTemplate(w, "register.html", ErrorMessage{EmailError: "Please enter a valid email address"})
// 		return
// 	} else if !ValidPassword(userPassword) {
// 		tmpl.ExecuteTemplate(w, "register.html", ErrorMessage{PasswordError: "Invalid Password"})
// 		return
// 		//only checking if the email exists in our database
// 	} else if emailExists {
// 		tmpl.ExecuteTemplate(w, "register.html", ErrorMessage{UsernameError: "Email already taken"})
// 		return
// 		//only checking if the username exists in our database
// 	} else if usernameExists {
// 		tmpl.ExecuteTemplate(w, "register.html", ErrorMessage{UsernameError: "Username is already taken"})
// 		return
// 	} else if userPassword != pass_confirm {
// 		tmpl.ExecuteTemplate(w, "register.html", ErrorMessage{ConfirmError: "Passwords do not match"})
// 		return
// 	}
// 	RegisterUser(userName, userEmail, userPassword)
// 	userID, err := GetUserIDfrom("users", "username", userName)
// 	fmt.Println("userid: ", userID)
// 	if err != nil {
// 		fmt.Println("there was an error: ", err)
// 	}
// 	SetUserSession(w, r, userID)
// 	http.Redirect(w, r, "/", http.StatusSeeOther)
// }

// // make a function to insert the user in you database
// func RegisterUser(username, useremail, userpass string) {
// 	var query = "INSERT INTO users(username, email, password) VALUES(?,?,?)"
// 	var hash []byte
// 	hash, err = bcrypt.GenerateFromPassword([]byte(userpass), bcrypt.DefaultCost) //generates hash for user password
// 	stmt, err := Database.Prepare(query)
// 	HandleError(err)

// 	defer stmt.Close()
// 	stmt.Exec(username, useremail, hash)
// }
