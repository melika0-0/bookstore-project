package middlewares

import (
	"strings"
"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/melika0-0/bookstore-project/api/config"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "no token",
			})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		cfg := config.GetConfig()

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {

			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			return []byte(cfg.Jwt.Secretkey), nil
		})

		if err != nil {
	fmt.Println("JWT ERROR:", err)
	c.AbortWithStatusJSON(401, gin.H{
		"error": "invalid token",
	})
	return
}

if !token.Valid {
	fmt.Println("TOKEN NOT VALID")
	c.AbortWithStatusJSON(401, gin.H{
		"error": "invalid token",
	})
	return
 }

		c.Next()
	}
}