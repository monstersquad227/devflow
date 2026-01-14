package service

import (
	"devflow/model"
	"devflow/repository"
	"encoding/json"
)

type FlowedgeService struct {
	FlowedgeRepository repository.FlowedgeRepositoryInterface
}

func NewFlowedgeService(repo repository.FlowedgeRepositoryInterface) *FlowedgeService {
	return &FlowedgeService{
		FlowedgeRepository: repo,
	}
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

func (fs *FlowedgeService) Create(flowedge *model.FlowedgeCreateRequest) (int64, error) {
	if flowedge.Metadata == nil || len(flowedge.Metadata) == 0 {
		flowedge.Metadata = map[string]string{
			"JVM_OPTIONS":   "TRUE",
			"APM_OPTIONS":   "FALSE",
			"SERVER_IP":     "true",
			"NETWORK_SCOPE": "PUBLIC",
		}
	}
	metadata, err := json.Marshal(flowedge.Metadata)
	if err != nil {
		return 0, err
	}
	req := &model.Flowedge{
		AgentID:  flowedge.AgentID,
		Hostname: flowedge.Hostname,
		Metadata: string(metadata),
		Version:  flowedge.Version,
		Status:   flowedge.Status,
	}
	return fs.FlowedgeRepository.CreateFlowedge(req)
}

func (fs *FlowedgeService) Update(flowedge *model.Flowedge) (int64, error) {
	return fs.FlowedgeRepository.UpdateFlowedgeLastHeartBeat(flowedge)
}

func (fs *FlowedgeService) PatchApplication(flowedge *model.FlowedgePatchRequest) (*model.FlowedgePatchResponse, error) {
	rowsAffected, err := fs.FlowedgeRepository.UpdateFlowedgeApplication(flowedge)
	if err != nil {
		return nil, err
	}
	return &model.FlowedgePatchResponse{RowsAffected: rowsAffected}, nil
}
