package utils

import "golang.org/x/crypto/bcrypt"

func HashString(str string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), 14)
	return string(bytes), err
}
