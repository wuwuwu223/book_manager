package global

import (
	"gorm.io/gorm"
)

var Db *gorm.DB
var Conf *Config

type Config struct {
	Mysql struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Dbname   string `mapstructure:"dbname"`
	} `mapstructure:"mysql"`
	ListenPort string `mapstructure:"listen_port"`
}
