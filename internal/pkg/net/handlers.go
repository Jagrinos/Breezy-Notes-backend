package net

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"uasbreezy/config/views"
	"uasbreezy/internal/pkg/db/users"
)

func (db DriverDb) GetAllUserHandler(c echo.Context) error {
	usersls, err := users.GetAll(db.Driver)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{
				"error": err.Error(),
			})
	}

	if usersls == nil {
		usersls = []views.User{}
	}

	return c.JSON(http.StatusOK, usersls)
}

func (db DriverDb) CreateUserHandler(c echo.Context) error {
	var newUser views.UserNoId
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]string{
				"error": "INVALID JSON",
			})
	}

	err := users.Create(db.Driver, views.User{
		Id:       uuid.NewString(),
		Login:    newUser.Login,
		Password: newUser.Password,
		Email:    newUser.Email,
		About:    newUser.About,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]string{
				"error": err.Error(),
			})
	}

	return c.JSON(http.StatusOK,
		map[string]string{
			"message": "user post successfully",
		})
}

func (db DriverDb) UpdateUserHandler(c echo.Context) error {
	var ud views.UserNoId
	if err := c.Bind(&ud); err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]string{
				"error": "INVALID JSON",
			})
	}

	id := c.QueryParam("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest,
			map[string]string{
				"error": "id parameter is required",
			})
	}

	if err := users.Update(db.Driver, ud, id); err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]string{
				"error": err.Error(),
			})
	}

	return c.JSON(http.StatusOK,
		map[string]string{
			"message": "user update successfully",
		})
}

func (db DriverDb) DeleteUserHandler(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest,
			map[string]string{
				"error": "id parameter is required",
			})
	}

	if err := users.Delete(db.Driver, id); err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]string{
				"error": err.Error(),
			})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "user delete successfully",
	})
}
