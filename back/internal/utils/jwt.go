package utils

import (
	"time"

	"github.com/ekosachev/spara/internal/config"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(userID uint, permissions []string) (string, error) {
	cfg := config.GetConfig()
	claims := jwt.MapClaims{
		"sub":         userID,
		"exp":         time.Now().Add(time.Duration(cfg.JWTExpirationSeconds) * time.Second).Unix(),
		"permissions": permissions,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecret))
}
