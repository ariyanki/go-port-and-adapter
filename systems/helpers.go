package systems

import (
	"crypto/sha256"
	"encoding/hex"
)

func GeneratePassword(username, password, salt string) string {
	hasher := sha256.New()
	_, err := hasher.Write([]byte(username + password + salt))
    if err != nil {
        return ""
    }
	hashBytes := hasher.Sum(nil)
	hashHex := hex.EncodeToString(hashBytes)
	return hashHex
}