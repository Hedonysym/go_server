package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetPolkaAuth(h http.Header) (string, error) {
	authHeader := h.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("no auth header")
	}
	cleaned, found := strings.CutPrefix(authHeader, "ApiKey")
	if !found {
		return "", fmt.Errorf("invalid apikey header")
	}
	return strings.TrimSpace(cleaned), nil
}
