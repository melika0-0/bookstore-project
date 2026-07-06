package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("Secretkey")

type Response struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Token   string `json:"token"`
}

type UserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func LoginHandler(c *gin.Context) {

	var userRequest UserRequest
	if err := c.BindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Error: "Invalid request",
		})
		return
	}

	if userRequest.Name == "admin" && userRequest.Password == "password" {

		claims := jwt.MapClaims{
			"username": userRequest.Name,
			"exp":      time.Now().Add(26 * time.Minute).Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		signedToken, err := token.SignedString(jwtSecret)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Response{
				Error: "Could not generate token",
			})
			return
		}

		c.JSON(http.StatusOK, Response{
			Message: "successful",
			Token:   signedToken,
		})
		return
	}

	c.JSON(http.StatusUnauthorized, Response{
		Error: "invalid",
	})
}


//we configure the rout handler 
