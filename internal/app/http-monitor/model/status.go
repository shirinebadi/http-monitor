package model

import (
	"time"

	"github.com/lib/pq"
)

type Status struct {
	ID         uint64 `gorm:"primaryKey;auto_increment"`
	Username   string `gorm:"not null"`
	Url        uint64
	StatusCode pq.Int32Array `gorm:"type:int[]"`
	Time       time.Time
}

type StatusI interface {
	Record(status *Status) error
	Search(username string) ([]Status, error)
	Update(status *Status) error
	SearchByUrl(username string, url uint64) (Status, error)
}
