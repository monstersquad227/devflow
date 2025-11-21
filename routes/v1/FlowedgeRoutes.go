package v1

import (
	"devflow/controller"
	"devflow/repository"
	"devflow/service"
	"github.com/gin-gonic/gin"
)

func FlowedgeRegister(api *gin.RouterGroup) {
	flowedgeController := controller.NewFlowEdgeController(
		service.NewFlowedgeService(
			repository.NewFlowedgeRepository(),
		),
	)

	api.GET("/flowedges", flowedgeController.List)
	api.GET("/flowedges/:flowedge", flowedgeController.ListByApplication)
	api.PATCH("/flowedges/:flowedge", flowedgeController.PatchApplication)
}
