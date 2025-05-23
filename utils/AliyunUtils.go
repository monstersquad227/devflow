package utils

import (
	"encoding/base64"
	"strings"
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
