package utils

import (
	"bytes"
	"devflow/config"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	url2 "net/url"
	"strings"
	"time"
)

/*
DeleteContainer 删除容器，通过容器名删除
*/

func DeleteContainer(addr, name string) {
	url := fmt.Sprintf("http://%s:%s/containers/%s?force=1", addr, config.GlobalConfig.Docker.RemotePort, name)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		logrus.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		logrus.Println(err)
		return
	}
}

/*
CreateContainer 创建容器

Body:

{
	"User": "root",
	"Image": "%s",								// 镜像
	"Env": [									// 环境变量
		"SERVER_HOST=%s",
		"SERVER_ENV=%s",
		"SERVER_PORT=8080",
		"ASPNETCORE_ENVIRONMENT=%s"
	],
	"HostConfig": {
		"Binds": [
			"/data/logs/:/data/logs",			// 日志目录
			"/etc/localtime:/etc/localtime:ro"	// 时间
		],
		"NetworkMode": "host",					// 使用宿主机IP
		"Privileged": true,
		"RestartPolicy": {
			"MaximumRetryCount": 5,
			"Name": "on-failure"
		}
	}
}
*/

func CreateContainer(addr, name, image string) {
	url := fmt.Sprintf("http://%s:%s/containers/create?name=%s", addr, config.GlobalConfig.Docker.RemotePort, name)
	body := fmt.Sprintf(`
		{
			"User": "root",
			"Image": "%s",
			"Env": [
				"SERVER_HOST=%s",
				"SERVER_ENV=%s",
				"SERVER_PORT=8080",
				"ASPNETCORE_ENVIRONMENT=%s"
			],
			"HostConfig": {
				"Binds": [
					"/data/logs/:/data/logs",
					"/etc/localtime:/etc/localtime:ro"
				],
				"NetworkMode": "host",
				"Privileged": true,
				"RestartPolicy": {
					"MaximumRetryCount": 5,
					"Name": "on-failure"
				}
			}
		}
	`, image, addr, "dev", "dev")
	bodyByte := []byte(body)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyByte))
	if err != nil {
		logrus.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		logrus.Println(err)
		return
	}
}

/*
StartContainer 启动容器
*/

func StartContainer(addr, name string) {
	url := fmt.Sprintf("http://%s:%s/containers/%s/start", addr, config.GlobalConfig.Docker.RemotePort, name)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		logrus.Println(err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		logrus.Println(err)
		return
	}
}

/*
PullImage 拉取镜像
*/

func PullImage(addr, image string) {
	imageSplit := strings.Split(image, ":")
	url := fmt.Sprintf(`http://%s:%s/images/create?fromImage=%s&tag=%s`, addr, config.GlobalConfig.Docker.RemotePort, url2.QueryEscape(imageSplit[0]), imageSplit[1])
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		logrus.Println(err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		logrus.Println(err)
		return
	}
}

func GenerateImage(env, name, tag string) string {
	raw := config.GlobalConfig.Harbor.URL + "/" + env + "/" + name + ":" + tag
	parts := strings.SplitN(raw, "://", 2)
	if len(parts) == 2 {
		return parts[1]
	}
	return raw
}
