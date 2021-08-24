package global

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type App struct {
	Name string `json:"name"`
	Port string `json:"port"`
	Mode string `json:"mode"`
}

type Config struct {
	App
}

func (c *Config) Init() {
	wd, err := os.Getwd()
	if err != nil {
		logrus.Error(err)
	}
	logrus.Println(wd)
	v := viper.New()
	v.SetConfigName("local")
	v.AddConfigPath(wd + "/config") //文件所在的目录路径
	v.SetConfigType("yml")          //这里是文件格式类型
	err = v.ReadInConfig()
	if err != nil {
		logrus.Error(err)
		return
	}
	configs := v.AllSettings()
	for k, val := range configs {
		v.SetDefault(k, val)
	}
	logrus.Println(configs)
	err = v.Unmarshal(c) //反序列化至结构体
	if err != nil {
		logrus.Fatal("读取配置错误：", err)
		return
	}
}
