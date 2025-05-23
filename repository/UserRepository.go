package repository

import (
	"devflow/model"
	"encoding/json"
)

type UserRepository struct{}

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

func (r *UserRepository) GetUsers(account string) (interface{}, error) {
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

func (r *UserRepository) GetPermissions(account string) (interface{}, error) {
	query := "SELECT permissions " +
		"FROM user WHERE account = ?"
	var str string
	if err := MysqlClient.QueryRow(query, account).Scan(&str); err != nil {
		return nil, err
	}
	var result []string
	if err := json.Unmarshal([]byte(str), &result); err != nil {
		return nil, err
	}
	return result, nil
}

/*
GetRoles 通过 account 获取 roles 字段
*/

func (r *UserRepository) GetRoles(account string) (interface{}, error) {
	query := "SELECT roles " +
		"FROM user WHERE account = ?"
	var str string
	if err := MysqlClient.QueryRow(query, account).Scan(&str); err != nil {
		return nil, err
	}
	var result []string
	if err := json.Unmarshal([]byte(str), &result); err != nil {
		return nil, err
	}
	return result, nil
}
