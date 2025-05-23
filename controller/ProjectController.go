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

type ProjectController struct {
	Service service.ProjectServiceInterface
}

func (ctrl *ProjectController) ListProjects(c *gin.Context) {
	pageNumber := c.Query("pageNumber")
	pageSize := c.Query("pageSize")

	number, err := strconv.Atoi(pageNumber)
	if err != nil {
		c.JSON(400, utils.Error(1, "pageNumber参数错误", err))
		return
	}
	size, err := strconv.Atoi(pageSize)
	if err != nil {
		c.JSON(400, utils.Error(1, "pageSize参数错误", err))
		return
	}

	result, err := ctrl.Service.List(number, size)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误:"+err.Error(), err))
		return
	}

	count, err := ctrl.Service.Count()
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误:"+err.Error(), err))
		return
	}

	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"data":  result,
		"total": count,
	}))
}

func (ctrl *ProjectController) CreateProject(c *gin.Context) {
	req := &model.Project{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(400, utils.Error(1, "参数错误", err))
		return
	}
	result, err := ctrl.Service.Create(req)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误", err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(map[string]int64{
		"lastID": result,
	}))
}

func (ctrl *ProjectController) UpdateProject(c *gin.Context) {
	projectId := c.Param("project")
	if projectId == "" {
		c.JSON(400, utils.Error(1, "参数错误", errors.New("project 参数不为空")))
		return
	}

	req := &model.Project{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(400, utils.Error(1, "JSON错误", err))
		return
	}

	result, err := ctrl.Service.Update(req)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误: "+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(map[string]int64{
		"rowAffected": result,
	}))
}

func (ctrl *ProjectController) DeleteProject(c *gin.Context) {
	projectId := c.Param("project")
	pId, err := strconv.Atoi(projectId)
	if err != nil {
		c.JSON(400, utils.Error(1, "参数错误", err))
		return
	}
	result, err := ctrl.Service.Delete(pId)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误: "+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}

func (ctrl *ProjectController) ListProjectApplications(c *gin.Context) {
	applications, err := ctrl.Service.ListProjectApplications()
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误: "+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(applications))
}

func (ctrl *ProjectController) ListBranches(c *gin.Context) {
	gitlabId := c.Param("project")
	gId, err := strconv.Atoi(gitlabId)
	if err != nil {
		c.JSON(400, utils.Error(1, "参数错误", err))
		return
	}
	branches, err := ctrl.Service.ListBranches(gId)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误", err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"branches": branches,
	}))
}

func (ctrl *ProjectController) ListBranchesDetails(c *gin.Context) {
	gitlabId := c.Param("project")
	branch := c.Param("branch")
	pid, err := strconv.Atoi(gitlabId)
	if err != nil {
		c.JSON(400, utils.Error(1, "参数错误: "+err.Error(), err))
		return
	}
	result, err := ctrl.Service.ListBranchesDetails(pid, branch)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误: "+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}

func (ctrl *ProjectController) BuildProject(c *gin.Context) {
	pid := c.Param("project") //项目ID
	projectId, err := strconv.Atoi(pid)
	if err != nil {
		c.JSON(400, utils.Error(1, "参数错误", err))
		return
	}
	req := &model.BuildParams{}
	if err = c.ShouldBindJSON(req); err != nil {
		c.JSON(400, utils.Error(1, "参数错误", err))
		return
	}
	createBy, _ := c.Get("account")
	req.CreatedBy = createBy.(string)
	result, err := ctrl.Service.Build(req, projectId)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误: "+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}

func (ctrl *ProjectController) DeployProject(c *gin.Context) {
	req := &model.ProjectDeploy{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(400, utils.Error(1, "参数错误", err))
		return
	}

	result, err := ctrl.Service.Deploy(req)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误"+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}

func (ctrl *ProjectController) ListBuildDetails(c *gin.Context) {
	projectId := c.Param("project")
	pid, err := strconv.Atoi(projectId)
	if err != nil {
		c.JSON(400, utils.Error(1, "参数错误", err))
		return
	}

	result, err := ctrl.Service.ListBuildDetails(pid)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误", err))
		return
	}

	count, err := ctrl.Service.CountBuildDetails(pid)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误: "+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"total": count,
		"data":  result,
	}))
}

func (ctrl *ProjectController) ListBuildDetailsText(c *gin.Context) {
	id := c.Param("id")
	pid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, utils.Error(1, "参数错误", err))
		return
	}
	text, err := ctrl.Service.ListBuildDetailsText(pid)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误: "+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(text))
}

func (ctrl *ProjectController) ListBuildStatus(c *gin.Context) {
	ing, err := ctrl.Service.ListBuildStatusING()
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误", err))
		return
	}
	fail, err := ctrl.Service.ListBuildStatusFail()
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误", err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"ing":  ing,
		"fail": fail,
	}))
}

func (ctrl *ProjectController) UpdateBuildStatus(c *gin.Context) {
	projectName := c.Param("project")
	jobIdStr := c.Param("jobId")
	status := c.Param("status")

	jobId, err := strconv.Atoi(jobIdStr[1:])
	if err != nil {
		c.JSON(400, utils.Error(1, "参数错误", err))
		return
	}

	result, err := ctrl.Service.UpdateBuildStatus(projectName, status, jobId)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误", err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(map[string]int64{
		"rowsAffected": result,
	}))
}

func (ctrl *ProjectController) ListProjectImageTags(c *gin.Context) {
	projectName := c.Param("project")
	env := c.Param("env")
	if projectName == "" || env == "" {
		c.JSON(400, utils.Error(1, "参数错误", errors.New("project or env 参数不存在")))
		return
	}
	result, err := ctrl.Service.ListProjectImageTags(projectName, env)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误", err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}
