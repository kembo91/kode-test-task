package utils

import (
	"fmt"
	"regexp"
)

var isUsernameStmt = regexp.MustCompile(`^[A-Za-z0-9_\-.]+$`).MatchString
var isAnagramStmt = regexp.MustCompile(`^[a-z]+$`).MatchString

//IsValidUsername checks if username is valid
func IsValidUsername(s string) error {
	if len(s) < 5 {
		return fmt.Errorf("username must be at least 5 characters long")
	}
	if !isUsernameStmt(s) {
		return fmt.Errorf(`username can contain only alphabetic letters, numbers, or _ - . symbols`)
	}
	return nil
}

//IsValidAnagram checks if an incoming anagram consists of only characters
func IsValidAnagram(s string) error {
	if !isAnagramStmt(s) {
		return fmt.Errorf("anagram must contain only alphabetical characters")
	}
	return nil
}

//IsValidPassword checks if a password is valid
func IsValidPassword(s string) error {
	if len(s) < 8 {
		return fmt.Errorf("password must contain 8 or more characters")
	}
	return nil
}
