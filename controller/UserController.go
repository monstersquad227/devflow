package controller

import (
	"devflow/model"
	"devflow/service"
	"devflow/utils"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	UserService service.UserServiceInterface
}

func NewUserController(svc service.UserServiceInterface) *UserController {
	return &UserController{
		UserService: svc,
	}
}

func (ctrl *UserController) Login(c *gin.Context) {
	req := &model.LoginRequest{}

	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(400, utils.Error(1, "参数错误: "+err.Error(), err))
		return
	}

	account, err := base64.StdEncoding.DecodeString(req.Account)
	if err != nil {
		c.JSON(400, utils.Error(1, "account 参数错误", err))
		return
	}

	password, err := base64.StdEncoding.DecodeString(req.Password)
	if err != nil {
		c.JSON(400, utils.Error(1, "password 参数错误", err))
		return
	}

	result, err := ctrl.UserService.Login(string(account), string(password))
	if err != nil {
		c.JSON(500, utils.Error(1, "登录失败: "+err.Error(), err))
		return
	}

	c.JSON(http.StatusOK, utils.Success(result))
	//c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
	//	"user": map[string]interface{}{
	//		"id":      3,
	//		"account": "yening",
	//		"name":    "叶宁",
	//		"email":   "yening@chengduoduo.net",
	//		"mobile":  "15056332824",
	//		"status":  1,
	//	},
	//
	//	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50IjoieWVuaW5nIiwiY3JlYXRlX2F0IjoiMjAyNS0xMi0xMFQwOTowMToxOS42Mjg4NzQzMzQrMDg6MDAiLCJleHAiOjE3NjUzNzE2Nzl9.d2NP-7RPOlCRExPQCUB2RS9w_OJJ5BHh-U4YuQ07A8w",
	//
	//	"roles": []map[string]interface{}{
	//		{
	//			"id":       1,
	//			"roleCode": "Owner",
	//			"roleName": "超级管理员",
	//		},
	//	},
	//
	//	"permissions": []string{
	//		"project", "project:edit", "project:delete", "project:deploy",
	//		"vm", "vm:add", "vm:edit", "vm:delete", "vm:view",
	//		"setting", "setting:add", "setting:edit", "setting:delete",
	//		"flowedge", "flowedge:add", "flowedge:edit", "flowedge:delete",
	//	},
	//
	//	"menus": []map[string]interface{}{
	//		{
	//			"id":             1,
	//			"permissionCode": "project",
	//			"permissionName": "项目列表",
	//			"path":           "/project",
	//			"children":       []interface{}{},
	//		},
	//		{
	//			"id":             8,
	//			"permissionCode": "vm",
	//			"permissionName": "机器列表",
	//			"path":           "/vm",
	//			"children":       []interface{}{},
	//		},
	//		{
	//			"id":             13,
	//			"permissionCode": "setting",
	//			"permissionName": "配置列表",
	//			"path":           "/setting",
	//			"children":       []interface{}{},
	//		},
	//		{
	//			"id":             19,
	//			"permissionCode": "flowedge",
	//			"permissionName": "Edge列表",
	//			"path":           "/flowedge",
	//			"children":       []interface{}{},
	//		},
	//	},
	//}))
}

func (ctrl *UserController) Users(c *gin.Context) {
	result, err := ctrl.UserService.List()
	if err != nil {
		c.JSON(500, utils.Error(1, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}

func (ctrl *UserController) Password(c *gin.Context) {
	req := &model.PasswordRequest{}

	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(400, utils.Error(1, "参数错误: "+err.Error(), err))
		return
	}

	account, err := base64.StdEncoding.DecodeString(req.Account)
	if err != nil {
		c.JSON(400, utils.Error(1, "account 参数错误", err))
		return
	}
	password, err := base64.StdEncoding.DecodeString(req.Password)
	if err != nil {
		c.JSON(400, utils.Error(1, "password 参数错误", err))
		return
	}
	newPassword, err := base64.StdEncoding.DecodeString(req.NewPassword)
	if err != nil {
		c.JSON(400, utils.Error(1, "new_password 参数错误", err))
		return
	}
	confirmNewPassword, err := base64.StdEncoding.DecodeString(req.ConfirmNewPassword)
	if err != nil {
		c.JSON(400, utils.Error(1, "confirm_new_password 参数错误", err))
		return
	}

	req.Account = string(account)
	req.Password = string(password)
	req.NewPassword = string(newPassword)
	req.ConfirmNewPassword = string(confirmNewPassword)

	result, err := ctrl.UserService.PasswordChange(req)
	if err != nil {
		c.JSON(500, utils.Error(1, "修改密码失败: "+err.Error(), err))
		return
	}

	c.JSON(200, utils.Success(result))
}

//func (ctrl *UserController) Permission(c *gin.Context) {
//	account, _ := c.Get("account")
//	result, err := ctrl.UserService.UserPermission(account.(string))
//	if err != nil {
//		c.JSON(500, utils.Error(1, "内部错误", err))
//		return
//	}
//	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
//		"permissions": result,
//	}))
//}
