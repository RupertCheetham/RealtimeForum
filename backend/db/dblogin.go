package db

import (
	"golang.org/x/crypto/bcrypt"
)

func GetLoginEntry(loginCheck UserEntry) map[string]string {
	dbLoginCheck, _ := GetUsersFromDatabase()

	message := map[string]string{
		"message": "",
	}

	for i := 0; i < len(dbLoginCheck); i++ {
		if dbLoginCheck[i].Username == loginCheck.Username {
			err := bcrypt.CompareHashAndPassword([]byte(dbLoginCheck[i].Password), []byte(loginCheck.Password))
			if err != nil {
				panic(err)
			}
			message["message"] = "Login successful"
			return message
		}
	}
	message["message"] = "Incorrect login details"
	return message
}
