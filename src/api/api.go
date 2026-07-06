package api

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"github.com/melika0-0/bookstore-project/api/config"
	"github.com/melika0-0/bookstore-project/api/middlewares"
	"github.com/melika0-0/bookstore-project/api/routers"
	"github.com/melika0-0/bookstore-project/api/validation"
)


func InitServer() {
	cfg := config.GetConfig()

	registerValidators()

	r := gin.New()

	// very first func in array we hand it gin context
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(middlewares.CORS(cfg))
	r.RedirectTrailingSlash = false

	public := r.Group("/")
	routers.AuthRouter(public)

	// snapshot recovery cors and logger
	// The snapshot keeps the groups isolated, freezing the middleware array at startup
	protected := r.Group("/protected")

	// Runs only for routes registered inside that specific group (e.g., /protected/books).
	// for limitations auth admin check
	protected.Use(middlewares.Auth())
	routers.ProtectedRouter(protected)

	if err := r.Run(fmt.Sprintf(":%s", cfg.Server.Port)); err != nil {
		log.Fatal(err)
	}
}

// register all custom validators used by Gin binding
func registerValidators() {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		log.Fatal("failed to get validator engine")
	}

	if err := v.RegisterValidation(
		"mobile",
		validation.IranMobileNumberValidator,
	); err != nil {
		log.Fatal("failed to register mobile validator:", err)
	}
}