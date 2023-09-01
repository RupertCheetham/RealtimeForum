package forum

import (
	"fmt"
	"regexp"
)

// make a function to validate username
func ValidUserName(username string) bool {
	var usernameRegex = regexp.MustCompile("^[a-zA-Z0-9]{5,50}")
	fmt.Println("username validation: ", usernameRegex.MatchString(username))
	return usernameRegex.MatchString(username)
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
