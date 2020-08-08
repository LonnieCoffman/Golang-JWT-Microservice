package validation

import (
	"regexp"
	"unicode"
)

// PasswordValid validates that all password requirements have been met
func PasswordValid(pass string) bool {
	var (
		hasMinLen, hasUpper, hasLower, hasNumber, hasSpecial = false, false, false, false, false
	)
	if len(pass) >= 7 || len(pass) > 255 {
		hasMinLen = true
	}
	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

// EmailValid validates that an email address is in the correct format and is less than 254 characters in length
func EmailValid(email string) bool {
	var validEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(email) > 254 || !validEmail.MatchString(email) {
		return false
	}
	return true
}
