package model

type Url struct {
	Username string `gorm:"unique;not null"`
	Urls     []Request
}
