package db

import "fmt"

func GetLoginEntry(loginCheck RegistrationEntry) {
	dbLoginCheck, _ := GetRegistrationFromDatabase()

	for i := 0; i < len(dbLoginCheck); i++ {
		if dbLoginCheck[i].Username == loginCheck.Username && dbLoginCheck[i].Password == loginCheck.Password {
			fmt.Println("blimey!")
			return
		}
	}
	fmt.Println("Oh no")
}
