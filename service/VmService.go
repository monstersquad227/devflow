package service

import (
	"devflow/model"
	"devflow/repository"
	"devflow/utils"
	"encoding/base64"
)

type VmService struct {
	VmRepo *repository.VmRepository
}

func (service *VmService) List(pageNumber, pageSize int) ([]*model.Vm, error) {
	return service.VmRepo.ListVms(pageNumber, pageSize)
}

func (service *VmService) Count() (int, error) {
	return service.VmRepo.CountVms()
}

func (service *VmService) Create(vm model.Vm) (int64, error) {
	encryptPassword, err := utils.EncryptAESGCM(vm.Password)
	if err != nil {
		return 0, err
	}
	vm.Password = encryptPassword
	return service.VmRepo.CreateVm(vm)
}

func (service *VmService) Update(vm model.Vm) (int64, error) {
	return service.VmRepo.UpdateVm(vm)
}

func (service *VmService) Delete(id int) (int64, error) {
	return service.VmRepo.DeleteVm(id)
}

func (service *VmService) FetchVmPasswordById(id int) (string, error) {
	password, err := service.VmRepo.GetVmPasswordById(id)
	if err != nil {
		return "", err
	}
	if password == "" {
		return "", nil
	}
	decryptPassword, err := utils.DecryptAESGCM(password)
	if err != nil {
		return "", err
	}
	encodePassword := base64.StdEncoding.EncodeToString([]byte(decryptPassword))
	return encodePassword, nil
}

func (service *VmService) FetchVmsByApplication(application string) (interface{}, error) {
	return service.VmRepo.GetVmsByApplication(application)
}
