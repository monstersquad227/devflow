package controller

import (
	"devflow/model"
	"devflow/service"
	"devflow/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type EnvController struct {
	EnvService service.EnvServiceInterface
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

	result, err := e.EnvService.Fetch(pageNumber, pageSize)
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

func (e *EnvController) Get(c *gin.Context) {
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

	result, err := e.EnvService.Fetch(pageNumber, pageSize)
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

func (e *EnvController) CreateEnv(c *gin.Context) {
	req := model.Env{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, utils.Error(1, "JSON错误", nil))
		return
	}
	account, _ := c.Get("account")
	req.CreatedBy = account.(string)
	req.UpdatedBy = account.(string)

	lastId, err := e.EnvService.SaveEnv(req)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误: "+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"lastInsertId": lastId,
	}))
}

func (e *EnvController) DeleteEnv(c *gin.Context) {
	envId := c.Param("env")
	if envId == "" {
		c.JSON(400, utils.Error(1, "参数错误: env", nil))
		return
	}
	id, err := strconv.Atoi(envId)
	if err != nil {
		c.JSON(400, utils.Error(1, "strconv 错误: "+err.Error(), err))
		return
	}
	rowAffected, err := e.EnvService.RemoveEnv(id)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误: "+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"rowAffected": rowAffected,
	}))
}

func (e *EnvController) UpdateEnv(c *gin.Context) {
	envId := c.Param("env")
	req := model.Env{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, utils.Error(1, "JSON错误: "+err.Error(), err))
		return
	}

	id, err := strconv.Atoi(envId)
	if err != nil {
		c.JSON(400, utils.Error(1, "strconv 错误: "+err.Error(), err))
		return
	}

	account, _ := c.Get("account")
	req.Id = id
	req.UpdatedBy = account.(string)

	rowAffected, err := e.EnvService.ModifyEnv(req)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误: "+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"rowAffected": rowAffected,
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
