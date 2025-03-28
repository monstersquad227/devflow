package v1

import (
	"devflow/controller"
	"devflow/repository"
	"devflow/service"
	"github.com/gin-gonic/gin"
)

func ProjectRegister(api *gin.RouterGroup) {
	projectController := &controller.ProjectController{
		Service: &service.ProjectService{
			Repo: &repository.ProjectRepository{},
		},
	}

	api.GET("/projects", projectController.GetProjects)                                          // 获取所有项目 √
	api.GET("/projects/:project/builds/details", projectController.GetBuildDetails)              // 获取项目构建详情
	api.GET("/projects/:project/branches", projectController.GetBranches)                        // 获取单个项目所有分支 √
	api.GET("/projects/:project/branches/:branch/details", projectController.GetBranchesDetails) // 获取项目分支的详细信息 √
	api.GET("/projects/build/status", projectController.GetBuildStatus)                          // 获取所有项目构建状态
	api.GET("/projects/:project/:env/tags", projectController.GetProjectTags)                    // 获取单个项目标签 √
	api.POST("/projects", projectController.CreateProject)                                       // 创建新项目 √
	api.POST("/projects/:project/build", projectController.BuildProject)                         // 构建项目 √
	api.POST("/projects/:project/deploy", projectController.DeployProject)                       // 部署项目

	api.PUT("/projects/:project", projectController.UpdateProject)                                         // 更新项目
	api.PUT("/projects/:project/builds/:jobId/status/:status", projectController.UpdateProjectBuildStatus) // 更新构建状态回调 √
	api.DELETE("/projects/:projectId", projectController.DeleteProject)                                    // 删除项目
}
