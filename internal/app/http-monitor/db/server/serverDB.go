package db

import (
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/model"
	"gorm.io/gorm"
)

type Mydb struct {
	DB *gorm.DB
}

func (d *Mydb) Login(username string, password string) (model.User, error) {
	var stored model.User
	err := d.DB.Where(&model.User{Username: username, Password: password}).First(&stored).Error

	return stored, err
}

func (d *Mydb) Register(user *model.User) error {
	return d.DB.Create(user).Error
}
