package helper

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID int64
	jwt.RegisteredClaims
}

func GenerateJwtToken(UserID int64, SecretKey []byte) (string, error) {
	claims := Claims{
		UserID: UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 300)),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(SecretKey)
	if err != nil {
		return "", fmt.Errorf("Unable to create JWT with the given parameters.")
	}

	return token, nil
}

func GetJwtValue(token, key string) (string, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return "", fmt.Errorf("JWT token is invalid.")
	}

	payloadBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", err
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(payloadBytes, &payload); err != nil {
		panic(err)
	}

	key, ok := payload[key].(string)
	if !ok {
		return "", fmt.Errorf("JWT key not found.")
	}

	return key, nil
}
