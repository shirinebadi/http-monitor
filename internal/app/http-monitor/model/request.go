package model

import "github.com/lib/pq"

type Request struct {
	Username string         `gorm:"primaryKey"`
	Urls     pq.StringArray `gorm:"type:text[]"`
}
