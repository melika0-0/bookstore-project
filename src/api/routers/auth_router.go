package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/melika0-0/bookstore-project/api/handlers"
)

func AuthRouter(r *gin.RouterGroup) {
	r.POST("/login", handlers.LoginHandler)
}