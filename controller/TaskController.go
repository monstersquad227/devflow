package controller

import (
	"devflow/service"
	"devflow/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TaskController struct {
	TaskService *service.TaskService
}

func (t *TaskController) GetTasks(c *gin.Context) {
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

	result, err := t.TaskService.FetchTasks(pageNumber, pageSize)
	if err != nil {
		c.JSON(500, utils.Error(1, "查询错误"+err.Error(), err))
		return
	}
	count, err := t.TaskService.FetchTasksCount()
	if err != nil {
		c.JSON(500, utils.Error(1, "查询错误"+err.Error(), err))
		return
	}

	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"data":  result,
		"total": count,
	}))
}
