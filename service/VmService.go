package service

import (
	"devflow/model"
	"devflow/repository"
	"devflow/utils"
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

func (service *VmService) SaveVm(vm model.Vm) (int64, error) {
	encryptPassword, err := utils.EncryptAESGCM(vm.Password)
	if err != nil {
		return 0, err
	}
	vm.Password = encryptPassword
	return service.VmRepo.CreateVm(vm)
}

func (service *VmService) RemoveVm(id int) (int64, error) {
	return service.VmRepo.DeleteVm(id)
}

func (service *VmService) FetchVmsByApplication(application string) (interface{}, error) {
	return service.VmRepo.GetVmsByApplication(application)
}
