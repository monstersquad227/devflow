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

type ImagesController struct {
	ImageService service.ImageServiceInterface
}

func (i *ImagesController) ListImages(c *gin.Context) {
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

	result, err := i.ImageService.List(pageNumber, pageSize)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误: "+err.Error(), err))
		return
	}
	count, err := i.ImageService.Count()
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误: "+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"data":  result,
		"total": count,
	}))
}

func (i *ImagesController) CreateImage(c *gin.Context) {
	req := &model.Image{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(400, utils.Error(1, "JSON 错误: "+err.Error(), err))
		return
	}
	account, _ := c.Get("account")
	req.CreatedBy = account.(string)
	req.UpdatedBy = account.(string)

	lastInsertId, err := i.ImageService.Create(req)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误: "+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"lastInsertId": lastInsertId,
	}))
}

func (i *ImagesController) UpdateImage(c *gin.Context) {
	imageId := c.Param("image")
	if imageId == "" {
		c.JSON(400, utils.Error(1, "参数错误", errors.New("image 参数不为空")))
		return
	}

	id, err := strconv.Atoi(imageId)
	if err != nil {
		c.JSON(400, utils.Error(1, "strconv 参数错误"+err.Error(), err))
		return
	}

	req := &model.Image{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(400, utils.Error(1, "JSON 错误: "+err.Error(), err))
		return
	}
	account, _ := c.Get("account")
	req.UpdatedBy = account.(string)
	req.Id = id

	rowAffected, err := i.ImageService.Update(req)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误: "+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"rowAffected": rowAffected,
	}))
}

func (i *ImagesController) DeleteImage(c *gin.Context) {
	imageId := c.Param("image")
	if imageId == "" {
		c.JSON(400, utils.Error(1, "参数错误", errors.New("image 参数不为空")))
		return
	}

	id, err := strconv.Atoi(imageId)
	if err != nil {
		c.JSON(400, utils.Error(1, "strconv 错误", err))
		return
	}

	rowAffected, err := i.ImageService.Delete(id)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误: "+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"rowAffected": rowAffected,
	}))
}
