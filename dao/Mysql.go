package dao

import (
	"book_manager/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func ConnectDb() *gorm.DB {
	//gorm连接mysql
	dest := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		global.Conf.Mysql.Username,
		global.Conf.Mysql.Password,
		global.Conf.Mysql.Host,
		global.Conf.Mysql.Port,
		global.Conf.Mysql.Dbname)

	Db, err := gorm.Open(mysql.Open(dest), &gorm.Config{})
	if err != nil {
		log.Println("连接数据库失败")
		panic(err)
	}
	return Db
}
