package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Application struct {
		Name string `mapstructure:"name"`
		Port string `mapstructure:"port"`
	}
	OpenLdap struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
		Ou   string `mapstructure:"ou"`
		Dc1  string `mapstructure:"dc1"`
		Dc2  string `mapstructure:"dc2"`
	}
	Jwt struct {
		Expire int64  `mapstructure:"expire"`
		Secret string `mapstructure:"secret"`
	}
	Aes struct {
		Secret string `mapstructure:"secret"`
	}
	Mysql struct {
		Username  string `mapstructure:"username"`
		Password  string `mapstructure:"password"`
		Addr      string `mapstructure:"addr"`
		Port      string `mapstructure:"port"`
		Databases string `mapstructure:"databases"`
		Charset   string `mapstructure:"charset"`
	}
	Redis struct {
		Addr string `mapstructure:"addr"`
		Port string `mapstructure:"port"`
		Db   int    `mapstructure:"db"`
	}
	Gitlab struct {
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Addr     string `mapstructure:"addr"`
	}
	Jenkins struct {
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Addr     string `mapstructure:"addr"`
		Token    string `mapstructure:"token"`
	}
	Harbor struct {
		URL      string `mapstructure:"url"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	}
	Docker struct {
		RemotePort string `mapstructure:"remote_port"`
	}
	Aliyun struct {
		AccessKey        string `mapstructure:"access_key"`
		SecretKey        string `mapstructure:"secret_key"`
		ImageId          string `mapstructure:"image_id"`
		VSwitchId        string `mapstructure:"vswitch_id"`
		SecurityGroupId1 string `mapstructure:"security_group_id1"`
		SecurityGroupId2 string `mapstructure:"security_group_id2"`
	}
	Huawei struct {
		AccessKey string `mapstructure:"access_key"`
		SecretKey string `mapstructure:"secret_key"`
	}
}

var GlobalConfig Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.SetConfigType("ini")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: , %v", err)
	}

	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		log.Fatalf("解析配置失败: , %v", err)
	}
}
