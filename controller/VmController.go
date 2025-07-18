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
	VmService service.VmServiceInterface
}

func (ctrl *VmController) ListVms(c *gin.Context) {
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
	data, err := ctrl.VmService.List(pageNumber, pageSize)
	if err != nil {
		c.JSON(500, utils.Error(1, "查询失败", err))
		return
	}
	count, err := ctrl.VmService.Count()
	if err != nil {
		c.JSON(500, utils.Error(1, "查询失败", err))
		return
	}

	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"data":  data,
		"total": count,
	}))
}

func (ctrl *VmController) CreateVm(c *gin.Context) {
	req := &model.Vm{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(400, utils.Error(1, "参数错误: "+err.Error(), err))
		return
	}

	password, err := base64.StdEncoding.DecodeString(req.Password)
	if err != nil {
		c.JSON(400, utils.Error(1, "base64: "+err.Error(), err))
		return
	}
	req.Password = string(password)

	switch req.CloudProvider {
	case "aliyun":
		lastId, err := ctrl.VmService.CreateAliyunVm(req)
		if err != nil {
			c.JSON(500, utils.Error(1, "内部错误Aliyun : "+err.Error(), err))
			return
		}
		c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
			"LastInsertId": lastId,
		}))
		return
	case "huawei":
		c.JSON(http.StatusOK, "ING...")
	case "tencent":
		c.JSON(http.StatusOK, "ING...")
	case "aws":
		c.JSON(http.StatusOK, "ING...")
	case "local":
		lastId, err := ctrl.VmService.Create(req)
		if err != nil {
			c.JSON(500, utils.Error(1, "内部错误: "+err.Error(), err))
			return
		}
		c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
			"LastInsertId": lastId,
		}))
		return
	}
}

func (ctrl *VmController) UpdateVm(c *gin.Context) {
	req := &model.Vm{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(400, utils.Error(1, "JSON错误: "+err.Error(), err))
		return
	}
	rowAffected, err := ctrl.VmService.Update(req)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误: "+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"rowAffected": rowAffected,
	}))
}

func (ctrl *VmController) DeleteVm(c *gin.Context) {
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
	affectedId, err := ctrl.VmService.Delete(id)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误"+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"rowsAffectedId": affectedId,
	}))
}

func (ctrl *VmController) GetVmPasswordById(c *gin.Context) {
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
	password, err := ctrl.VmService.FetchVmPasswordById(id)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误: "+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(password))
}

func (ctrl *VmController) GetVmsByApplication(c *gin.Context) {
	application := c.Param("vm")
	if application == "" {
		c.JSON(400, utils.Error(1, "application不能为空", nil))
		return
	}
	vms, err := ctrl.VmService.FetchVmsByApplication(application)
	if err != nil {
		c.JSON(400, utils.Error(1, "查询失败: "+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(vms))
}

func (ctrl *VmController) GetUsersByVm(c *gin.Context) {
	vm := c.Param("vm")
	vmId, err := strconv.Atoi(vm)
	if err != nil {
		c.JSON(400, utils.Error(1, err.Error(), nil))
		return
	}
	result, err := ctrl.VmService.FetchUserByVm(vmId)
	if err != nil {
		c.JSON(500, utils.Error(1, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}
