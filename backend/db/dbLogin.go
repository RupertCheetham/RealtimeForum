package db

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func GetLoginEntry(loginCheck UserEntry) (message map[string]string, id int, err error) {
	id = 0
	message = map[string]string{
		"message": "",
	}

	dbLoginCheck, err := FindUserFromDatabase(loginCheck.Username)

	fmt.Println("dbLoginCheck:", dbLoginCheck)

	if err != nil {
		return message, id, err
	}

	// Check if the input contains an '@' sign
	isEmail := strings.Contains(loginCheck.Username, "@")

	if (isEmail && dbLoginCheck.Email == loginCheck.Email) || (!isEmail && dbLoginCheck.Username == loginCheck.Username) {
		err := bcrypt.CompareHashAndPassword([]byte(dbLoginCheck.Password), []byte(loginCheck.Password))
		if err != nil {
			message["message"] = "Incorrect username or password"
		} else {
			id = dbLoginCheck.Id
			message["message"] = "Login successfulley"
		}
	}
	return message, id, err
}
