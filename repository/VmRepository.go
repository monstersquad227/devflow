package repository

import (
	"database/sql"
	"devflow/model"
	"fmt"
)

type VmRepository struct{}

func (repo *VmRepository) ListVms(pageNumber, pageSize int) ([]*model.Vm, error) {
	query := "SELECT id, instance_id, instance_name, private_ip, public_ip, spec, application, region, cloud_provider, os, created_at, updated_at " +
		"FROM vm WHERE is_deleted = 0 " +
		"ORDER BY " +
		"CASE cloud_provider " +
		"WHEN 'local' THEN 1 " +
		"WHEN 'aliyun' THEN 2 " +
		"WHEN 'huawei' THEN 3 " +
		"WHEN 'tencent' THEN 4 " +
		"WHEN 'aws' THEN 5 " +
		"ELSE 99 " +
		"END ASC, " +
		"CASE spec " +
		"WHEN 'small' THEN 1 " +
		"WHEN 'medium' THEN 2 " +
		"WHEN 'large' THEN 3 " +
		"WHEN 'xlarge' THEN 4 " +
		"WHEN '2xlarge' THEN 5 " +
		"WHEN 'ultra' THEN 6 " +
		"ELSE 99 " +
		"END ASC " +
		"LIMIT ? OFFSET ? "
	rows, err := MysqlClient.Query(query, pageSize, (pageNumber-1)*pageSize)
	if err != nil {
		return nil, err
	}
	data := make([]*model.Vm, 0)
	for rows.Next() {
		obj := &model.Vm{}
		if err = rows.Scan(&obj.Id, &obj.InstanceId, &obj.InstanceName, &obj.PrivateIp, &obj.PublicIp, &obj.Spec,
			&obj.Application, &obj.Region, &obj.CloudProvider, &obj.Os, &obj.CreatedAt, &obj.UpdatedAt); err != nil {
			fmt.Println(err)
			return nil, err
		}
		data = append(data, obj)
	}
	return data, nil
}

func (repo *VmRepository) CountVms() (int, error) {
	query := "SELECT count(id) " +
		"FROM vm WHERE is_deleted = 0 "
	var count int
	if err := MysqlClient.QueryRow(query).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (repo *VmRepository) CreateVm(vm *model.Vm) (int64, error) {
	query := "INSERT " +
		"INTO vm(instance_id, instance_name, private_ip, public_ip, spec, application, region, cloud_provider, os, password) " +
		"VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := MysqlClient.Exec(query, vm.InstanceId, vm.InstanceName, vm.PrivateIp, vm.PublicIp, vm.Spec, vm.Application, vm.Region, vm.CloudProvider, vm.Os, vm.Password)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (repo *VmRepository) UpdateVm(vm *model.Vm) (int64, error) {
	query := "UPDATE vm " +
		"SET instance_name = ?, private_ip = ?, public_ip = ?, spec = ?, application = ?, region = ?, cloud_provider = ?, os = ? " +
		"WHERE id = ?"
	result, err := MysqlClient.Exec(query, vm.InstanceName, vm.PrivateIp, vm.PublicIp, vm.Spec, vm.Application, vm.Region, vm.CloudProvider, vm.Os, vm.Id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (repo *VmRepository) DeleteVm(id int) (int64, error) {
	query := "UPDATE vm " +
		"SET is_deleted = 1 WHERE id = ?"
	result, err := MysqlClient.Exec(query, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (repo *VmRepository) GetVmPasswordById(id int) (string, error) {
	var password string
	query := "SELECT password " +
		"FROM vm WHERE id = ?"
	if err := MysqlClient.QueryRow(query, id).Scan(&password); err != nil {
		return "", err
	}
	return password, nil
}

func (repo *VmRepository) GetVmsByApplication(application string) (interface{}, error) {
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

func (repo *VmRepository) GetCloudProviderById(id int) (string, error) {
	var cloud string
	query := "SELECT cloud_provider " +
		"FROM vm WHERE id = ?"
	err := MysqlClient.QueryRow(query, id).Scan(&cloud)
	if err != nil {
		return "", err
	}
	return cloud, nil
}

func (repo *VmRepository) GetInstanceIDById(id int) (string, error) {
	var instanceId string
	query := "SELECT instance_id " +
		"FROM vm WHERE id = ?"
	err := MysqlClient.QueryRow(query, id).Scan(&instanceId)
	if err != nil {
		return "", err
	}
	return instanceId, nil
}

func (repo *VmRepository) GetRegionById(id int) (string, error) {
	var region string
	query := "SELECT region " +
		"FROM vm WHERE id = ?"
	err := MysqlClient.QueryRow(query, id).Scan(&region)
	if err != nil {
		return "", err
	}
	return region, nil
}

func (repo *VmRepository) GetUserByVm(id int) ([]*model.User, error) {
	query := "SELECT " +
		"    u.id, " +
		"    u.name " +
		"FROM " +
		"    user u " +
		"INNER JOIN " +
		"    bastion b ON u.id = b.user_id " +
		"WHERE " +
		"    b.vm_id = ?;"
	rows, err := MysqlClient.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			return
		}
	}(rows)
	data := make([]*model.User, 0)
	for rows.Next() {
		obj := &model.User{}
		if err = rows.Scan(&obj.ID, &obj.Name); err != nil {
			return nil, err
		}
		data = append(data, obj)
	}
	return data, nil
}

func (repo *VmRepository) UpdateUserByVm(vmID int, userIDs []int) (int64, error) {
	deleteSql := "DELETE " +
		"FROM bastion " +
		"WHERE " +
		"	vm_id = ?"
	insertSql := "INSERT " +
		"INTO bastion(vm_id, user_id) " +
		"VALUES(?, ?);"

	tx, err := MysqlClient.Begin()
	if err != nil {
		return 0, err
	}

	_, err = tx.Exec(deleteSql, vmID)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	if len(userIDs) == 0 {
		return 0, tx.Commit()
	}

	stmt, err := tx.Prepare(insertSql)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			return
		}
	}(stmt)

	var count int64
	for _, userID := range userIDs {
		res, err := stmt.Exec(vmID, userID)
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return 0, err
			}
			return 0, err
		}
		rows, _ := res.RowsAffected()
		count += rows
	}
	if err := tx.Commit(); err != nil {
		return 0, err
	}
	
	return count, nil
}
