package users

import (
	"foodcal/models/response"
	"foodcal/models/users"
	"net/http"

	"foodcal/configs"

	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {

	var user users.User
	c.Bind(&user)

	if user.Password == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			http.StatusBadRequest,
			"Password Kosong",
			nil,
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		http.StatusOK,
		"success",
		user,
	})

}

func InsertUsersController(c echo.Context) error {
	var user users.User
	c.Bind(&user)

	result := configs.DB.Create(&user)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			http.StatusInternalServerError,
			result.Error.Error(),
			nil,
		})
	}

	var response = response.BaseResponse{
		http.StatusOK,
		"success",
		user,
	}

	return c.JSON(http.StatusOK, response)
}
