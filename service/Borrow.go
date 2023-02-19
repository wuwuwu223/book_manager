package service

import (
	"book_manager/global"
	"book_manager/model"
	"log"
	"strconv"
	"time"
)

func FindAllBorrowRecords(raaccount string, page string) ([]*model.BorrowRecordList, int64) {
	var records []*model.BorrowRecord
	var count int64
	pageInt, _ := strconv.Atoi(page)
	var result []*model.BorrowRecordList
	if raaccount == "" {
		err := global.Db.Model(model.BorrowRecord{}).Count(&count).Offset((pageInt - 1) * 10).Limit(10).Find(&records)
		if err.Error != nil {
			log.Println(err.Error)
			return nil, 0
		}
		for _, record := range records {
			var resultRecord model.BorrowRecordList
			resultRecord.BorrowRecord = *record
			var reader model.Reader
			err = global.Db.Model(model.Reader{}).Where("rid = ?", record.Rid).First(&reader)
			if err.Error != nil {
				log.Println(err.Error)
				return nil, 0
			}
			var album model.Albums
			err = global.Db.Model(model.Albums{}).Where("aid = ?", record.Aid).First(&album)
			if err.Error != nil {
				log.Println(err.Error)
				return nil, 0
			}
			var subalbum model.Subalbum
			err = global.Db.Model(model.Subalbum{}).Where("sid = ?", record.Sid).First(&subalbum)
			if err.Error != nil {
				log.Println(err.Error)
				return nil, 0
			}
			resultRecord.Album = album
			resultRecord.Subalbum = subalbum
			resultRecord.Reader = reader
			result = append(result, &resultRecord)
		}
		return result, count
	}
	err := global.Db.Model(model.BorrowRecord{}).Where("raccount like ?", "%"+raaccount+"%").Count(&count).Offset((pageInt - 1) * 10).Limit(10).Find(&records)
	if err.Error != nil {
		log.Println(err.Error)
		return nil, 0
	}
	for _, record := range records {
		var resultRecord model.BorrowRecordList
		resultRecord.BorrowRecord = *record
		var reader model.Reader
		err = global.Db.Model(model.Reader{}).Where("rid = ?", record.Rid).First(&reader)
		if err.Error != nil {
			log.Println(err.Error)
			return nil, 0
		}
		var album model.Albums
		err = global.Db.Model(model.Albums{}).Where("aid = ?", record.Aid).First(&album)
		if err.Error != nil {
			log.Println(err.Error)
			return nil, 0
		}
		var subalbum model.Subalbum
		err = global.Db.Model(model.Subalbum{}).Where("sid = ?", record.Sid).First(&subalbum)
		if err.Error != nil {
			log.Println(err.Error)
			return nil, 0
		}
		resultRecord.Album = album
		resultRecord.Subalbum = subalbum
		resultRecord.Reader = reader
		result = append(result, &resultRecord)
	}
	return result, count
}

func DelBorrowRecord(bid, sid string) {
	global.Db.Model(model.BorrowRecord{}).Where("bid = ?", bid).Delete(model.BorrowRecord{})
	//更新书的状态
	global.Db.Model(model.Subalbum{}).Where("sid = ?", sid).Update("condi", 0)
}

func AddBorrowRecord(raccount, rid, aid string) {
	var reader model.Reader
	global.Db.Model(model.Reader{}).Where("rid = ?", rid).First(&reader)
	var borrowRecord model.BorrowRecord
	borrowRecord.Rid = reader.Rid
	borrowRecord.Raccount = raccount
	borrowRecord.Aid, _ = strconv.Atoi(aid)
	//先查找是否有可借的书
	var subalbum model.Subalbum
	global.Db.Model(model.Subalbum{}).Where("aid = ? and condi = ?", aid, 0).First(&subalbum)
	if subalbum.Sid == 0 {
		return
	}
	sid := strconv.Itoa(subalbum.Sid)
	//更新书的状态
	global.Db.Model(model.Subalbum{}).Where("sid = ?", sid).Update("condi", 1)
	borrowRecord.Sid, _ = strconv.Atoi(sid)
	t := time.Now()
	borrowRecord.Time = t.Format("2006-01-02 15:04:05")
	t2 := t.AddDate(0, 0, 15)
	borrowRecord.Backtime = t2.Format("2006-01-02 15:04:05")
	borrowRecord.Inttime = strconv.FormatInt(t2.UnixMilli(), 10)
	global.Db.Model(model.BorrowRecord{}).Create(&borrowRecord)
}
