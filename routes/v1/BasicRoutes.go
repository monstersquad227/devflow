package v1

import (
	"devflow/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BasicRegister(api *gin.RouterGroup) {

	api.GET("demo/testGet", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	api.POST("demo/testPost", func(c *gin.Context) {
		var req struct {
			Id int `json:"id"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
			return
		}
		c.JSON(http.StatusOK, req)
	})

	api.GET("/actuator/health", func(c *gin.Context) {
		c.String(http.StatusOK, config.GlobalConfig.Application.Name+" OK")
	})
}
