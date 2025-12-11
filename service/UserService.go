package service

import (
	"devflow/config"
	"devflow/model"
	"devflow/repository"
	"devflow/utils"
	"errors"
	"fmt"
)

type UserService struct {
	UserRepo repository.UserRepositoryInterface
}

func NewUserService(repo repository.UserRepositoryInterface) *UserService {
	return &UserService{
		UserRepo: repo,
	}
}

func (repo *UserService) Login(account, password string) (*model.LoginResponse, error) {

	if err := LdapClient.Bind(fmt.Sprintf("cn=%s,ou=%s,dc=%s,dc=%s",
		account,
		config.GlobalConfig.OpenLdap.Ou,
		config.GlobalConfig.OpenLdap.Dc1,
		config.GlobalConfig.OpenLdap.Dc2), password); err != nil {
		return nil, err
	}

	token, err := utils.GenerateToken(account)
	if err != nil {
		return nil, err
	}

	encryptToken, err := utils.EncryptAESGCM(token)
	if err != nil {
		return nil, err
	}

	rows, err := repo.UserRepo.UpdateTokenByAccount(account, encryptToken)
	if err != nil {
		return nil, err
	}
	if rows == 0 {
		return nil, errors.New("数据库未更改")
	}

	users, err := repo.UserRepo.GetUsers(account)
	if err != nil {
		return nil, err
	}

	roles, err := repo.UserRepo.GetRoles(int64(users.ID))
	if err != nil {
		return nil, err
	}

	permissions, err := repo.UserRepo.GetPermissions(int64(users.ID))
	if err != nil {
		return nil, err
	}

	menus, err := repo.UserRepo.GetMenus(int64(users.ID))
	if err != nil {
		return nil, err
	}

	result := &model.LoginResponse{
		Token:       token,
		User:        users,
		Roles:       roles,
		Permissions: permissions,
		Menus:       menus,
	}

	return result, nil
}

//func (repo *UserService) UserPermission(account string) (interface{}, error) {
//	return repo.UserRepo.GetPermissions(account)
//}

func (repo *UserService) List() ([]*model.User, error) {
	return repo.UserRepo.ListUsers()
}
