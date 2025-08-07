package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func JWTTest1(t *testing.T) {
	pass, err := HashPassword("mah balls")
	if err != nil {
		t.Errorf("hash error: %v", err)
	}
	ogId := uuid.New()
	token, err := MakeJWT(ogId, pass, time.Duration(100000))
	if err != nil {
		t.Errorf("make token failed: %v", err)
	}
	newId, err := ValidateJWT(token, pass)
	if err != nil {
		t.Errorf("validate token failed: %v", err)
	}
	if ogId != newId {
		t.Errorf("id mismatch: %v, %v", ogId, newId)
	}
}

func JWTTest2(t *testing.T) {
	pass, err := HashPassword("mah balls")
	if err != nil {
		t.Errorf("hash error: %v", err)
	}
	ogId := uuid.New()
	token, err := MakeJWT(ogId, pass, time.Duration(-500))
	if err != nil {
		t.Errorf("make token failed: %v", err)
	}
	newId, err := ValidateJWT(token, pass)
	if err == nil {
		t.Errorf("validate token failed: %v", err)
	}
	if ogId != newId {
		t.Errorf("id mismatch: %v, %v", ogId, newId)
	}
}

func JWTTest3(t *testing.T) {
	pass, err := HashPassword("mah balls")
	if err != nil {
		t.Errorf("hash error: %v", err)
	}
	pass2, err := HashPassword("mah dick")
	if err != nil {
		t.Errorf("hash error: %v", err)
	}
	ogId := uuid.New()
	token, err := MakeJWT(ogId, pass, time.Duration(100000))
	if err != nil {
		t.Errorf("make token failed: %v", err)
	}
	newId, err := ValidateJWT(token, pass2)
	if err == nil {
		t.Errorf("validate token failed: %v", err)
	}
	if ogId != newId {
		t.Errorf("id mismatch: %v, %v", ogId, newId)
	}
}
