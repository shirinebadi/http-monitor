package model

type Url struct {
	ID     uint64 `gorm:"primaryKey;auto_increment"`
	Body   string `gorm:"not null"`
	Period int    `gorm:"default:5"`
}

type UrlI interface {
	AddUrl(url *Url) error
	SearchId(url string) ([]Url, error)
}
