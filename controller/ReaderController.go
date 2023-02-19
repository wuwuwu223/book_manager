package controller

import (
	"book_manager/model"
	"book_manager/service"
	"github.com/gin-gonic/gin"
)

func ReaderLogin(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	reader := service.FindUserByAccount(account)
	if reader == nil {
		c.JSON(200, gin.H{
			"result": "no",
		})
		return
	}
	if reader.Password != password {
		c.JSON(200, gin.H{
			"result": "no",
		})
		return
	}
	c.JSON(200, model.LoginResponse{
		Result:    "yes",
		Condi:     reader.Condi,
		LoginUser: reader,
	})
}

func GetAllReaders(c *gin.Context) {
	account := c.PostForm("account")
	page := c.PostForm("currentPage")
	readers, count := service.FindAllReaders(account, page)
	c.JSON(200, model.GetReadersResponse{
		Readers: readers,
		PageInfo: model.PageInfo{
			Total: int(count),
		},
	})
}

func DelReader(c *gin.Context) {
	account := c.PostForm("account")
	service.DelReader(account)
	c.JSON(200, gin.H{
		"result": "yes",
	})
}

func AddReader(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	name := c.PostForm("name")
	sex := c.PostForm("sex")
	condi := c.PostForm("condi")
	service.AddReader(account, name, password, sex, condi)
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func UpdateReader(c *gin.Context) {
	rid := c.PostForm("rid")
	account := c.PostForm("account")
	password := c.PostForm("password")
	name := c.PostForm("name")
	sex := c.PostForm("sex")
	condi := c.PostForm("condi")
	service.UpdateReader(rid, account, name, password, sex, condi)
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
