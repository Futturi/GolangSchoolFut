package utils

import (
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
