package service

import (
	"devflow/config"
	"devflow/model"
	"devflow/repository"
	"devflow/utils"
	"encoding/base64"
	"fmt"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v2/client"
	"github.com/alibabacloud-go/tea/tea"
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

func (service *VmService) CreateAliyunVm(vm model.Vm) (int64, error) {
	client, err := NewAliyunClient()
	if err != nil {
		return 0, err
	}
	fmt.Println(vm.Region)
	instanceType := utils.GenerateInstanceTypeBySpec(vm.Spec)
	request := &ecs20140526.RunInstancesRequest{
		InstanceName:    tea.String(vm.InstanceName),
		Password:        tea.String(vm.Password),
		InstanceType:    tea.String(instanceType),
		RegionId:        tea.String("cn-shanghai"),
		ImageId:         tea.String(config.GlobalConfig.Aliyun.ImageId),
		UserData:        tea.String(utils.GenerateUserData()),
		SecurityGroupId: tea.String("sg-uf60f2wmh0rdiabpbweq"),
		SystemDisk: &ecs20140526.RunInstancesRequestSystemDisk{
			Category: tea.String("cloud_efficiency"),
			Size:     tea.String(utils.GenerateDiskSizeBySpec(vm.Spec)),
		},
		//ZoneId:             tea.String("cn-shanghai-g"),
		InstanceChargeType: tea.String("PostPaid"),
		SpotStrategy:       tea.String("NoSpot"),
		Amount:             tea.Int32(1),
		VSwitchId:          tea.String("vsw-uf65avmqpg8bfdtkjg7i1"),
	}
	resp, err := client.RunInstances(request)
	if err != nil {
		return 0, err
	}
	instances := resp.Body.InstanceIdSets.InstanceIdSet[0]
	encryptPassword, err := utils.EncryptAESGCM(vm.Password)
	if err != nil {
		return 0, err
	}
	vm.Password = encryptPassword
	vm.InstanceId = *instances
	return service.VmRepo.CreateVm(vm)
}
