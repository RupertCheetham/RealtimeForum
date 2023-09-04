package forum

import (
	"fmt"
	"regexp"
)

// make a function to validate nickname
func ValidUserNickName(nickname string) bool {
	var nicknameRegex = regexp.MustCompile("^[a-zA-Z0-9]{5,50}")
	fmt.Println("username validation: ", nicknameRegex.MatchString(nickname))
	return nicknameRegex.MatchString(nickname)
}

// make a function to validate user password
func ValidPassword(userpass string) bool {
	var passRegex = regexp.MustCompile("[A-Za-z0-9!@#$%^&*(),.?:{}|<>]{8,50}")
	fmt.Println("pass validation: ", passRegex.MatchString(userpass))
	return passRegex.MatchString(userpass)
}

// make a function to validate user email
func ValidEmail(useremail string) bool {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	fmt.Println("mail validation: ", emailRegex.MatchString(useremail))
	return emailRegex.MatchString(useremail)
}

// make a function to validate user First or Last name
func ValidName(firstOrLastName string) bool {
	var firstOrLastNameRegex = regexp.MustCompile("^[a-zA-Z0-9] {1,}")
	fmt.Println("mail validation: ", firstOrLastNameRegex.MatchString(firstOrLastName))
	return firstOrLastNameRegex.MatchString(firstOrLastName)
}
