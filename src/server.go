package main

import (
	"github.com/TKUCSC/ecosphere-node/src/config"
	"github.com/TKUCSC/ecosphere-node/src/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initialize() *echo.Echo {
	e := echo.New()

	/********************** middleware **********************/

	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	/********************** routes **********************/

	// add routes for api
	apiRoute := e.Group("/api")
	new(controller.APIController).Data(apiRoute)
	new(controller.APIController).Meta(apiRoute)

	// top level shortcuts
	{
		/** api **/
	}

	return e
}

func main() {
	e := initialize()

	e.Logger.Fatal(e.Start(config.PORT))
}
