package controller

import (
	"devflow/service"
	"devflow/utils"
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

func (v *VmController) GetVmsByApplication(c *gin.Context) {
	application := c.Param("application")
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
