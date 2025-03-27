package service

import (
	"context"
	"devflow/config"
	"github.com/bndr/gojenkins"
	"github.com/go-ldap/ldap/v3"
	"github.com/xanzy/go-gitlab"
	"log"
)

var (
	GitlabClient  *gitlab.Client
	JenkinsClient *gojenkins.Jenkins
	LdapClient    *ldap.Conn
)

func InitGitlab() {
	var err error
	GitlabClient, err = gitlab.NewBasicAuthClient(
		config.GlobalConfig.Gitlab.Username,
		config.GlobalConfig.Gitlab.Password,
		gitlab.WithBaseURL(config.GlobalConfig.Gitlab.Addr))
	if err != nil {
		log.Fatalf("初始化 GitLab 客户端失败: %v", err)
	}
	log.Println("GitLab 客户端初始化成功")
}

func InitJenkins() {
	JenkinsClient = gojenkins.CreateJenkins(nil,
		config.GlobalConfig.Jenkins.Addr,
		config.GlobalConfig.Jenkins.Username,
		config.GlobalConfig.Jenkins.Password,
	)
	_, err := JenkinsClient.Init(context.Background())
	if err != nil {
		log.Fatalf("初始化 Jenkins 客户端失败: %v", err)
	}
	log.Println("Jenkins 客户端初始化成功")
}

func InitOpenLdap() {
	l, err := ldap.Dial("tcp", config.GlobalConfig.OpenLdap.Host+":"+config.GlobalConfig.OpenLdap.Port)
	if err != nil {
		log.Fatalf("初始化 OpenLdap 客户端失败: %v", err)
	}
	LdapClient = l
	log.Println("OpenLdap 客户端初始化成功")
}
