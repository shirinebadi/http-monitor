package request

type User struct {
	Username string `json:"name"`
	Password string `json:"password"`
}

func NewUser(username string, password string) *User {
	user := &User{Username: username, Password: password}

	return user
}
