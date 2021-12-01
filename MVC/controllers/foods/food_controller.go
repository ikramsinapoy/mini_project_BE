package foods

import (
	"foodcal/configs"
	"foodcal/models/foods"
	"foodcal/models/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InsertFoodsController(c echo.Context) error {
	var food foods.Food
	c.Bind(&food)

	result := configs.DB.Create(&food)

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
		food,
	}

	return c.JSON(http.StatusOK, response)

}
