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

	imageController := &controller.ImagesController{
		ImageService: &service.ImageService{
			ImageRepository: &repository.ImageRepository{},
		},
	}

	taskController := &controller.TaskController{
		TaskService: &service.TaskService{
			TaskRepository: &repository.TaskRepository{},
		},
	}

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
		imageRoutes.GET("", imageController.ListImages)         // 获取镜像列表
		imageRoutes.POST("", imageController.CreateImage)       // 创建镜像
		imageRoutes.PUT("/:id", imageController.UpdateImage)    // 更新镜像
		imageRoutes.DELETE("/:id", imageController.DeleteImage) // 删除镜像
	}

	// 任务管理路由
	taskRoutes := api.Group("/setting/tasks")
	{
		taskRoutes.GET("", taskController.ListTasks)         // 获取任务列表
		taskRoutes.POST("", taskController.CreateTask)       // 创建任务
		taskRoutes.PUT("/:id", taskController.UpdateTask)    // 更新任务
		taskRoutes.DELETE("/:id", taskController.DeleteTask) // 删除任务
	}
}
