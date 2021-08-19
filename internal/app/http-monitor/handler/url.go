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
	UrlI     model.UrlI
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

	url := new(model.Url)
	url.Body = req.UrlBody
	url.Period = req.Period

	if err := h.UrlI.Sound(url); err != nil {
		fmt.Print(err)
	}

	fmt.Println("url:", url)

	reqToDB := &model.Request{Username: username}
	reqToDB.Urls = append(reqToDB.Urls, fmt.Sprint(url.ID))

	result, err := h.RequestI.Search(username)

	if err != nil {
		log.Error("Error in Add: ", err)
	}

	if result {
		fmt.Print("1")
		if err := h.RequestI.Update(reqToDB); err != nil {
			log.Error("Error in Add: ", err)
		}
	} else {
		fmt.Print("2")
		if err := h.RequestI.Record(reqToDB); err != nil {
			log.Error("Error in Add: ", err)
		}
	}

	return c.JSON(http.StatusOK, map[string]string{
		"url":    req.UrlBody,
		"Status": "Sent to Scheduler",
	})

}
