package model

type Response struct {
	Username string  `gorm:"unique;not null"`
	Request  Request `gorm:"foreignkey:UrlId"`
	Status   []Status
}
