package model

import "time"

type Status struct {
	Id         uint64 `gorm:"primaryKey"`
	Username   string `gorm:"not null"`
	Url        uint64
	StatusCode int
	Time       time.Time
}

type StatusI interface {
	Record(status *Status) error
	Search(username string) (bool, error)
	Update(status *Status) error
}
