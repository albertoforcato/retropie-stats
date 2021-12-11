package api_helpers

import (
	"github.com/gin-gonic/gin"
)

type Res struct {
	Status  int
	Message interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

//Response returns basic response structure
func Response(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, map[string]interface{}{"message": data})
}

//ResponseWithError returns basic response structure with error
func ResponseWithError(c *gin.Context, statusCode int, error interface{}) {
	c.JSON(statusCode, map[string]interface{}{"error": error})
}
