package api

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{}
)

// Hello gonna do something cool here idk yet
func Hello() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		// todo
		return c.JSON(200, map[string]string{"sucess": "true"})
	}
}
