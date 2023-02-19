package controller

import (
	"book_manager/service"
	"github.com/gin-gonic/gin"
)

func GetAllAlbums(c *gin.Context) {
	title := c.PostForm("title")
	page := c.PostForm("currentPage")
	albums, count := service.FindAllAlbums(title, page)
	c.JSON(200, gin.H{
		"albums": albums,
		"pageInfo": gin.H{
			"total": count,
		},
	})
}

func DelAlbum(c *gin.Context) {
	aid := c.PostForm("aid")
	service.DelAlbum(aid)
	c.JSON(200, gin.H{
		"result": "yes",
	})
}

func AddAlbum(c *gin.Context) {
	title := c.PostForm("title")
	publisher := c.PostForm("publisher")
	author := c.PostForm("author")
	publishtime := c.PostForm("publishtime")
	descri := c.PostForm("descri")
	service.AddAlbum(title, publisher, author, publishtime, descri)
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func AddSubAlbum(c *gin.Context) {
	number := c.PostForm("number")
	aid := c.PostForm("aid")
	ok := service.AddSubAlbum(number, aid)
	if ok {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	} else {
		c.JSON(200, gin.H{
			"status": "no",
		})
	}
}
