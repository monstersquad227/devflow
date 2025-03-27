package controller

import (
	"devflow/service"
	"devflow/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type EnvController struct {
	EnvService *service.EnvService
}

func (e *EnvController) GetEnvs(c *gin.Context) {
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

	result, err := e.EnvService.FetchEnvs(pageNumber, pageSize)
	if err != nil {
		c.JSON(500, utils.Error(1, "查询失败: "+err.Error(), err))
		return
	}

	count, err := e.EnvService.FetchEnvsCount()
	if err != nil {
		c.JSON(500, utils.Error(1, "查询失败:"+err.Error(), err))
		return
	}

	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"data":  result,
		"total": count,
	}))
}

func (e *EnvController) GetNamespacesByEnv(c *gin.Context) {
	env := c.Param("env")
	result, err := e.EnvService.FetchNamespacesByEnv(env)
	if err != nil {
		c.JSON(500, utils.Error(1, "查询失败: "+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}
