package v1

import (
	"devflow/controller"
	"devflow/repository"
	"devflow/service"
	"github.com/gin-gonic/gin"
)

func UserRegister(api *gin.RouterGroup) {
	userController := controller.NewUserController(
		service.NewUserService(
			repository.NewUserRepository(),
		),
	)

	api.POST("/user/login", userController.Login)
	api.GET("/users", userController.Users)
	api.POST("/user/password", userController.Password)
	//api.GET("/getPermission", userController.Permission)
}
