package api

import (
	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"
)

// DataUpload uploads sensor data
func DataUpload() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.JSON(fasthttp.StatusOK, map[string]string{
			"Upload!": "Hello World!",
		})
	}
}

// DataDownload downloads sensor data
func DataDownload() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.JSON(fasthttp.StatusOK, map[string]string{
			"Download": "Hello World!",
		})
	}
}
