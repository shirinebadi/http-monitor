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

func (d *Mydb) Add(request *model.Request) error {
	return d.DB.Create(request).Error
}

func (d *Mydb) Update(request *model.Request) error {
	return d.DB.Save(request).Error
}

func (d *Mydb) Search(username string) (bool, error) {
	var stored model.User
	if err := d.DB.Where(&model.Request{Username: username}).First(&stored).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
