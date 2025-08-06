package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	var fail uuid.NullUUID
	var claims jwt.RegisteredClaims
	token, err := jwt.ParseWithClaims(tokenString, claims, nil)
	if err != nil {
		return fail.UUID, fmt.Errorf("jwt error: %v", err)
	}
	sub, err := token.Claims.GetSubject()
	if err != nil {
		return fail.UUID, fmt.Errorf("jwt error: %v", err)
	}
	id, err := uuid.Parse(sub)
	if err != nil {
		return fail.UUID, fmt.Errorf("jwt error: %v", err)
	}
	return id, nil
}
