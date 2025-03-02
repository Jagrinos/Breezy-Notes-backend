package net

import (
	"BreeZy_Backend_vol_0/internal/net/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupEcho() (*echo.Echo, error) {
	e := echo.New()
	e.Use(middleware.Logger(), middleware.Recover())
	e.GET("/api/getest", routes.Getest)
	e.POST("/api/postest", routes.Postest)

	return e, nil
}
