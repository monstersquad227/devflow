package v1

import (
	"devflow/controller"
	"devflow/repository"
	"devflow/service"
	"github.com/gin-gonic/gin"
)

func SettingRegister(api *gin.RouterGroup) {
	envController := &controller.EnvController{
		EnvService: &service.EnvService{
			EnvRepository: &repository.EnvRepository{},
		},
	}

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

	api.GET("/setting/envs", envController.ListEnvs)
	api.POST("/setting/envs", envController.CreateEnv)
	api.PUT("/setting/envs/:env", envController.UpdateEnv)
	api.DELETE("/setting/envs/:env", envController.DeleteEnv)
	api.GET("/setting/envs/:env/namespaces", envController.GetNamespacesByEnv)

	api.GET("/setting/images", imageController.ListImages)            // √
	api.POST("/setting/images", imageController.CreateImage)          // √
	api.DELETE("/setting/images/:image", imageController.DeleteImage) //√
	api.PUT("/setting/images/:image", imageController.UpdateImage)    // √

	api.GET("/setting/tasks", taskController.ListTasks)
	api.POST("/setting/tasks", taskController.CreateTask)
	api.PUT("/setting/tasks/:task", taskController.UpdateTask)
	api.DELETE("/setting/tasks/:task", taskController.DeleteTask)
}
