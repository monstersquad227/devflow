package controller

import (
	"devflow/service"
	"devflow/utils"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	Service *service.UserService
}

func (controller *UserController) UserLogin(c *gin.Context) {
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

	token, info, err := controller.Service.UserLogin(string(account), string(password))
	if err != nil {
		c.JSON(500, utils.Error(1, "登录失败: "+err.Error(), err))
		return
	}

	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"token": token,
		"user":  info,
	}))
}

func (controller *UserController) UserPermission(c *gin.Context) {
	account, _ := c.Get("account")
	result, err := controller.Service.UserPermission(account.(string))
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误", err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"permissions": result,
	}))
}
