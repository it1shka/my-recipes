package myutils

import (
	"strings"
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

func IsPasswordStrong(pwd string) (valid bool, message string) {
	var hasNumber, hasUpper, hasSpecial bool

	for _, char := range pwd {
		switch {
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		case !(unicode.IsLetter(char) || char == ' '):
			valid = false
			message = "unexpected character encountered"
			return
		}
	}

	problems := []string{}
	if !hasNumber {
		problems = append(problems, "you should add a number")
	}
	if !hasUpper {
		problems = append(problems, "you should add an uppercase letter")
	}
	if !hasSpecial {
		problems = append(problems, "you should add a special character")
	}
	if len([]rune(pwd)) < 8 {
		problems = append(problems, "your password should be at least 8 characters long")
	}

	valid = (len(problems) == 0)
	message = strings.Join(problems, ", ")

	return
}
