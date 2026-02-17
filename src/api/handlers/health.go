package handlers

import(
	"net/http"
	"github.com/gin-gonic/gin"
)


type HealthHandler struct{

}
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{} 
} // no dependency it return itself
func (h *HealthHandler) Health(c *gin.Context){
	c.JSON(http.StatusOK,"working")

}

func (h *HealthHandler) HealthPost(c *gin.Context){
	c.JSON(http.StatusOK,"working post")

}