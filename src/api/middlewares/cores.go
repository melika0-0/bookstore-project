package middlewares

import (
	"github.com/melika0-0/bookstore-project/api/config"

	"github.com/gin-gonic/gin"
)
//Cross-Origin Resource Sharing.stick with web brows so cores are like password stamp let frontend read backend data 
func CORS(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, Content-Length, Accept-Encoding")
		c.Header("Access-Control-Max-Age", "21600")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}