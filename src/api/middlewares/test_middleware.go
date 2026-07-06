package middlewares
import (
"net/http"
	"github.com/gin-gonic/gin"
)
func TestMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
            "result": "Api key is required",
        })
    }
}
