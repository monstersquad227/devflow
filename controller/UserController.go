package controller

import (
	"devflow/service"
	"devflow/utils"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	UserService service.UserServiceInterface
}

func (ctrl *UserController) Login(c *gin.Context) {
	var userReq struct {
		Account  string `json:"account"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(400, utils.Error(1, "参数错误: "+err.Error(), err))
		return
	}

	account, err := base64.StdEncoding.DecodeString(userReq.Account)
	if err != nil {
		c.JSON(400, utils.Error(1, "account 参数错误", err))
		return
	}

	password, err := base64.StdEncoding.DecodeString(userReq.Password)
	if err != nil {
		c.JSON(400, utils.Error(1, "password 参数错误", err))
		return
	}

	token, info, err := ctrl.UserService.Login(string(account), string(password))
	if err != nil {
		c.JSON(500, utils.Error(1, "登录失败: "+err.Error(), err))
		return
	}

	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"token": token,
		"user":  info,
	}))
}

func (ctrl *UserController) Users(c *gin.Context) {
	result, err := ctrl.UserService.List()
	if err != nil {
		c.JSON(500, utils.Error(1, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}

func (ctrl *UserController) Permission(c *gin.Context) {
	account, _ := c.Get("account")
	result, err := ctrl.UserService.UserPermission(account.(string))
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误", err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"permissions": result,
	}))
}
