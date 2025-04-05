package net

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupEcho() (*echo.Echo, error) {
	e := echo.New()
	e.Use(middleware.Logger(), middleware.Recover())
	//e.GET("/api/users/all", mongonet.GetAllUsersHandler)
	//e.POST("/api/users", mongonet.PostUserHandler)
	//e.PUT("/api/users", mongonet.PutUserHandler)
	//e.DELETE("/api/users", mongonet.DeleteUserHandler)

	return e, nil
}
