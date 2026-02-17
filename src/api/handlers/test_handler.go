package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type header struct{
	UserId string 
	Browser string
}
//bodyyy
type personDate struct {
	firstname string `json :"first_name" , binding: "required,max=10` //using tags for validation
	lastname string 
	MobileNumber string

}

type TestHandler struct{
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

//send json responnse back to client
func (h *TestHandler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "Test",
	})
}

func (h *TestHandler) AddUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"result": "AddUser",
		"id": id,
	})
}
// trying to find userid through header
func (h *TestHandler) HeaderBinder1(c *gin.Context){
	Userid := c.GetHeader("UserId")
	c.JSON(http.StatusOK, gin.H{
		"result": "Headerbinder1",
		"id": Userid,
	})

}

func (h *TestHandler) HeaderBinder2(c *gin.Context){
	header := header{}
     c.BindHeader(&header) //bind in obj //dont need error
	c.JSON(http.StatusOK, gin.H{
		"result": "Headerbinder2",
		"header" : header,
	})

}

func (h *TestHandler) UriBinder(c *gin.Context){
	id := c.Param("id")
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"result": "UriBinder",
		"id" : id,
		"name" : name ,
	})

}

func (h *TestHandler) BodyBinder(c *gin.Context){
	p := personDate{}
	err := c.ShouldBindJSON(&p)
	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"validationError": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "BodyBinder", 
		"person" : p,
	})

}
