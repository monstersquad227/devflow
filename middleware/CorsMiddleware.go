package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Token, Accept, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		if method == "OPTIONS" {
			c.Status(http.StatusOK)
		}
		c.Next()
	}
}
