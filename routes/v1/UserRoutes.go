package v1

import (
	"devflow/controller"
	"devflow/repository"
	"devflow/service"
	"github.com/gin-gonic/gin"
)

func UserRegister(api *gin.RouterGroup) {

	userController := &controller.UserController{
		Service: &service.UserService{
			Repo: &repository.UserRepository{},
		},
	}

	api.POST("/user/login", userController.UserLogin) // √
	api.GET("/getPermission", userController.UserPermission)
}
