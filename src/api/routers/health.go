package routers

import (
	"github.com/melika0-0/bookstore-project/api/handlers"

	"github.com/gin-gonic/gin"
)

func Health (r *gin.RouterGroup){
	handler := handlers.NewHealthHandler()

	r.GET("/", handler.Health) // recieve get from api
	r.POST("/" , handler.Health)
	
}