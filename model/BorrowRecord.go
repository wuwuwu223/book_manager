package model

type BorrowRecord struct {
	Bid      int    `json:"bid"`
	Rid      int    `json:"rid"`
	Sid      int    `json:"sid"`
	Aid      int    `json:"aid"`
	Raccount string `json:"raccount"`
	Time     string `json:"time"`
	Backtime string `json:"backtime"`
	Inttime  string `json:"inttime"`
}

func (BorrowRecord) TableName() string {
	return "borrowrecord"
}

type BorrowRecordList struct {
	BorrowRecord
	Reader   Reader   `json:"reader"`
	Album    Albums   `json:"album"`
	Subalbum Subalbum `json:"subalbum"`
}
