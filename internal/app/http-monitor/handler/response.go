package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/config"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/model"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/request"
)

type ResponseHandler struct {
	Cfg     config.Config
	StatusI model.StatusI
	Token   Token
	UrlI    model.UrlI
}

func (h *ResponseHandler) Get(c echo.Context) error {
	token := c.Request().Header.Get("Token")
	if token == "" {
		return c.NoContent(http.StatusUnauthorized)
	}

	username, err := h.Token.Parse(token)
	if err != nil {
		log.Error(err)
	}

	totalStatus, err := h.StatusI.Search(username)
	if err != nil {
		log.Error("Error in Response, Get:", err)
	}

	keyVal := make(map[string][]int32)

	for _, s := range totalStatus {
		body, err := h.UrlI.SearchUrl(s.ID)
		if err != nil {
			log.Error("Error in Get Results: ", err)
			continue
		}

		keyVal[body.Body] = s.StatusCode
	}

	return c.JSON(http.StatusOK, keyVal)
}

func (h *ResponseHandler) Post(c echo.Context) error {
	req := new(request.Url)

	if err := c.Bind(req.UrlBody); err != nil {
		log.Error("Error in Response: %s", err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	token := c.Request().Header.Get("Token")
	if token == "" {
		return c.NoContent(http.StatusUnauthorized)
	}

	username, err := h.Token.Parse(token)
	if err != nil {
		log.Error(err)
	}

	urls, _ := h.UrlI.SearchId(req.UrlBody)

	for _, u := range urls {
		resp, err := h.StatusI.SearchByUrl(username, u.ID)
		if err != nil {
			log.Error(err)
			continue
		}

		return c.JSON(http.StatusOK, map[string][]int32{
			req.UrlBody: resp.StatusCode,
		})

	}

	return c.NoContent(http.StatusBadRequest)

}

func (h *ResponseHandler) Alert(c echo.Context) error {
	alert := 0
	alerts := make(map[string]bool)

	token := c.Request().Header.Get("Token")
	if token == "" {
		return c.NoContent(http.StatusUnauthorized)
	}

	username, err := h.Token.Parse(token)
	if err != nil {
		log.Error(err)
	}

	totalStatus, err := h.StatusI.Search(username)
	if err != nil {
		log.Error("Error in Response, Get:", err)
	}

	for _, s := range totalStatus {
		alert = 0
		body, err := h.UrlI.SearchUrl(s.ID)
		if err != nil {
			log.Error("Error in Alert: ", err)
		}

		statusCode := s.StatusCode
		for _, c := range statusCode {
			if c != int32(400) {
				alert++
			}
		}

		if alert >= body.Threshold {
			alerts[body.Body] = true
		}
	}

	return c.JSON(http.StatusOK, alerts)

}
