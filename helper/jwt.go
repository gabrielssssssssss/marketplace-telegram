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

func GetUserID(claims *Claims) int64 {
	return claims.UserID
}

func GetClaims(claims *Claims) *Claims {
	return claims
}

func NewJwtToken(userID int64, secretKey []byte) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 300)),
		},
	}).SignedString(secretKey)
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
		return secretKey, nil
	})
	return jwt, err
}

func GetJwtValue(token, secretKey string) (string, error) {
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
		return "", err
	}

	key, ok := payload[secretKey].(string)
	if !ok {
		return "", fmt.Errorf("JWT key not found.")
	}

	return key, nil
}
