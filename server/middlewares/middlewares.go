package middlewares

import "github.com/gin-gonic/gin"

// Middlewares function to add middlewares to the application
func Middlewares() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Code for middlewares
		c.Next()
	}
}
