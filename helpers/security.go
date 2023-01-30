package helpers

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"goshaka/configs"
	"io"

	"golang.org/x/crypto/bcrypt"
)

func CompareHash(hashed string, plain string) bool {
	errHash := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return errHash == nil
}

func CreateHash(text string) (string, error) {
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedToken), nil
}

func PadKey(key []byte, blockSize int) []byte {
	padding := blockSize - len(key)%blockSize
	pad := bytes.Repeat([]byte{0}, padding)
	return append(key, pad...)
}

func EncryptText(originText string) (string, error) {
	key := PadKey([]byte(configs.GetEnv("APP_KEY")), aes.BlockSize)
	plaintext := []byte(originText)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func DecryptText(encryptedText string) (string, error) {
	decodedString, _ := base64.StdEncoding.DecodeString(encryptedText)
	ciphertext := decodedString
	key := PadKey([]byte(configs.GetEnv("APP_KEY")), aes.BlockSize)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext is too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}
