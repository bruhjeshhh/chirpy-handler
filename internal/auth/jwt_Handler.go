package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {
	token := headers.Get("Authorization")

	ans := strings.HasPrefix(token, "Bearer ")

	if ans == false {
		return "", fmt.Errorf("auth header missing")
	}

	trimmed := strings.TrimSpace(strings.TrimPrefix(token, "Bearer "))
	if trimmed == "" {
		return "", fmt.Errorf("the auth token is invalid")
	}

	return trimmed, nil

}

func GetAPIKey(headers http.Header) (string, error) {
	token := headers.Get("Authorization")

	ans := strings.HasPrefix(token, "ApiKey ")

	if ans == false {
		return "", fmt.Errorf("auth header missing")
	}

	trimmed := strings.TrimSpace(strings.TrimPrefix(token, "ApiKey "))
	if trimmed == "" {
		return "", fmt.Errorf("the auth token is invalid")
	}

	return trimmed, nil
}
