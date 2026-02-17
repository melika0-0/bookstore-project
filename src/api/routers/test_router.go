package routers

import (
	"bookstore-project/api/handlers"

	"github.com/gin-gonic/gin"
)

func TestRouter(r *gin.RouterGroup) {
	h := handlers.NewTestHandler()

	r.GET("/", h.Test)
	r.GET("/usres", h.Users)
	r.GET("/user/:id", h.UserById)
	r.GET("/user/get-user-by-username/:username", h.UserByUsername)
	r.GET("/user/:id/accounts", h.Accounts)
	r.POST("/add-user", h.AddUser)
	r.POST("/binder/header1", h.HeaderBinder1)
	r.POST("/binder/header2", h.HeaderBinder2)
	r.POST("/binder/Uri", h.UriBinder)
	r.POST("/binder/Body", h.BodyBinder)

}
