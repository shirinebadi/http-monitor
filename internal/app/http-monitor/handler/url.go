package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/model"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/request"
)

type UrlHandler struct {
	RequestI model.RequestI
	Token    Token
}

func (h *UrlHandler) Send(c echo.Context) error {

	token := c.Request().Header.Get("Token")
	if token == "" {
		return c.NoContent(http.StatusUnauthorized)
	}
	username, err := h.Token.Parse(token)
	if err != nil {
		log.Error(err)
	}
	fmt.Print(username)

	req := new(request.Url)

	if err := c.Bind(req); err != nil {
		log.Info("Error in Send: %s", err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	fmt.Print("request: ", req)

	url := &model.Url{Body: req.UrlBody, Period: req.Period}
	fmt.Println("url:", url)

	reqToDB := &model.Request{Username: username}
	reqToDB.Urls = append(reqToDB.Urls, fmt.Sprint(url.ID))

	if err := h.RequestI.Add(reqToDB); err != nil {
		log.Error(err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"url":    req.UrlBody,
		"Status": "Sent to Scheduler",
	})

}
