package config

import (
	"book_manager/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	//使用viper读取配置文件
	v := viper.New()
	v.SetConfigFile("config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	//读取配置文件中的配置到全局变量中

	err = v.Unmarshal(&global.Conf)
	if err != nil {
		panic(err)
	}
}
