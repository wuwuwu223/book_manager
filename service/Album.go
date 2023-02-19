package service

import (
	"book_manager/global"
	"book_manager/model"
	"log"
	"strconv"
	"time"
)

func FindAllAlbums(title, page string) ([]*model.Albums, int64) {
	var albums []*model.Albums
	var count int64
	pageInt, _ := strconv.Atoi(page)
	if title == "" {
		err := global.Db.Model(model.Albums{}).Preload("Subalbums").Count(&count).Offset((pageInt - 1) * 10).Limit(10).Find(&albums)
		if err.Error != nil {
			log.Println(err.Error)
			return nil, 0
		}
		return albums, count
	}
	err := global.Db.Model(model.Albums{}).Preload("Subalbums").Where("title like ?", "%"+title+"%").Count(&count).Offset((pageInt - 1) * 10).Limit(10).Find(&albums)
	if err.Error != nil {
		log.Println(err.Error)
		return nil, 0
	}
	return albums, count

}

func DelAlbum(aid string) {
	err := global.Db.Where("aid = ?", aid).Delete(&model.Albums{})
	if err.Error != nil {
		log.Println(err.Error)
	}
	return
}

func AddAlbum(title, publisher, author, publishtime, descri string) {
	var album model.Albums
	album.Title = title
	album.Publisher = publisher
	album.Author = author
	album.Publishtime = publishtime
	album.Descri = descri
	err := global.Db.Create(&album)
	if err.Error != nil {
		log.Println(err.Error)
	}
	return
}

func AddSubAlbum(number, aid string) bool {
	var subalbum model.Subalbum
	subalbum.Number = number
	aidInt, _ := strconv.Atoi(aid)
	subalbum.Aid = aidInt
	subalbum.Time = time.Now().Format("2006-01-02 15:04:05")
	//检查是否有重复的number
	var count int64
	err := global.Db.Model(model.Subalbum{}).Where("number = ? and aid=?", number, aid).Count(&count)
	if err.Error != nil {
		log.Println(err.Error)
		return false
	}
	if count != 0 {
		return false
	}
	err = global.Db.Create(&subalbum)
	if err.Error != nil {
		log.Println(err.Error)
	}
	//更新album表中的count
	var album model.Albums
	err = global.Db.Where("aid = ?", aid).First(&album)
	if err.Error != nil {
		log.Println(err.Error)
		return false
	}
	album.Num = album.Num + 1
	err = global.Db.Model(model.Albums{}).Where("aid=?", aid).Update("num", album.Num)
	return true
}
