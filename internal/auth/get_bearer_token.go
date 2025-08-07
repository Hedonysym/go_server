package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetBearerToken(h http.Header) (string, error) {
	authHeader := h.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("no auth header")
	}
	cleaned, found := strings.CutPrefix(authHeader, "Bearer")
	if !found {
		return "", fmt.Errorf("invalid token header")
	}
	return strings.TrimSpace(cleaned), nil
}
