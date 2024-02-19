package utils

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
)

const (
	salt = "erogjnrgjjlsa2oqkjpo12j0i13ju491u3hrijwfjnf"
)

func HashPass(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func GenerateRefresh() string {
	result := make([]byte, 32)
	rand.Read(result)
	return fmt.Sprintf("%x", result)
}

func GenerateTokenForAccess() string {
	result := make([]byte, 30)
	rand.Read(result)
	return fmt.Sprintf("%x", result)
}
