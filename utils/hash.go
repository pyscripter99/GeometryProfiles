package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

func HashName(name string) string {
	profileHash := sha256.Sum256([]byte(name))
	return base64.URLEncoding.EncodeToString(profileHash[:])
}
