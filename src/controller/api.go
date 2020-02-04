// Package controller s
package controller

import (
	"github.com/TKUCSC/ecosphere-node/src/controller/api"
	"github.com/labstack/echo/v4"
)

// APIController s
type APIController struct{}

// File collection of routes for files
func (c *APIController) Data(g *echo.Group) {
	/********************** file **********************/

	{
		/* upload data */
		g.Any("/data/upload", api.DataUpload())

		/* get data */
		g.Any("/data/download", api.DataDownload())
	}
}

// Meta collection of routes regarding miscellaneous actions
func (c *APIController) Meta(g *echo.Group) {
	/********************** meta **********************/

	metaGroup := g.Group("/meta")
	{
		metaGroup.Any("/hello", api.Hello())
	}
}
