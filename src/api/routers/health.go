package routers

import (
	"bookstore-project/api/handlers"

	"github.com/gin-gonic/gin"
)

func Health (r *gin.RouterGroup){
	handler := handlers.NewHealthHandler()

	r.GET("/", handler.Health) // recieve get from api
	r.POST("/" , handler.Health)
	
}