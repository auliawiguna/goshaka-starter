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
	"sync"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/crypto/bcrypt"
)

var (
	RateLimiter = make(map[string]int64)
	PointLimit  = make(map[string]int64)
	Mutex       sync.Mutex
)

// Throttle function/request by key
//
//	param key string
//	param maxAttempt int
//	param maxTimeInSeconds int
//	return bool
func RateLimit(key string, maxAttempt int, maxTimeInSeconds int) bool {
	Mutex.Lock()
	defer Mutex.Unlock()

	now := time.Now().Unix()

	if v, ok := RateLimiter[key]; ok && now-v < int64(maxTimeInSeconds) {
		RateLimiter[key]++
		PointLimit[key]++
		if PointLimit[key] > int64(maxAttempt) {
			return false
		}
	} else {
		RateLimiter[key] = now
		PointLimit[key] = 1
	}

	return true
}

// Compare hash and a plain text, will returns true if hash is match
//
//	param	hashed	string
//	param	plain	string
//	return	bool
func CompareHash(hashed string, plain string) bool {
	errHash := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return errHash == nil
}

// Convert plain text to hash string
//
//	param	text	string
//	return	string, error
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

// Convert plain text to encrypted string
//
//	param	originText	string
//	return	string, error
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

// Convert decrypted text to plain text
//
//	param	encryptedText	string
//	return	string, error
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

// sanitise text
//
//	param	str	string
//	return	string
func SanitiseText(str string) string {
	sanitise := bluemonday.UGCPolicy()
	return sanitise.Sanitize(str)
}
