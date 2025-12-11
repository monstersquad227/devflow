package service

import (
	"devflow/config"
	"devflow/model"
	"devflow/repository"
	"devflow/utils"
	"errors"
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

type UserService struct {
	UserRepo repository.UserRepositoryInterface
}

func NewUserService(repo repository.UserRepositoryInterface) *UserService {
	return &UserService{
		UserRepo: repo,
	}
}

func (svc *UserService) Login(account, password string) (*model.LoginResponse, error) {

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

	rows, err := svc.UserRepo.UpdateTokenByAccount(account, encryptToken)
	if err != nil {
		return nil, err
	}
	if rows == 0 {
		return nil, errors.New("数据库未更改")
	}

	users, err := svc.UserRepo.GetUsers(account)
	if err != nil {
		return nil, err
	}

	roles, err := svc.UserRepo.GetRoles(int64(users.ID))
	if err != nil {
		return nil, err
	}

	permissions, err := svc.UserRepo.GetPermissions(int64(users.ID))
	if err != nil {
		return nil, err
	}

	menus, err := svc.UserRepo.GetMenus(int64(users.ID))
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

func (svc *UserService) List() ([]*model.User, error) {
	return svc.UserRepo.ListUsers()
}

func (svc *UserService) PasswordChange(req *model.PasswordRequest) (*model.PasswordResponse, error) {
	resp := &model.PasswordResponse{}

	if err := LdapClient.Bind(fmt.Sprintf("cn=%s,ou=%s,dc=%s,dc=%s",
		req.Account,
		config.GlobalConfig.OpenLdap.Ou,
		config.GlobalConfig.OpenLdap.Dc1,
		config.GlobalConfig.OpenLdap.Dc2), req.Password); err != nil {
		resp.Message = "原密码错误"
		return resp, err
	}

	_, err := LdapClient.PasswordModify(&ldap.PasswordModifyRequest{
		UserIdentity: fmt.Sprintf("cn=%s,ou=%s,dc=%s,dc=%s",
			req.Account,
			config.GlobalConfig.OpenLdap.Ou,
			config.GlobalConfig.OpenLdap.Dc1,
			config.GlobalConfig.OpenLdap.Dc2),
		OldPassword: req.Password,
		NewPassword: req.NewPassword,
	})
	if err != nil {
		resp.Message = "修改密码失败"
		return resp, err
	}

	resp.Message = "修改密码成功"
	return resp, nil
}
