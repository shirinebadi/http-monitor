package handler

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/config"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/model"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/request"
)

type UserHandler struct {
	UserI model.UserI
}

func (h *UserHandler) Register(c echo.Context) error {

	req := new(request.User)

	if err := c.Bind(req); err != nil {
		log.Info("Error in Register: %s", err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	userToDB := &model.User{Username: req.Username, Password: req.Password}
	if err := h.UserI.Register(userToDB); err != nil {
		log.Fatal(err)
	}

	token, _ := h.generateJWT(*userToDB)

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func (h *UserHandler) Login(c echo.Context) error {
	req := new(request.User)

	if err := c.Bind(req); err != nil {
		log.Info("Error in Register: %s", err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	user, err := h.UserI.Login(req.Username, req.Password)
	if err != nil {
		log.Info("Error in Login: %s", err.Error())
		return c.NoContent(http.StatusNetworkAuthenticationRequired)
	}

	token, _ := h.generateJWT(user)

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})

}

func (h *UserHandler) generateJWT(user model.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	var cfg config.Config

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(cfg.JWT.Expiration)).Unix()

	return token.SignedString([]byte(cfg.JWT.Secret))
}
