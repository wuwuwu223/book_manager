package main

import (
	"book_manager/config"
	"book_manager/controller"
	"book_manager/dao"
	"book_manager/global"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置文件
	{
		config.LoadConfig()
		global.Db = dao.ConnectDb()
	}
	// 初始化路由
	r := gin.Default()
	r.Use(Cors())
	x := r.Group("/excise")

	// 读者
	{
		x.POST("/login", controller.ReaderLogin)
		x.POST("/getAllReaders", controller.GetAllReaders)
		x.POST("/removeReader", controller.DelReader)
		x.POST("/addReader", controller.AddReader)
		x.POST("/updateReader", controller.UpdateReader)
	}

	// 图书
	{
		x.POST("/getAllAlbums", controller.GetAllAlbums)
		x.POST("/removeAlbum", controller.DelAlbum)
		x.POST("/addAlbum", controller.AddAlbum)
		x.POST("/addSubAlbum", controller.AddSubAlbum)
	}

	// 借阅
	{
		x.POST("/getAllBorrowRecords", controller.GetAllBorrowRecords)
		x.POST("/reback", controller.DelBorrowRecord)
		x.POST("/borrow", controller.AddBorrowRecord)
	}
	r.Run(fmt.Sprintf(":%s", global.Conf.ListenPort))
}

// Cors 跨域中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
