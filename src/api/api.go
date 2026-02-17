package api

import (
	"bookstore-project/api/config"
	"bookstore-project/api/routers"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New() //returns pointer for gin.Engine
	//register new func here
	val,ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile",validation.IrnMobileNumberValidator, true)
	}
	r.Use(gin.Logger(), gin.Recovery())
	//for logging user and recover 500 if there's a panic global middleware
	r.RedirectTrailingSlash = false
    api := r.Group("/api")
	v1 := api.Group("api/v1") // route group and api versioning
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")


		routers.Health(health) //gets router group
		routers.TestRouter(test_router)

	}
    
	
	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
