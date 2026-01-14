package v1

import (
	"devflow/controller"
	"devflow/repository"
	"devflow/service"
	"github.com/gin-gonic/gin"
)

func SettingRegister(api *gin.RouterGroup) {
	// envController 初始化
	envController := controller.NewEnvController(
		service.NewEnvService(
			repository.NewEnvRepository(),
		),
	)

	// imageController 初始化
	imageController := controller.NewImagesController(
		service.NewImageService(
			repository.NewImageRepository(),
		),
	)

	// taskController 初始化
	taskController := controller.NewTaskController(
		service.NewTaskService(
			repository.NewTaskRepository(),
		),
	)

	// 环境管理路由
	envRoutes := api.Group("/setting/envs")
	{
		envRoutes.GET("", envController.List)                         // 获取环境列表
		envRoutes.POST("", envController.Create)                      // 创建环境
		envRoutes.PUT("/:id", envController.Update)                   // 更新环境
		envRoutes.DELETE("/:id", envController.Delete)                // 删除环境
		envRoutes.GET("/:id/namespaces", envController.GetNamespaces) // 获取环境的命名空间
	}

	// 镜像管理路由
	imageRoutes := api.Group("/setting/images")
	{
		imageRoutes.GET("", imageController.List)          // 获取镜像列表
		imageRoutes.POST("", imageController.Create)       // 创建镜像
		imageRoutes.PUT("/:id", imageController.Update)    // 更新镜像
		imageRoutes.DELETE("/:id", imageController.Delete) // 删除镜像
	}

	// 任务管理路由
	taskRoutes := api.Group("/setting/tasks")
	{
		taskRoutes.GET("", taskController.List)                 // 获取任务列表
		taskRoutes.POST("", taskController.Create)              // 创建任务
		taskRoutes.PUT("/:id", taskController.Update)           // 更新任务
		taskRoutes.DELETE("/:id", taskController.Delete)        // 删除任务
		taskRoutes.GET("/:id/detail", taskController.GetDetail) // 获取单个任务配置
	}
}
