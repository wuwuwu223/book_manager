package controller

import (
	"book_manager/service"
	"github.com/gin-gonic/gin"
)

func GetAllBorrowRecords(c *gin.Context) {
	raaccount := c.PostForm("raaccount")
	page := c.PostForm("currentPage")
	records, count := service.FindAllBorrowRecords(raaccount, page)
	c.JSON(200, gin.H{
		"borrowrecords": records,
		"pageInfo": gin.H{
			"total": count,
		},
	})

}

func DelBorrowRecord(c *gin.Context) {
	bid := c.PostForm("bid")
	sid := c.PostForm("sid")
	service.DelBorrowRecord(bid, sid)
	c.JSON(200, gin.H{
		"status": "yes",
	})
}

func AddBorrowRecord(c *gin.Context) {
	raccount := c.PostForm("raccount")
	rid := c.PostForm("rid")
	aid := c.PostForm("aid")
	service.AddBorrowRecord(raccount, rid, aid)
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
