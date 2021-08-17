package model

import "github.com/lib/pq"

type Request struct {
	Username string         `gorm:"primaryKey"`
	Urls     pq.StringArray `gorm:"type:text[]"`
}

type RequestI interface {
	Add(request *Request) error
	Search(username string) (bool, error)
	Update(request *Request) error
}
