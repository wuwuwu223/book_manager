package model

type Albums struct {
	Aid         int        `json:"aid"`
	Title       string     `json:"title"`
	Author      string     `json:"author"`
	Publisher   string     `json:"publisher"`
	Publishtime string     `json:"publishtime"`
	Num         int        `json:"num"`
	Descri      string     `json:"descri"`
	Time        string     `json:"time"`
	Subalbums   []Subalbum `json:"subalbums" gorm:"foreignKey:Aid;references:Aid"`
}

func (Albums) TableName() string {
	return "album"
}

type Subalbum struct {
	Sid    int    `json:"sid"`
	Aid    int    `json:"aid"`
	Condi  int    `json:"condi"`
	Number string `json:"number"`
	Time   string `json:"time"`
}

func (Subalbum) TableName() string {
	return "subalbum"
}
