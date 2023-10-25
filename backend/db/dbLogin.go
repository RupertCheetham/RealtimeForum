package db

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func GetLoginEntry(loginCheck UserEntry) (message map[string]string, id int, err error) {
	id = 0
	message = map[string]string{
		"message": "",
	}

	dbLoginCheck, err := FindUserFromDatabase(loginCheck.Username)

	if err != nil {
		return message, 0, err
	}

	// Check if the input contains an '@' sign
	isEmail := strings.Contains(loginCheck.Username, "@")

	if (isEmail && dbLoginCheck.Email == loginCheck.Email) || (!isEmail && dbLoginCheck.Username == loginCheck.Username) {
		err := bcrypt.CompareHashAndPassword([]byte(dbLoginCheck.Password), []byte(loginCheck.Password))
		if err != nil {
			message["message"] = "Incorrect username or password"
			return message, id, err
		}
		id = dbLoginCheck.Id
		message["message"] = "Login successfully"
	}

	return message, id, err
}
