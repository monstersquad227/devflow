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

	api.GET("/setting/envs", envController.GetEnvs)                            // √
	api.POST("/setting/envs", envController.CreateEnv)                         // √
	api.DELETE("/setting/envs/:env", envController.DeleteEnv)                  // √
	api.PUT("/setting/envs/:env", envController.UpdateEnv)                     //√
	api.GET("/setting/envs/:env/namespaces", envController.GetNamespacesByEnv) // √
	api.GET("/setting/images", imageController.GetImages)                      // √
	api.GET("/setting/tasks", taskController.GetTasks)                         // √
}
