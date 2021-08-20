package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/model"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/request"
)

type UserHandler struct {
	UserI model.UserI
	Token Token
}

func (h *UserHandler) Register(c echo.Context) error {

	req := new(request.User)

	if err := c.Bind(req); err != nil {
		log.Info("Error in Register: %s", err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	userToDB := model.NewUser(req.Username, req.Password)
	if err := h.UserI.Register(userToDB); err != nil {
		log.Error(err)
	}

	token, _ := h.Token.GenerateJWT(*userToDB)

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func (h *UserHandler) Login(c echo.Context) error {
	req := new(request.User)

	if err := c.Bind(req); err != nil {
		log.Error("Error in Login: %s", err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	user, err := h.UserI.Login(req.Username, req.Password)
	if err != nil {
		log.Error("Error in Login: %s", err.Error())
		return c.NoContent(http.StatusUnauthorized)
	}

	token, _ := h.Token.GenerateJWT(user)

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})

}
