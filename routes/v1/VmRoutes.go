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

	api.GET("/vms", vmController.ListVms)                        // √
	api.POST("/vms", vmController.CreateVm)                      // √
	api.PUT("/vms", vmController.UpdateVm)                       // √
	api.DELETE("/vms/:vm", vmController.DeleteVm)                // √
	api.GET("/vms/:vm/password", vmController.GetVmPasswordById) // √
	api.GET("/vms/:vm", vmController.GetVmsByApplication)        // √
}
