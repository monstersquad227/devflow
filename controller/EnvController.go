package controller

import (
	"devflow/model"
	"devflow/service"
	"devflow/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// EnvController 环境管理控制器
type EnvController struct {
	EnvService service.EnvServiceInterface
}

// NewEnvController 创建环境控制器实例
func NewEnvController(svc service.EnvServiceInterface) *EnvController {
	return &EnvController{
		EnvService: svc,
	}
}

// List 获取环境
func (e *EnvController) List(c *gin.Context) {
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

	result, err := e.EnvService.List(pageNumber, pageSize)
	if err != nil {
		c.JSON(500, utils.Error(1, "查询失败: "+err.Error(), err))
		return
	}

	count, err := e.EnvService.Count()
	if err != nil {
		c.JSON(500, utils.Error(1, "查询失败:"+err.Error(), err))
		return
	}

	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"data":  result,
		"total": count,
	}))
}

// Create 创建环境
func (e *EnvController) Create(c *gin.Context) {
	req := &model.EnvCreateRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(400, utils.Error(1, "JSON错误", nil))
		return
	}
	account, _ := c.Get("account")
	req.CreatedBy = account.(string)
	req.UpdatedBy = account.(string)

	result, err := e.EnvService.Create(req)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误: "+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}

// Update 更新环境
func (e *EnvController) Update(c *gin.Context) {
	//envId := c.Param("id")
	req := &model.EnvUpdateRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(400, utils.Error(1, "JSON错误: "+err.Error(), err))
		return
	}

	//id, err := strconv.Atoi(envId)
	//if err != nil {
	//	c.JSON(400, utils.Error(1, "strconv 错误: "+err.Error(), err))
	//	return
	//}

	account, _ := c.Get("account")
	//req.Id = id
	req.UpdatedBy = account.(string)

	rowAffected, err := e.EnvService.Update(req)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误: "+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"rowAffected": rowAffected,
	}))
}

// Delete 删除环境
func (e *EnvController) Delete(c *gin.Context) {
	envId := c.Param("id")
	if envId == "" {
		c.JSON(400, utils.Error(1, "参数错误: env", nil))
		return
	}
	id, err := strconv.Atoi(envId)
	if err != nil {
		c.JSON(400, utils.Error(1, "strconv 错误: "+err.Error(), err))
		return
	}
	rowAffected, err := e.EnvService.Delete(id)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误: "+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"rowAffected": rowAffected,
	}))
}

// GetNamespaces 获取环境的命名空间列表
func (e *EnvController) GetNamespaces(c *gin.Context) {
	env := c.Param("id")
	result, err := e.EnvService.GetNamespaces(env)
	if err != nil {
		c.JSON(500, utils.Error(1, "查询失败: "+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}
