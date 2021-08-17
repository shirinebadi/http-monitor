package handler

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/gommon/log"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/config"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/model"
)

type Token struct {
	Cfg config.Config
}

func (t *Token) GenerateJWT(user model.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Minute * time.Duration((t.Cfg.JWT.Expiration))).Unix()

	return token.SignedString([]byte(t.Cfg.JWT.Secret))
}

func (t *Token) Parse(token string) (string, error) {
	claims := jwt.MapClaims{}

	fmt.Println("chi shode?")
	fmt.Println("token: ", token)
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return []byte(t.Cfg.JWT.Secret), nil
	})

	if err != nil {
		log.Error(err)
		return "", err
	}

	username := claims["username"].(string)

	return username, nil

}
