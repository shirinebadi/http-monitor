package model

type Url struct {
	UrlBody string `gorm:"not null"`
	Period  int    `gorm:"default:5"`
}
