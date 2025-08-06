package auth

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
)

func JWTTest1(t *testing.T) {
	ogId := uuid.New()
	token, err := MakeJWT(ogId, "mah balls", time.Duration(100000))
	if err != nil {
		fmt.Errorf("make token failed: %v", err)
	}
	newId, err := ValidateJWT(token, "mah balls")
	if err != nil {
		fmt.Errorf("validate token failed: %v", err)
	}
	if ogId != newId {
		fmt.Errorf("id mismatch: %v, %v", ogId, newId)
	}
}

func JWTTest2(t *testing.T) {
	ogId := uuid.New()
	token, err := MakeJWT(ogId, "mah balls", time.Duration(-500))
	if err != nil {
		fmt.Errorf("make token failed: %v", err)
	}
	newId, err := ValidateJWT(token, "mah balls")
	if err == nil {
		fmt.Errorf("validate token failed: %v", err)
	}
	if ogId != newId {
		fmt.Errorf("id mismatch: %v, %v", ogId, newId)
	}
}

func JWTTest3(t *testing.T) {
	ogId := uuid.New()
	token, err := MakeJWT(ogId, "mah balls", time.Duration(100000))
	if err != nil {
		fmt.Errorf("make token failed: %v", err)
	}
	newId, err := ValidateJWT(token, "mah dick")
	if err == nil {
		fmt.Errorf("validate token failed: %v", err)
	}
	if ogId != newId {
		fmt.Errorf("id mismatch: %v, %v", ogId, newId)
	}
}
