package utils

import "github.com/gin-gonic/gin"

func RespondJSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{
		"status":  statusCode,
		"message": "success",
		"data":    data,
	})
}

func ResponseError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"status":  statusCode,
		"message": message,
	})
}
