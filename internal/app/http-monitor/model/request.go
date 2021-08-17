package model

type Request struct {
	Username string `gorm:"primaryKey"`
	Urls     []Url  `gorm:"embedded"`
}

type RequestI interface {
	Add(request *Request) error
}
