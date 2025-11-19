package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
)

var passwordSecret = os.Getenv("PASSWORD_SECRET")

func HashPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password + passwordSecret))
	return hex.EncodeToString(h.Sum(nil))
}
