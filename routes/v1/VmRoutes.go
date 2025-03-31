package v1

import (
	"devflow/controller"
	"devflow/repository"
	"devflow/service"
	"github.com/gin-gonic/gin"
)

func VmRegister(api *gin.RouterGroup) {
	vmController := &controller.VmController{
		VmService: &service.VmService{
			VmRepo: &repository.VmRepository{},
		},
	}
	api.GET("/vms", vmController.GetVms)    // √
	api.POST("/vms", vmController.CreateVm) // √
	api.DELETE("/vms/:vm", vmController.DeleteVm)
	api.GET("/vms/:application", vmController.GetVmsByApplication) // √
}
