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

func (repo *UserService) Login(account, password string) (interface{}, interface{}, error) {

	if err := LdapClient.Bind(fmt.Sprintf("cn=%s,ou=%s,dc=%s,dc=%s",
		account,
		config.GlobalConfig.OpenLdap.Ou,
		config.GlobalConfig.OpenLdap.Dc1,
		config.GlobalConfig.OpenLdap.Dc2), password); err != nil {
		return nil, nil, err
	}

	token, err := utils.GenerateToken(account)
	if err != nil {
		return nil, nil, err
	}

	encryptToken, err := utils.EncryptAESGCM(token)
	if err != nil {
		return nil, nil, err
	}

	rows, err := repo.UserRepo.UpdateTokenByAccount(account, encryptToken)
	if err != nil {
		return nil, nil, err
	}
	if rows == 0 {
		return nil, nil, errors.New("数据库未更改")
	}

	result, err := repo.UserRepo.GetUsers(account)
	if err != nil {
		return nil, nil, err
	}

	return token, result, nil
}

func (repo *UserService) UserPermission(account string) (interface{}, error) {
	return repo.UserRepo.GetPermissions(account)
}

func (repo *UserService) List() ([]*model.User, error) {
	return repo.UserRepo.ListUsers()
}

func (repo *UserService) UserRoles(account string) (interface{}, error) {
	return repo.UserRepo.GetRoles(account)
}
