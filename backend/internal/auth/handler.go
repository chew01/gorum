package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"strings"
)

type UserClaims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func Handler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if !strings.HasPrefix(authHeader, "Bearer ") {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("claims", claims)
	c.Next()
}
