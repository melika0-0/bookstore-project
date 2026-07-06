package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/melika0-0/bookstore-project/api/handlers"
)

func ProtectedRouter(r *gin.RouterGroup) {

	r.GET("/", handlers.ProtectedHandler)

	bookHandler := handlers.NewBookHandler()

	r.POST("/books", bookHandler.CreateBook)
	r.GET("/books", bookHandler.GetBooks)
}