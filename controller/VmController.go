package controller

import (
	"devflow/model"
	"devflow/service"
	"devflow/utils"
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type VmController struct {
	VmService *service.VmService
}

func (v *VmController) GetVms(c *gin.Context) {
	number := c.Query("pageNumber")
	size := c.Query("pageSize")

	pageNumber, err := strconv.Atoi(number)
	if err != nil {
		c.JSON(400, utils.Error(1, "pageNumber错误", err))
		return
	}
	pageSize, err := strconv.Atoi(size)
	if err != nil {
		c.JSON(400, utils.Error(1, "pageSize错误", err))
		return
	}
	data, count, err := v.VmService.FetchVms(pageNumber, pageSize)
	if err != nil {
		c.JSON(500, utils.Error(1, "查询失败", err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"data":  data,
		"total": count,
	}))
}

func (v *VmController) CreateVm(c *gin.Context) {
	req := model.Vm{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, utils.Error(1, "参数错误: "+err.Error(), err))
		return
	}

	password, err := base64.StdEncoding.DecodeString(req.Password)
	if err != nil {
		c.JSON(400, utils.Error(1, "base64: "+err.Error(), err))
		return
	}
	req.Password = string(password)

	lastId, err := v.VmService.SaveVm(req)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误: "+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"LastInsertId": lastId,
	}))
}

func (v *VmController) DeleteVm(c *gin.Context) {
	vmId := c.Param("vm")
	if vmId == "" {
		c.JSON(400, utils.Error(1, "参数错误", errors.New(":vm 为空")))
		return
	}
	id, err := strconv.Atoi(vmId)
	if err != nil {
		c.JSON(400, utils.Error(1, "strconv 错误: "+err.Error(), err))
		return
	}
	affectedId, err := v.VmService.RemoveVm(id)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误"+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"rowsAffectedId": affectedId,
	}))
}

func (v *VmController) GetVmPasswordById(c *gin.Context) {
	vmId := c.Param("vm")
	if vmId == "" {
		c.JSON(400, utils.Error(1, "参数错误", errors.New(":vm 为空")))
		return
	}
	id, err := strconv.Atoi(vmId)
	if err != nil {
		c.JSON(400, utils.Error(1, "strconv 错误: "+err.Error(), err))
		return
	}
	password, err := v.VmService.FetchVmPasswordById(id)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误: "+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(password))
}

func (v *VmController) GetVmsByApplication(c *gin.Context) {
	application := c.Param("vm")
	if application == "" {
		c.JSON(400, utils.Error(1, "application不能为空", nil))
		return
	}
	vms, err := v.VmService.FetchVmsByApplication(application)
	if err != nil {
		c.JSON(400, utils.Error(1, "查询失败: "+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(vms))
}
