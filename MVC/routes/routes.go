package routes

import (
	"foodcal/controllers/users"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/login", users.LoginController)
	e.POST("/register", users.InsertUsersController)
	return e
}
