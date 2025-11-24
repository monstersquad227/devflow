package v1

import (
	"devflow/controller"
	"devflow/repository"
	"devflow/service"
	"github.com/gin-gonic/gin"
)

func VmRegister(api *gin.RouterGroup) {
	vmController := controller.NewVmController(
		service.NewVmService(
			repository.NewVmRepository(),
		),
	)

	api.GET("/vms", vmController.List)                     // 机器列表
	api.POST("/vms", vmController.Create)                  // 创建机器
	api.GET("/vms/:vm", vmController.Get)                  // 获取机器详情
	api.PUT("/vms/:vm", vmController.Update)               // 更新机器
	api.DELETE("/vms/:vm", vmController.Delete)            // 删除机器
	api.GET("/vms/:vm/password", vmController.GetPassword) // 获取机器密码
	api.GET("/vms/:vm/users", vmController.ListUsers)      // 获取机器用户列表
	api.PUT("/vms/:vm/users", vmController.UpdateUsers)    // 更新机器用户
}
