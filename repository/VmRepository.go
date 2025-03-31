package repository

import (
	"devflow/model"
	"fmt"
)

type VmRepository struct{}

func (receiver *VmRepository) GetVms(pageNumber, pageSize int) ([]*model.Vm, error) {
	query := "SELECT id, instance_id, instance_name, private_ip, public_ip, spec, region, cloud_provider, os, created_at, updated_at " +
		"FROM vm WHERE is_deleted = 0 LIMIT ? OFFSET ? "
	rows, err := MysqlClient.Query(query, pageSize, (pageNumber-1)*pageSize)
	if err != nil {
		return nil, err
	}
	data := make([]*model.Vm, 0)
	for rows.Next() {
		obj := &model.Vm{}
		if err = rows.Scan(&obj.Id, &obj.InstanceId, &obj.InstanceName, &obj.PrivateIp, &obj.PublicIp, &obj.Spec,
			&obj.Region, &obj.CloudProvider, &obj.Os, &obj.CreatedAt, &obj.UpdatedAt); err != nil {
			fmt.Println(err)
			return nil, err
		}
		data = append(data, obj)
	}
	return data, nil
}

func (receiver *VmRepository) CreateVm(vm model.Vm) (int64, error) {
	query := "INSERT " +
		"INTO vm(instance_id, instance_name, private_ip, public_ip, spec, region, cloud_provider, os, password) " +
		"VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := MysqlClient.Exec(query, vm.InstanceId, vm.InstanceName, vm.PrivateIp, vm.PublicIp, vm.Spec, vm.Region, vm.CloudProvider, vm.Os, vm.Password)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (receiver *VmRepository) DeleteVm(id int) (int64, error) {
	query := "UPDATE vm " +
		"SET is_deleted = 1 WHERE id = ?"
	result, err := MysqlClient.Exec(query, id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (receiver *VmRepository) GetVmsCount() (int, error) {
	query := "SELECT count(id) " +
		"FROM vm WHERE is_deleted = 0 "
	var count int
	if err := MysqlClient.QueryRow(query).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (receiver *VmRepository) GetVmsByApplication(application string) (interface{}, error) {
	query := "SELECT private_ip, public_ip, instance_name " +
		"FROM vm WHERE application = ?"
	rows, err := MysqlClient.Query(query, application)
	if err != nil {
		return nil, err
	}
	data := make([]*model.Vm, 0)
	for rows.Next() {
		obj := &model.Vm{}
		if err = rows.Scan(&obj.PrivateIp, &obj.PublicIp, &obj.InstanceName); err != nil {
			return nil, err
		}
		data = append(data, obj)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data, nil
}
