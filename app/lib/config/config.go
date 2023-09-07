package config

import (
	"github.com/spf13/viper"
	"os"
)

func Init() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	viper.Reset()
	//设置配置文件目录
	viper.AddConfigPath(path)
	//设置配置文件名称
	viper.SetConfigName("env")
	//设置配置文件类型
	viper.SetConfigType("yaml")
	//设置监听配置文件修改变化
	viper.WatchConfig()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
