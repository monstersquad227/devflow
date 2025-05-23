package v1

import (
	"devflow/controller"
	"devflow/repository"
	"devflow/service"
	"github.com/gin-gonic/gin"
)

func ProjectRegister(api *gin.RouterGroup) {
	crtl := &controller.ProjectController{
		Service: &service.ProjectService{
			Repo: &repository.ProjectRepository{},
		},
	}

	api.GET("/projects", crtl.ListProjects)
	api.POST("/projects", crtl.CreateProject)
	api.PUT("/projects/:project", crtl.UpdateProject)
	api.DELETE("/projects/:project", crtl.DeleteProject)
	api.GET("/projects/applications", crtl.ListProjectApplications)
	api.GET("/projects/:project/branches", crtl.ListBranches)
	api.GET("/projects/:project/branches/:branch/details", crtl.ListBranchesDetails)
	api.POST("/projects/:project/build", crtl.BuildProject)
	api.POST("/projects/:project/deploy", crtl.DeployProject)
	api.GET("/projects/:project/builds/details", crtl.ListBuildDetails)
	api.GET("/projects/build/details/:id/text", crtl.ListBuildDetailsText)
	api.GET("/projects/build/status", crtl.ListBuildStatus)
	api.PUT("/projects/:project/builds/:jobId/status/:status", crtl.UpdateBuildStatus)
	api.GET("/projects/:project/:env/tags", crtl.ListProjectImageTags)
}
