package middleware

import (
	"fmt"
	"strings"

	"github.com/ekosachev/spara/internal/config"
	"github.com/ekosachev/spara/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware() gin.HandlerFunc {
	var cfg = config.GetConfig()
	var secretKey = cfg.JWTSecret

	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(401, dto.ApiResponse{Success: false, Error: "Authorization header required"})
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.AbortWithStatusJSON(401, dto.ApiResponse{Success: false, Error: "Invalid auth header"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			ctx.AbortWithStatusJSON(401, dto.ApiResponse{Success: false, Error: err.Error()})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := uint(claims["sub"].(float64))
			ctx.Set("userID", userID)
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(401, dto.ApiResponse{Success: false, Error: "Invalid token"})
		}
	}
}
