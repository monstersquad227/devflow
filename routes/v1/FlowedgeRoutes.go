package v1

import (
	"devflow/controller"
	"devflow/repository"
	"devflow/service"
	"github.com/gin-gonic/gin"
)

func FlowedgeRegister(api *gin.RouterGroup) {
	flowedgeController := &controller.FlowEdgeController{
		FlowedgeService: &service.FlowedgeService{
			FlowedgeRepository: &repository.FlowedgeRepository{},
		},
	}

	api.GET("/flowedges", flowedgeController.ListFlowedges)
	api.GET("/flowedges/:flowedge", flowedgeController.GetFlowedgesByApplication)
}
