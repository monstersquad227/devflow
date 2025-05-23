package controller

import (
	"devflow/model"
	"devflow/service"
	"devflow/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TaskController struct {
	TaskService service.TaskServiceInterface
}

func (t *TaskController) ListTasks(c *gin.Context) {
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

	result, err := t.TaskService.List(pageNumber, pageSize)
	if err != nil {
		c.JSON(500, utils.Error(1, "查询错误"+err.Error(), err))
		return
	}
	count, err := t.TaskService.Count()
	if err != nil {
		c.JSON(500, utils.Error(1, "查询错误"+err.Error(), err))
		return
	}

	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"data":  result,
		"total": count,
	}))
}

func (t *TaskController) CreateTask(c *gin.Context) {
	req := &model.Task{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(400, utils.Error(1, "JSON错误: "+err.Error(), err))
		return
	}

	account, _ := c.Get("account")
	req.CreatedBy = account.(string)
	req.UpdatedBy = account.(string)

	latestInsertId, err := t.TaskService.Create(req)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误: "+err.Error(), err))
		return
	}

	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"lastInsertId": latestInsertId,
	}))
}

func (t *TaskController) UpdateTask(c *gin.Context) {
	taskId := c.Param("task")
	if taskId == "" {
		c.JSON(400, utils.Error(1, "参数为空", errors.New("task 参数不能为空")))
		return
	}

	req := &model.Task{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(400, utils.Error(1, "JSON参数错误", err))
		return
	}

	account, _ := c.Get("account")
	req.UpdatedBy = account.(string)
	rowAffected, err := t.TaskService.Update(req)

	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误: "+err.Error(), err))
		return
	}

	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"rowAffected": rowAffected,
	}))
}

func (t *TaskController) DeleteTask(c *gin.Context) {
	taskId := c.Param("task")
	if taskId == "" {
		c.JSON(400, utils.Error(1, "参数为空", errors.New("task 参数不能为空")))
		return
	}

	id, err := strconv.Atoi(taskId)
	if err != nil {
		c.JSON(400, utils.Error(1, "strconv 参数错误: "+err.Error(), err))
		return
	}

	rowAffected, err := t.TaskService.Delete(id)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误: "+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"rowAffected": rowAffected,
	}))
}
