package middlewares
// access  to request
func TestMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context)
	   apiKey := ctx.GetHeader("x-api-key")
	   if apiKey == "1" {
		ctx.Next()
	   }
	   ctx.AbortWithStatusJSON(http.StatusUnathorized , gin.H{
		"result": "Api key is required"
	   })
}