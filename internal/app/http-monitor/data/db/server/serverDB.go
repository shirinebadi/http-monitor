package db

import (
	"fmt"

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

func (d *Mydb) Record(request *model.Request) error {
	return d.DB.Create(request).Error
}

func (d *Mydb) Sound(url *model.Url) error {
	fmt.Print("adddd")
	return d.DB.Create(url).Error
}

func (d *Mydb) Search(username string) (bool, error) {
	var stored model.Request
	if err := d.DB.Where(&model.Request{Username: username}).First(&stored).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (d *Mydb) Update(request *model.Request) error {
	return d.DB.Save(request).Error
}

func (d *Mydb) GetFirst() (model.Url, error) {
	var stored model.Url
	err := d.DB.First(&stored).Error

	return stored, err
}

func (d *Mydb) GetRecent(id uint64) (model.Url, error) {
	var stored model.Url
	err := d.DB.Where("ID > ?", id).First(&stored).Error

	return stored, err
}