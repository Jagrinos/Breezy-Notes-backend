package net

import (
	"errors"
	jwt2 "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"uasbreezy/config/views"
	"uasbreezy/internal/pkg/jwt"
	"uasbreezy/pkg/users"
)

func (db DriverDb) AuthenticationHandler(c echo.Context) error {
	var u views.UserAuth
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]string{
				"error": "INVALID JSON",
			})
	}

	err := users.Auth(db.Driver, u)
	if err != nil {
		return c.JSON(http.StatusBadGateway, //TODO StatusUnauthorized
			map[string]string{
				"error": err.Error(),
			})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "user found",
	})
}

func (db DriverDb) AuthorizationHandler(c echo.Context) error {
	login := c.QueryParam("login")
	if login == "" {
		return c.JSON(http.StatusBadRequest,
			map[string]string{
				"error": "login is required",
			})
	}

	at, err := jwt.GenerateToken(login, "ACCESS")
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{
				"error": err.Error(),
			})
	}
	rt, err := jwt.GenerateToken(login, "REFRESH")
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{
				"error": err.Error(),
			})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"accesstoken":  at,
		"refreshtoken": rt,
	})
}

func (db DriverDb) RefreshHandler(c echo.Context) error {
	rt := c.QueryParam("refreshtoken")
	if rt == "" {
		return c.JSON(http.StatusBadRequest,
			map[string]string{
				"error": "refresh token is required",
			})
	}

	login, err := jwt.GetLoginFromToken(rt)
	if err != nil {
		if errors.Is(err, jwt2.ErrTokenExpired) {
			return c.JSON(http.StatusUnauthorized,
				map[string]string{
					"error": errors.New("token is expired").Error(),
				})
		}
		return c.JSON(http.StatusInternalServerError,
			map[string]string{
				"error": err.Error(),
			})
	}

	at, err := jwt.Refresh(rt, login)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{
				"error": err.Error(),
			})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"accesstoken": at,
	})
}

func (db DriverDb) CheckTokenHandler(c echo.Context) error {
	at := c.QueryParam("accesstoken")
	if at == "" {
		return c.JSON(http.StatusBadRequest,
			map[string]string{
				"error": "access token is required",
			})
	}
	_, err := jwt.VerifyToken(at)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, //TODO TOKEN IS EXPIRED ERROR
			map[string]string{
				"error": err.Error(),
			})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "token good",
	})
}

func (db DriverDb) GetInfoHandler(c echo.Context) error {
	login := c.QueryParam("login")
	if login == "" {
		return c.JSON(http.StatusBadRequest,
			map[string]string{
				"error": "login is required",
			})
	}

	u, err := users.GetInfo(db.Driver, login)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{
				"error": err.Error(),
			})
	}

	return c.JSON(http.StatusOK, u)
}
