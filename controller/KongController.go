package controller

import (
	"devflow/service"
	"devflow/utils"
	"github.com/gin-gonic/gin"
)

type KongController struct {
	KongService *service.KongService
}

func (k *KongController) GetUpstreams(c *gin.Context) {
	result, err := k.KongService.FetchUpstreams()
	if err != nil {
		c.JSON(500, utils.Error(1, "upstreams获取失败: "+err.Error(), err))
		return
	}

	c.JSON(200, utils.Success(result))
}
