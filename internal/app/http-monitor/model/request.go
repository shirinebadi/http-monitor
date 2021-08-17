package model

type Request struct {
	UrlId   int    `gorm:"primaryKey;AUTO_INCREMENT"`
	UrlBody string `gorm:"not null"`
	Period  int    `gorm:"default:5"`
}
