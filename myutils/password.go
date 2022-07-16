package myutils

import "golang.org/x/crypto/bcrypt"

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
