package helpers

import (
	"crypto/rand"
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
