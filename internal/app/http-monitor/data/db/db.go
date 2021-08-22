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

func (d *Mydb) Record(status *model.Status) error {
	return d.DB.Create(status).Error
}

func (d *Mydb) AddUrl(url *model.Url) error {
	err := d.DB.Create(url).Error

	return err
}

func (d *Mydb) SearchUrl(id uint64) (model.Url, error) {
	var stored model.Url
	err := d.DB.Where(&model.Url{ID: id}).First(&stored).Error

	return stored, err
}

func (d *Mydb) SearchId(url string) ([]model.Url, error) {
	var stored []model.Url
	err := d.DB.Where(&model.Url{Body: url}).First(&stored).Error

	return stored, err
}

func (d *Mydb) Search(username string) ([]model.Status, error) {
	var stored []model.Status
	err := d.DB.Where(&model.Status{Username: username}).Find(&stored).Error

	return stored, err
}

func (d *Mydb) SearchByUrl(username string, url uint64) (model.Status, error) {
	var stored model.Status
	err := d.DB.Where(&model.Status{Username: username, Url: url}).First(&stored).Error

	return stored, err
}

func (d *Mydb) Update(status *model.Status) error {
	return d.DB.Save(status).Error
}

func (d *Mydb) GetFirst(id uint64) (model.Status, error) {
	var stored model.Status
	err := d.DB.Where(&model.Status{ID: id}).First(&stored).Error

	return stored, err
}

func (d *Mydb) GetRecent(id uint64) (model.Status, error) {
	var stored model.Status
	err := d.DB.Where("ID > ?", id).First(&stored).Error

	return stored, err
}
