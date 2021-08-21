package model

type Url struct {
	ID        uint64 `gorm:"primaryKey;auto_increment"`
	Body      string `gorm:"not null"`
	Threshold int    `gorm:"default:5"`
}

func NewUrl(body string, threshold int) *Url {
	url := &Url{Body: body, Threshold: threshold}

	return url
}

type UrlI interface {
	AddUrl(url *Url) error
	SearchId(url string) ([]Url, error)
	SearchUrl(id uint64) (Url, error)
}
