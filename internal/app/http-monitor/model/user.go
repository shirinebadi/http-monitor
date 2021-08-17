package model

type User struct {
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

type UserI interface {
	Login(username string, password string) (User, error)
	Register(user *User) error
}
