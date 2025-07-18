package v1

import (
	"devflow/controller"
	"devflow/repository"
	"devflow/service"
	"github.com/gin-gonic/gin"
)

func UserRegister(api *gin.RouterGroup) {

	UserController := &controller.UserController{
		UserService: &service.UserService{
			UserRepo: &repository.UserRepository{},
		},
	}

	api.POST("/user/login", UserController.Login) // √
	api.GET("/users", UserController.Users)
	api.GET("/getPermission", UserController.Permission)
}
