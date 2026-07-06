package handlers


import (
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"

)

func ProtectedHandler(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(401, gin.H{"error": "no token"})
		return
	}
//deletin that shit and get the pure code 
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		c.JSON(401, gin.H{"error": "invalid token"})
		return
	}

	c.JSON(200, gin.H{"message": "you are inside protected route"})
}