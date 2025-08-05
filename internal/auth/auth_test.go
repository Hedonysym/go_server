package auth

import (
	"fmt"
	"testing"
)

func TestAuth1(t *testing.T) {
	password := "fuckass"
	hash, err := HashPassword(password)
	if err != nil {
		fmt.Errorf("hashing failed: %v", err)
	}
	err = CheckPasswordHash(password, hash)
	if err != nil {
		fmt.Errorf("checking failed: %v", err)
	}
}

func TestAuth2(t *testing.T) {
	password := "fuckass"
	hash, err := HashPassword(password)
	if err != nil {
		fmt.Errorf("hashing failed: %v", err)
	}
	err = CheckPasswordHash("wrong password", hash)
	if err == nil {
		fmt.Errorf("checking failed: %v", err)
	}
}

func TestAuth3(t *testing.T) {
	password := "this will be an entirely too long password, this is so long ang shit theres no way its valid. holy shit its long and invalid, thats fcking wild bro frfr nikka 6969696969696969696969696969696969696969696996969696969"
	hash, err := HashPassword(password)
	if err == nil {
		fmt.Errorf("hashing failed: %v", err)
	}
	err = CheckPasswordHash(password, hash)
	if err != nil {
		fmt.Errorf("checking failed: %v", err)
	}
}
