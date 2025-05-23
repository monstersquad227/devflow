package controller

import (
	"devflow/service"
	"devflow/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FlowEdgeController struct {
	FlowedgeService service.FlowedgeServiceInterface
}

func (crtl *FlowEdgeController) ListFlowedges(c *gin.Context) {
	number := c.Query("pageNumber")
	size := c.Query("pageSize")

	pageNumber, err := strconv.Atoi(number)
	if err != nil {
		c.JSON(400, utils.Error(1, "pageNumber 参数错误", err))
		return
	}
	pageSize, err := strconv.Atoi(size)
	if err != nil {
		c.JSON(400, utils.Error(1, "pageSize 参数错误", err))
		return
	}

	result, err := crtl.FlowedgeService.List(pageNumber, pageSize)
	if err != nil {
		c.JSON(500, utils.Error(1, "查询失败: "+err.Error(), err))
		return
	}

	count, err := crtl.FlowedgeService.Count()
	if err != nil {
		c.JSON(500, utils.Error(1, "查询失败:"+err.Error(), err))
		return
	}

	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"data":  result,
		"total": count,
	}))
}

func (crtl *FlowEdgeController) GetFlowedgesByApplication(c *gin.Context) {
	application := c.Param("flowedge")
	if application == "" {
		c.JSON(400, utils.Error(1, "application不能为空", nil))
		return
	}
	flowedges, err := crtl.FlowedgeService.FetchFlowedgesByApplication(application)
	if err != nil {
		c.JSON(400, utils.Error(1, "查询失败: "+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(flowedges))
}
