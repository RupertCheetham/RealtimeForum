package main

import (
	"fmt"
	"regexp"
)

// make a function to validate nickname
func ValidUsername(username string) bool {
	var usernameRegex = regexp.MustCompile("^[a-zA-Z0-9]{5,50}")
	fmt.Println("username validation: ", usernameRegex.MatchString(username))
	return usernameRegex.MatchString(username)
}

// make a funciton to validate first and last names
func ValidName(userFirstOrLastName string) bool {
	var nameRegex = regexp.MustCompile("^[a-zA-Z- ]{1,}")
	fmt.Println("first or last name validation: ", nameRegex.MatchString(userFirstOrLastName))
	return nameRegex.MatchString(userFirstOrLastName)
}

// make a function to validate user email
func ValidEmail(userEmail string) bool {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	fmt.Println("email validation: ", emailRegex.MatchString(userEmail))
	return emailRegex.MatchString(userEmail)
}

// make a function to validate user password
func ValidPassword(userPassword string) bool {
	var passwordRegex = regexp.MustCompile("[A-Za-z0-9!@#$%^&*(),.?:{}|<>]{8,50}")
	fmt.Println("password validation: ", passwordRegex.MatchString(userPassword))
	return passwordRegex.MatchString(userPassword)
}
