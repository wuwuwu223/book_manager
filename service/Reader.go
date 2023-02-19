package service

import (
	"book_manager/global"
	"book_manager/model"
	"log"
	"strconv"
	"time"
)

func FindUserByAccount(account string) (reader *model.Reader) {
	err := global.Db.Where("account = ?", account).First(&reader)
	if err.Error != nil {
		log.Println(err.Error)
		return
	}
	return
}

func FindAllReaders(account, page string) ([]*model.Reader, int64) {
	var readers []*model.Reader
	var count int64
	pageInt, _ := strconv.Atoi(page)
	if account == "" {
		err := global.Db.Model(model.Reader{}).Count(&count).Offset((pageInt - 1) * 10).Limit(10).Find(&readers)
		if err.Error != nil {
			log.Println(err.Error)
			return nil, 0
		}
		return readers, count
	}
	err := global.Db.Model(model.Reader{}).Where("account like ?", "%"+account+"%").Count(&count).Offset((pageInt - 1) * 10).Limit(10).Find(&readers)
	if err.Error != nil {
		log.Println(err.Error)
		return nil, 0
	}
	return readers, count
}

func DelReader(account string) {
	err := global.Db.Where("account = ?", account).Delete(&model.Reader{})
	if err.Error != nil {
		log.Println(err.Error)
	}
	return
}

func AddReader(account, name, password, sex, condi string) {
	var reader model.Reader
	reader.Account = account
	reader.Name = name
	reader.Password = password
	reader.Sex = sex
	condiInt, _ := strconv.Atoi(condi)
	reader.Condi = condiInt
	reader.Time = time.Now().Format("2006-01-02 15:04:05")
	err := global.Db.Create(&reader)
	if err.Error != nil {
		log.Println(err.Error)
		return
	}
}

func UpdateReader(rid, account, name, password, sex, condi string) {
	var reader *model.Reader
	err := global.Db.Where("rid = ?", rid).First(&reader)
	reader.Account = account
	reader.Name = name
	reader.Password = password
	reader.Sex = sex
	condiInt, _ := strconv.Atoi(condi)
	reader.Condi = condiInt
	//reader.Time = time.Now().Format("2006-01-02 15:04:05")
	err = global.Db.Where("rid = ?", rid).Save(&reader)
	if err.Error != nil {
		log.Println(err.Error)
		return
	}
}
