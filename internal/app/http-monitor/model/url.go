package model

type Url struct {
	ID     uint64 `gorm:"primaryKey"`
	Body   string `gorm:"not null"`
	Period int    `gorm:"default:5"`
}
