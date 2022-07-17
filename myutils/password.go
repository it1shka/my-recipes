package myutils

import (
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	pwd := []byte(password)
	bytes, err := bcrypt.GenerateFromPassword(pwd, 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	pwd, hpwd := []byte(password), []byte(hash)
	err := bcrypt.CompareHashAndPassword(hpwd, pwd)
	return err == nil
}

func IsPasswordStrong(pwd string) bool {
	var length, number, upper, special bool
	letters := 0
	for _, char := range pwd {
		switch {
		case unicode.IsNumber(char):
			number = true
		case unicode.IsUpper(char):
			upper = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			special = true
		case unicode.IsLetter(char) || char == ' ':
			letters++
		default:
			return false
		}
	}
	length = letters > 7
	return length && number && upper && special
}
