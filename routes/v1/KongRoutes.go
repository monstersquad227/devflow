package v1

import (
	"devflow/controller"
	"devflow/service"
	"github.com/gin-gonic/gin"
)

func KongRegister(api *gin.RouterGroup) {

	kongController := &controller.KongController{
		KongService: &service.KongService{},
	}

	api.GET("/kong/routes")
	api.GET("/kong/services")
	api.GET("/kong/upstreams", kongController.GetUpstreams)
	api.GET("/kong/targets")
	api.GET("/plugins")

}
