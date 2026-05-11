package middlewares

import (
	"github.com/melika0-0/bookstore-project/api/config"

	"github.com/gin-gonic/gin"
)

func cores(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Accsess-Control-Allow-Origin", cfg.Cors.AllowedOrigin)
		c.Header("Access-Control-Allow-Origin-credential", "true")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,Update")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization,Content-Type,Content-Length,Accept-Encoding")
		c.Header("Access-max-age", "21600")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return




	}

	}
}