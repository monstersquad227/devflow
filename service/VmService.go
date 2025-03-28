package service

import (
	"devflow/model"
	"devflow/repository"
)

type VmService struct {
	VmRepo *repository.VmRepository
}

func (service *VmService) FetchVms(pageNumber, pageSize int) ([]*model.Vm, int, error) {
	vms, err := service.VmRepo.GetVms(pageNumber, pageSize)
	if err != nil {
		return nil, 0, err
	}
	count, err := service.VmRepo.GetVmsCount()
	if err != nil {
		return nil, 0, err
	}
	return vms, count, nil
}

func (service *VmService) FetchVmsByApplication(application string) (interface{}, error) {
	return service.VmRepo.GetVmsByApplication(application)
}
