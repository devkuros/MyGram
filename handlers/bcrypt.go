package handlers

import "golang.org/x/crypto/bcrypt"

func HashPassword(p string) string {
	salt := 8
	password := []byte(p)
	hash, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hash)
}

func ComparePassword(h, p []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(p))

	return err == nil
}
