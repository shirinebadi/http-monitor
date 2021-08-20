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
	StatusI model.StatusI
	UrlI    model.UrlI
	Token   Token
	Jobs    chan model.Status
}

func (h *UrlHandler) Send(c echo.Context) error {
	req := new(request.Url)

	if err := c.Bind(req); err != nil {
		log.Info("Error in Send: %s", err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	url := model.NewUrl(req.UrlBody, req.Threshold)
	if err := h.UrlI.AddUrl(url); err != nil {
		fmt.Print(err)
	}

	token := c.Request().Header.Get("Token")
	if token == "" {
		return c.NoContent(http.StatusUnauthorized)
	}

	username, err := h.Token.Parse(token)
	if err != nil {
		log.Error(err)
	}

	status := model.NewStatus(username, url.ID)
	if err := h.StatusI.Record(status); err != nil {
		log.Error("Error in Add: ", err)
	}

	h.Jobs <- *status

	return c.JSON(http.StatusOK, map[string]string{
		"url":    req.UrlBody,
		"Status": "Sent to Scheduler",
	})

}
