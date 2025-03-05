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
	//user
	e.GET("/api/users/getall", db.GetAllUserHandler)
	e.POST("/api/users/create", db.CreateUserHandler)
	e.PUT("/api/users/update", db.UpdateUserHandler)
	e.DELETE("/api/users/delete", db.DeleteUserHandler)
	e.GET("/api/users/getinfo", db.GetInfoHandler)

	//auth
	e.POST("/api/authentication", db.AuthenticationHandler)
	e.GET("/api/authorization", db.AuthorizationHandler)
	e.GET("/api/refresh", db.RefreshHandler)
	e.GET("/api/checktoken", db.CheckTokenHandler)

	return Echo{Echo: e}
}
