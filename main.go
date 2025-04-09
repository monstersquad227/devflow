package main

import (
	"devflow/config"
	"devflow/middleware"
	"devflow/repository"
	v1 "devflow/routes/v1"
	"devflow/service"
	"github.com/gin-gonic/gin"
	"log"
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

	v1.KongRegister(baseRouter)

	// 启动服务
	err := Application.Run(":" + config.GlobalConfig.Application.Port)
	if err != nil {
		log.Fatal(config.GlobalConfig.Application.Name+" 启动失败 : ", err)
	}
}
