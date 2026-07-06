package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Header struct {
	UserID  string `header:"UserId"`
	Browser string `header:"Browser"`
}

// DTO (request body)
type PersonData struct {
	FirstName    string `json:"first_name" binding:"required,max=10"`
	LastName     string `json:"last_name"`
	MobileNumber string `json:"mobile_number"`
}

type TestHandler struct{}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "Test",
	})
}

func (h *TestHandler) AddUser(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"result": "AddUser",
		"id":     id,
	})
}

// header manually
func (h *TestHandler) HeaderBinder1(c *gin.Context) {
	userID := c.GetHeader("UserId")

	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder1",
		"id":     userID,
	})
}

// header bind into struct (DTO-style)
func (h *TestHandler) HeaderBinder2(c *gin.Context) {
	var header Header

	if err := c.ShouldBindHeader(&header); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder2",
		"header": header,
	})
}

// URI params
func (h *TestHandler) UriBinder(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")

	c.JSON(http.StatusOK, gin.H{
		"result": "UriBinder",
		"id":     id,
		"name":   name,
	})
}

// JSON body binding (DTO usage)
func (h *TestHandler) BodyBinder(c *gin.Context) {
	var p PersonData

	if err := c.ShouldBindJSON(&p); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"validationError": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "BodyBinder",
		"person": p,
	})
}