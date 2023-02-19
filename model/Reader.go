package model

type Reader struct {
	Rid      int    `json:"rid"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Sex      string `json:"sex"`
	Time     string `json:"time"`
	Condi    int    `json:"condi"`
}

func (Reader) TableName() string {
	return "reader"
}

type LoginResponse struct {
	Result    string  `json:"result"`
	Condi     int     `json:"condi"`
	LoginUser *Reader `json:"loginUser"`
}

type GetReadersResponse struct {
	Readers  []*Reader `json:"readers"`
	PageInfo PageInfo  `json:"pageInfo"`
}
type PageInfo struct {
	Total int `json:"total"`
}
