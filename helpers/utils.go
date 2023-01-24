package helpers

import (
	"crypto/rand"

	"golang.org/x/crypto/bcrypt"
)

func RandomNumber(numLength int) string {
	token := make([]byte, numLength)
	_, err := rand.Read(token)
	if err != nil {
		return ""
	}

	for i := 0; i < len(token); i++ {
		token[i] = token[i]%10 + '0'
	}
	return string(token)
}

func CompareHash(hashed string, plain string) bool {
	errHash := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return errHash == nil
}
