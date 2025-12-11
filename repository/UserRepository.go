package repository

import (
	"database/sql"
	"devflow/model"
	"encoding/json"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

/*
UpdateTokenByAccount 通过 account 更新 token 字段
*/

func (r *UserRepository) UpdateTokenByAccount(account, token string) (int64, error) {
	query := "UPDATE user " +
		"SET token = ? " +
		"WHERE account = ?"

	result, err := MysqlClient.Exec(query, token, account)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (r *UserRepository) ListUsers() ([]*model.User, error) {
	query := "SELECT id, name " +
		"FROM user"
	rows, err := MysqlClient.Query(query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)
	data := make([]*model.User, 0)
	for rows.Next() {
		obj := &model.User{}
		err := rows.Scan(&obj.ID, &obj.Name)
		if err != nil {
			return nil, err
		}
		data = append(data, obj)
	}
	return data, nil
}

func (r *UserRepository) GetUsers(account string) (*model.User, error) {
	var roles, permissions string
	var obj model.User
	query := "SELECT id, account, name, email, mobile, roles, permissions, created_at, updated_at " +
		"FROM user WHERE account = ? AND deleted = 0 "

	err := MysqlClient.QueryRow(query, account).Scan(&obj.ID, &obj.Account, &obj.Name, &obj.Email, &obj.Mobile, &roles, &permissions, &obj.CreatedAt, &obj.UpdatedAt)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal([]byte(roles), &obj.Roles); err != nil {
		return nil, err
	}
	if err = json.Unmarshal([]byte(permissions), &obj.Permissions); err != nil {
		return nil, err
	}
	return &obj, err
}

/*
GetPermissions 通过 account 获取 permissions 字段
*/

func (r *UserRepository) GetPermissions(id int64) ([]string, error) {
	query := "SELECT  " +
		"    p.permission_code " +
		"FROM permission p " +
		"INNER JOIN role_permission rp ON p.id = rp.permission_id " +
		"INNER JOIN user_role ur ON rp.role_id = ur.role_id " +
		"WHERE ur.user_id = ? " +
		"    AND p.status = 1 " +
		"    AND p.deleted = 0 " +
		"ORDER BY p.id;"
	var data []string
	rows, err := MysqlClient.Query(query, id)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var obj string
		err := rows.Scan(&obj)
		if err != nil {
			return nil, err
		}
		data = append(data, obj)
	}
	return data, nil
}

/*
GetRoles 通过 account 获取 roles 字段
*/

func (r *UserRepository) GetRoles(id int64) ([]*model.Role, error) {
	query := "SELECT " +
		"    r.id, " +
		"    r.role_code AS roleCode, " +
		"    r.role_name AS roleName " +
		"FROM role r " +
		"INNER JOIN user_role ur ON r.id = ur.role_id " +
		"WHERE ur.user_id = ? " +
		"    AND r.status = 1 " +
		"    AND r.deleted = 0;"
	data := make([]*model.Role, 0)
	rows, err := MysqlClient.Query(query, id)

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		obj := &model.Role{}
		err = rows.Scan(&obj.ID, &obj.RoleCode, &obj.RoleName)
		if err != nil {
			return nil, err
		}
		data = append(data, obj)
	}
	return data, nil
}

/*
GetRoles 通过 account 获取 menus
*/

func (r *UserRepository) GetMenus(id int64) ([]*model.Menu, error) {
	query := "SELECT " +
		"    p.id, " +
		"    p.permission_code AS permissionCode, " +
		"    p.permission_name AS permissionName, " +
		"    p.path " +
		"FROM permission p " +
		"INNER JOIN role_permission rp ON p.id = rp.permission_id " +
		"INNER JOIN user_role ur ON rp.role_id = ur.role_id " +
		"WHERE ur.user_id = ? " +
		"    AND p.resource_type = 'menu' " +
		"    AND p.status = 1 " +
		"  AND p.deleted = 0 " +
		"ORDER BY p.parent_id, p.id;"
	data := make([]*model.Menu, 0)
	rows, err := MysqlClient.Query(query, id)

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		obj := &model.Menu{}
		err = rows.Scan(&obj.ID, &obj.PermissionCode, &obj.PermissionName, &obj.Path)
		if err != nil {
			return nil, err
		}
		data = append(data, obj)
	}
	return data, nil
}
