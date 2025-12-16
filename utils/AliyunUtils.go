package utils

import (
	"devflow/config"
	"encoding/base64"
	"encoding/json"
	"fmt"
	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/gogo/protobuf/proto"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

func GenerateUserData() string {
	return base64.StdEncoding.EncodeToString([]byte("#!/bin/bash\ncurl 192.168.200.1/initialization.sh | bash"))
}

func GenerateDiskSizeBySpec(spec string) string {
	switch spec {
	case "small":
		return "40"
	case "medium":
		return "80"
	case "large":
		return "100"
	case "extraLarge":
		return "150"
	default:
		return "40"
	}
}

func GenerateInstanceTypeBySpec(spec string) string {
	switch spec {
	case "small":
		return "ecs.t5-lc1m2.small"
	case "medium":
		return "ecs.t5-lc1m2.large"
	case "large":
		return "ecs.t5-c1m2.xlarge"
	case "xlarge":
		return "ecs.t5-c1m2.2xlarge"
	default:
		return "ecs.t5-lc1m2.small"
	}
}

func StringToLower(val string) string {
	return strings.ToLower(val)
}

func SendToSls(fields logrus.Fields) {
	slsClient := sls.CreateNormalInterface(
		config.GlobalConfig.Aliyun.SlsEndpoint,
		config.GlobalConfig.Aliyun.AccessKey,
		config.GlobalConfig.Aliyun.SecretKey,
		"")

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("发送SLS日志失败: %v\n", r)
			}
		}()

		// 构造日志内容
		contents := make([]*sls.LogContent, 0, len(fields))
		for key, value := range fields {
			// 将值转换为字符串
			var valueStr string
			switch v := value.(type) {
			case string:
				valueStr = v
			case int, int64, float64:
				valueStr = fmt.Sprintf("%v", v)
			default:
				// 复杂类型序列化为JSON
				if jsonBytes, err := json.Marshal(v); err == nil {
					valueStr = string(jsonBytes)
				} else {
					valueStr = fmt.Sprintf("%v", v)
				}
			}
			contents = append(contents, &sls.LogContent{
				Key:   proto.String(key),
				Value: proto.String(valueStr),
			})
		}
		// 添加时间戳和主机名
		hostname, _ := os.Hostname()
		contents = append(contents, &sls.LogContent{
			Key:   proto.String("hostname"),
			Value: proto.String(hostname),
		})

		// 构造日志
		log := &sls.Log{
			Time:     proto.Uint32(uint32(time.Now().Unix())),
			Contents: contents,
		}

		// 构造LogGroup
		logGroup := &sls.LogGroup{
			Topic:  proto.String(""),
			Source: proto.String(hostname),
			Logs:   []*sls.Log{log},
		}

		// 发送日志
		if err := slsClient.PutLogs(
			config.GlobalConfig.Aliyun.SlsProject,
			config.GlobalConfig.Aliyun.SlsLogstore,
			logGroup,
		); err != nil {
			fmt.Printf("发送SLS日志失败: %v\n", err)
		}
	}()
}
