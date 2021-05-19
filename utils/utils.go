package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

func HashSecret(secret string) string {
	h := sha256.Sum256([]byte(secret))
	return base64.StdEncoding.EncodeToString(h[:])
}
