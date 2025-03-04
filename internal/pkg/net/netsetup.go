package net

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Echo struct {
	Echo *echo.Echo
}

func GetEcho(db DriverDb) Echo {
	e := echo.New()
	e.Use(middleware.Logger(), middleware.Recover())

	e.GET("/api/users/getall", db.GetAllUserHandler)
	e.POST("/api/users/create", db.CreateUserHandler)
	e.PUT("/api/users/update", db.UpdateUserHandler)
	e.DELETE("/api/users/delete", db.DeleteUserHandler)

	return Echo{Echo: e}
}
