package service

import (
	"devflow/config"
	"devflow/model"
	"devflow/repository"
	"devflow/utils"
	"encoding/base64"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

type VmService struct {
	VmRepo *repository.VmRepository
}

func (svc *VmService) List(pageNumber, pageSize int) ([]*model.Vm, error) {
	return svc.VmRepo.ListVms(pageNumber, pageSize)
}

func (svc *VmService) Count() (int, error) {
	return svc.VmRepo.CountVms()
}

func (svc *VmService) Create(vm *model.Vm) (int64, error) {
	encryptPassword, err := utils.EncryptAESGCM(vm.Password)
	if err != nil {
		return 0, err
	}
	vm.Password = encryptPassword
	return svc.VmRepo.CreateVm(vm)
}

func (svc *VmService) Update(vm *model.Vm) (int64, error) {
	return svc.VmRepo.UpdateVm(vm)
}

func (svc *VmService) Delete(id int) (int64, error) {
	cloudProvider, err := svc.VmRepo.GetCloudProviderById(id)
	if err != nil {
		return 0, err
	}

	if cloudProvider == "aliyun" {
		instanceID, err := svc.VmRepo.GetInstanceIDById(id)
		if err != nil {
			return 0, err
		}
		client, err := NewAliyunClient()
		if err != nil {
			return 0, err
		}
		regionID, err := svc.VmRepo.GetRegionById(id)
		if err != nil {
			return 0, err
		}
		request := &ecs20140526.DeleteInstancesRequest{
			Force:      tea.Bool(true),
			InstanceId: tea.StringSlice([]string{instanceID}),
			RegionId:   tea.String(regionID),
		}
		_, err = client.DeleteInstances(request)
		if err != nil {
			return 0, err
		}
	}

	return svc.VmRepo.DeleteVm(id)
}

func (svc *VmService) FetchVmPasswordById(id int) (string, error) {
	password, err := svc.VmRepo.GetVmPasswordById(id)
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

func (svc *VmService) FetchVmsByApplication(application string) (interface{}, error) {
	return svc.VmRepo.GetVmsByApplication(application)
}

func (svc *VmService) CreateAliyunVm(vm *model.Vm) (int64, error) {
	client, err := NewAliyunClient()
	if err != nil {
		return 0, err
	}

	request := &ecs20140526.RunInstancesRequest{
		InstanceName:       tea.String(vm.InstanceName),                                                                                         // 实例名称
		HostName:           tea.String(utils.StringToLower(vm.InstanceName)),                                                                    // 实例主机名
		UserData:           tea.String(utils.GenerateUserData()),                                                                                // 实例初始化脚本 base64 加密
		PrivateIpAddress:   tea.String(vm.PrivateIp),                                                                                            // 实例内网IP地址
		Password:           tea.String(vm.Password),                                                                                             // 实例密码
		InstanceType:       tea.String(utils.GenerateInstanceTypeBySpec(vm.Spec)),                                                               // 实例类型
		RegionId:           tea.String(vm.Region),                                                                                               // 实例地域：cn-shanghai
		ImageId:            tea.String(config.GlobalConfig.Aliyun.ImageId),                                                                      // 系统镜像
		InstanceChargeType: tea.String("PrePaid"),                                                                                               // 包年包月
		PeriodUnit:         tea.String("Month"),                                                                                                 // 包年包月单位：月
		Period:             tea.Int32(1),                                                                                                        // 包年包月单位时长: 1个月
		Amount:             tea.Int32(1),                                                                                                        // 实例数量
		SecurityGroupIds:   tea.StringSlice([]string{config.GlobalConfig.Aliyun.SecurityGroupId1, config.GlobalConfig.Aliyun.SecurityGroupId2}), // 安全组ID
		VSwitchId:          tea.String(config.GlobalConfig.Aliyun.VSwitchId),                                                                    // 交换机ID
		SystemDisk: &ecs20140526.RunInstancesRequestSystemDisk{
			Category: tea.String("cloud_efficiency"),                    // 磁盘类型
			Size:     tea.String(utils.GenerateDiskSizeBySpec(vm.Spec)), // 磁盘大小
		},
		/*
			包年包月 配置
			InstanceChargeType: tea.String("PrePaid"),                                  // 包年包月
			PeriodUnit:         tea.String("Month"),                                    // 包年包月单位：月
			Period:             tea.Int32(1),                                           // 包年包月单位时长: 1个月
			按量付费 配置
			InstanceChargeType: tea.String("PostPaid"),									// 按量付费
			SpotStrategy:       tea.String("NoSpot"),									// 正常按量付费
			启动脚本
			UserData:           tea.String(utils.GenerateUserData()),                   // 实例启动初始化脚本 base64 加密
		*/
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
	return svc.VmRepo.CreateVm(vm)
}

func (svc *VmService) FetchUserByVm(id int) ([]*model.User, error) {
	return svc.VmRepo.GetUserByVm(id)
}
