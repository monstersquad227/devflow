package service

import (
	"devflow/repository"
	"devflow/utils"
	"errors"
	"fmt"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (s *UserService) UserLogin(account, password string) (interface{}, interface{}, error) {

	if err := LdapClient.Bind(fmt.Sprintf("cn=%s,ou=technology,dc=chengdd,dc=com", account), password); err != nil {
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

	rows, err := s.Repo.UpdateTokenByAccount(account, encryptToken)
	if err != nil {
		return nil, nil, err
	}
	if rows == 0 {
		return nil, nil, errors.New("数据库未更改")
	}

	result, err := s.Repo.GetUsers(account)
	if err != nil {
		return nil, nil, err
	}

	return token, result, nil
}

func (s *UserService) UserPermission(account string) (interface{}, error) {
	return s.Repo.GetPermissions(account)
}

func (s *UserService) UserRoles(account string) (interface{}, error) {
	return s.Repo.GetRoles(account)
}
