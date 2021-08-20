package model

type User struct {
	Username string `gorm:"primaryKey"`
	Password string `gorm:"not null"`
}

func NewUser(username string, password string) *User {
	user := &User{Username: username, Password: password}

	return user
}

type UserI interface {
	Login(username string, password string) (User, error)
	Register(user *User) error
}
