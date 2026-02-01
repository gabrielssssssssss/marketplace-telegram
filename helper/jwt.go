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
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func NewJwtToken(userID int64, secretKey string) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 300)),
		},
	}).SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("Unable to create JWT with the given parameters.")
	}
	return token, nil
}

func VerifyJwtToken(token, secretKey string) (*jwt.Token, error) {
	jwt, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return false, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	return jwt, err
}

func GetJwtValue(token, key string) (int64, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid token format")
	}

	payloadBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return 0, fmt.Errorf("decode payload: %w", err)
	}

	var payload map[string]int64
	if err := json.Unmarshal(payloadBytes, &payload); err != nil {
		return 0, fmt.Errorf("unmarshal payload: %w", err)
	}

	val, ok := payload[key]
	if !ok {
		return 0, fmt.Errorf("key %q not found", key)
	}

	return val, nil
}
