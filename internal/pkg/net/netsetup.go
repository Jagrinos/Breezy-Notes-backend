package net

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

type Echo struct {
	Echo *echo.Echo
}

func setupLimiter() echo.MiddlewareFunc {
	config := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper, //skip for specific IP or request
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{
				Rate:      rate.Limit(10),
				Burst:     15,
				ExpiresIn: 3 * time.Minute,
			},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(c echo.Context, err error) error {
			if err == nil {
				err = errors.New("something went wrong in LIMITER")
			}
			return c.JSON(http.StatusInternalServerError,
				map[string]string{
					"error": err.Error(),
				})
		},
		DenyHandler: func(c echo.Context, identifier string, err error) error {
			if err == nil {
				err = errors.New("too many request")
			}
			return c.JSON(http.StatusTooManyRequests,
				map[string]string{
					"error":      err.Error(),
					"identifier": identifier,
				})
		},
	}

	return middleware.RateLimiterWithConfig(config)
}

func GetEcho(db DriverDb) Echo {
	e := echo.New()
	e.Use(middleware.Logger(), middleware.Recover())
	e.Use(setupLimiter())
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
