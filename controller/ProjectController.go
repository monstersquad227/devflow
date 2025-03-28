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
	Service *service.ProjectService
}

/*
GetProjects 获取项目
Params:

	pageNumber	int
	pageSize	int
*/

func (controller *ProjectController) GetProjects(c *gin.Context) {
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

	result, err := controller.Service.FetchProjects(number, size)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误:"+err.Error(), err))
		return
	}

	count, err := controller.Service.FetchProjectsCount()
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误:"+err.Error(), err))
		return
	}

	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"data":  result,
		"total": count,
	}))
}

/*
DeleteProject 删除项目
Params:

	projectId	int
*/

func (controller *ProjectController) DeleteProject(c *gin.Context) {
	projectId := c.Param("project")
	pId, err := strconv.Atoi(projectId)
	if err != nil {
		c.JSON(400, utils.Error(1, "projectId参数错误", err))
		return
	}
	result, err := controller.Service.RemoveProject(pId)
	if err != nil {
		c.JSON(500, utils.Error(1, "Sql错误", err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}

/*
CreateProject 创建项目
Body:

	project	model.project
*/

func (controller *ProjectController) CreateProject(c *gin.Context) {
	req := model.Project{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, utils.Error(1, "参数错误", err))
		return
	}
	result, err := controller.Service.SaveProject(req)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误", err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(map[string]int64{
		"lastID": result,
	}))
}

/*
UpdateProject 更新项目
Body:

	gitlab_id				int
	gitlab_repo				string
	build_template_id		int
	project_build_path		string
	project_package_name	string
*/

func (controller *ProjectController) UpdateProject(c *gin.Context) {
	var req model.Project
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, utils.Error(1, "参数错误", err))
		return
	}

	result, err := controller.Service.ModifyProject(req)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误", err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(map[string]int64{
		"rowsAffected": result,
	}))
}

/*
GetBranches 获取项目分支
Params:

	gitlabID	int
*/

func (controller *ProjectController) GetBranches(c *gin.Context) {
	gitlabId := c.Param("project")
	gId, err := strconv.Atoi(gitlabId)
	if err != nil {
		c.JSON(400, utils.Error(1, "参数错误", err))
		return
	}
	branches, err := controller.Service.FetchBranches(gId)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误", err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(map[string]interface{}{
		"branches": branches,
	}))
}

func (controller *ProjectController) GetBranchesDetails(c *gin.Context) {
	gitlabId := c.Param("project")
	branch := c.Param("branch")
	pid, err := strconv.Atoi(gitlabId)
	if err != nil {
		c.JSON(400, utils.Error(1, "参数错误: "+err.Error(), err))
		return
	}
	result, err := controller.Service.FetchBranchesDetails(pid, branch)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误: "+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}

/*
BuildProject 构建项目
Params:

	pid	int

Body:

	params	model.BuildParams
*/

func (controller *ProjectController) BuildProject(c *gin.Context) {
	pid := c.Param("project") //项目ID
	projectId, err := strconv.Atoi(pid)
	if err != nil {
		c.JSON(400, utils.Error(1, "参数错误", err))
		return
	}
	req := model.BuildParams{}
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, utils.Error(1, "参数错误", err))
		return
	}
	createBy, _ := c.Get("account")
	req.CreatedBy = createBy.(string)
	result, err := controller.Service.BuildProjectV2(req, projectId)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误: "+err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}

/*
DeployProject 发布项目
Params:

	deploymentName string

Body:

	deploy_type	string
*/

func (controller *ProjectController) DeployProject(c *gin.Context) {
	req := model.ProjectDeploy{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, utils.Error(1, "参数错误", err))
		return
	}

	result, err := controller.Service.DeployProject(req)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误", err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}

/*
GetBuildDetails 获取构建项目详情
Params:

	projectId	int
*/

func (controller *ProjectController) GetBuildDetails(c *gin.Context) {
	projectId := c.Param("project")
	pid, err := strconv.Atoi(projectId)
	if err != nil {
		c.JSON(400, utils.Error(1, "参数错误", err))
		return
	}

	result, err := controller.Service.GetBuildDetails(pid)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误", err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}

/*
UpdateProjectBuildStatus 更新项目构建状态
Params:

	projectName	string
	jobId		int
	status		string
*/

func (controller *ProjectController) UpdateProjectBuildStatus(c *gin.Context) {
	projectName := c.Param("project")
	jobIdStr := c.Param("jobId")
	status := c.Param("status")

	jobId, err := strconv.Atoi(jobIdStr[1:])
	if err != nil {
		c.JSON(400, utils.Error(1, "参数错误", err))
		return
	}

	result, err := controller.Service.ModifyProjectBuildStatus(projectName, status, jobId)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误", err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(map[string]int64{
		"rowsAffected": result,
	}))
}

/*
GetBuildStatus 获取 Jenkins 构建状态的项目
*/

func (controller *ProjectController) GetBuildStatus(c *gin.Context) {
	result, err := controller.Service.GetBuildStatus()
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误", err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}

/*
GetProjectTags 通过项目名和环境名获取对应的版本
Params:

	project 项目名
	env	环境名
*/

func (controller *ProjectController) GetProjectTags(c *gin.Context) {
	projectName := c.Param("project")
	env := c.Param("env")
	if projectName == "" || env == "" {
		c.JSON(400, utils.Error(1, "参数错误", errors.New("project or env 参数不存在")))
		return
	}
	result, err := controller.Service.FetchProjectsTags(projectName, env)
	if err != nil {
		c.JSON(500, utils.Error(1, "内部错误", err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}
