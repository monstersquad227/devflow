package main

import (
	"crypto/tls"
	"crypto/x509"
	"devflow/config"
	"devflow/middleware"
	"devflow/repository"
	v1 "devflow/routes/v1"
	"devflow/service"
	"github.com/gin-gonic/gin"
	pb "github.com/monstersquad227/flowedge-proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	// 加载配置文件
	config.LoadConfig()

	// 初始化
	repository.InitMysql()
	repository.InitRedis()
	service.InitGitlab()
	service.InitJenkins()
	service.InitOpenLdap()

	Application := gin.New()

	// 加载中间件
	Application.Use(middleware.Cors())                    // 跨域
	Application.Use(middleware.Jwt())                     // Jwt
	Application.Use(middleware.Logger(middleware.HttpIn)) // 日志
	Application.Use(middleware.RecoveryWithLogger())      // Panic

	baseRouter := Application.Group(config.GlobalConfig.Application.Name)

	// 注册路由
	v1.BasicRegister(baseRouter)
	v1.UserRegister(baseRouter)
	v1.ProjectRegister(baseRouter)
	v1.VmRegister(baseRouter)
	v1.SettingRegister(baseRouter)
	v1.FlowedgeRegister(baseRouter)

	// 启动服务
	go func() {
		err := Application.Run(":" + config.GlobalConfig.Application.Port)
		if err != nil {
			log.Fatal(config.GlobalConfig.Application.Name+" 启动失败 : ", err)
		}
	}()

	// 设置 TLS 和启动 gRPC 服务
	cert, err := tls.LoadX509KeyPair("./certs/server.crt", "./certs/server.key")
	if err != nil {
		log.Fatal("server 证书文件加载失败: ", err)
	}
	caCert, err := ioutil.ReadFile("./certs/ca.crt")
	if err != nil {
		log.Fatal("ca 证书文件加载失败: ", err)
	}
	caPool := x509.NewCertPool()
	caPool.AppendCertsFromPEM(caCert)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    caPool,
	})

	s := grpc.NewServer(grpc.Creds(creds))
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("grpc 监听失败: ", err)
	}
	pb.RegisterFlowEdgeServer(s, service.GlobalFlowEdgeServer)

	err = s.Serve(lis)
	if err != nil {
		log.Fatal("grpc 启动失败: ", err)
	}

}
