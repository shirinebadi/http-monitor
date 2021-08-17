package model

type Response struct {
	Username string `gorm:"unique;not null"`
	Url      Url    `gorm:"foreignKey:UrlId"`
	Status   []Status
}
