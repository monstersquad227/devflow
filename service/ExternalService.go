package service

import (
	"context"
	"devflow/config"
	aliyunopenapi "github.com/alibabacloud-go/darabonba-openapi/client"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/bndr/gojenkins"
	"github.com/go-ldap/ldap/v3"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/region"
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

func NewAliyunClient() (*ecs20140526.Client, error) {
	cfg := &aliyunopenapi.Config{
		AccessKeyId:     tea.String(config.GlobalConfig.Aliyun.AccessKey),
		AccessKeySecret: tea.String(config.GlobalConfig.Aliyun.SecretKey),
	}
	cfg.Endpoint = tea.String("ecs.cn-shanghai.aliyuncs.com")
	client, err := ecs20140526.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewHuaweiClient() *ecs.EcsClient {
	return ecs.NewEcsClient(
		ecs.EcsClientBuilder().
			WithRegion(region.ValueOf("cn-east-3")).
			WithCredential(basic.NewCredentialsBuilder().
				WithAk(config.GlobalConfig.Huawei.AccessKey).
				WithSk(config.GlobalConfig.Huawei.SecretKey).
				Build()).
			Build())
}
