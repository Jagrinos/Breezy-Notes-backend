package routes

import (
	"BreeZy_Backend_vol_0/Views"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Getest(c echo.Context) error {
	msg := c.Request().Header.Get("message")
	if msg == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "no message"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": msg + " change"})
}

func Postest(c echo.Context) error {
	var user Views.UserWithoutIdJson

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "INVALID JSON"})
	}

	if user.Login == "" || user.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "EMPTY FIELDS"})
	}

	return c.JSON(http.StatusOK,
		map[string]string{
			"message": fmt.Sprintf("Привет, я тебя узнал, %s %s", user.Login, user.Password),
			"warning": "берегись гнида",
		})
}
