package auth

import (
	"crypto/rand"
	"encoding/hex"
)

func MakeRefreshToken() (string, error) {
	bytes := []byte{}
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	str := hex.EncodeToString(bytes)
	return str, nil
}
