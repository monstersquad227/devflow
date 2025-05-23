package service

import (
	"devflow/model"
	"devflow/repository"
)

type FlowedgeService struct {
	FlowedgeRepository *repository.FlowedgeRepository
}

func NewFlowedgeService() *FlowedgeService {
	return &FlowedgeService{}
}

func (fs *FlowedgeService) List(pageNumber, pageSize int) ([]*model.Flowedge, error) {
	return fs.FlowedgeRepository.ListFlowedges(pageNumber, pageSize)
}
func (fs *FlowedgeService) Count() (int, error) {
	return fs.FlowedgeRepository.CountFlowedges()
}

func (fs *FlowedgeService) FetchFlowedgesByApplication(application string) (interface{}, error) {
	return fs.FlowedgeRepository.GetFlowedgeByApplication(application)
}

func (fs *FlowedgeService) Create(flowedge model.Flowedge) (int64, error) {
	return fs.FlowedgeRepository.CreateFlowedge(flowedge)
}

func (fs *FlowedgeService) Update(flowedge model.Flowedge) (int64, error) {
	return fs.FlowedgeRepository.UpdateFlowedgeLastHeartBeat(flowedge)
}
