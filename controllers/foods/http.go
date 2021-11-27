package controllers

import (
	"foodcal/business/foods"
	"foodcal/controllers"
	"foodcal/controllers/foods/request"
	"foodcal/controllers/foods/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FoodController struct {
	usecase foods.FoodUsecaseInterface
}

func NewFoodController(uc foods.FoodUsecaseInterface) *FoodController {
	return &FoodController{
		usecase: uc,
	}
}

func (controller *FoodController) InsertFood(c echo.Context) error {
	var insertFood request.Foods
	err := c.Bind(&insertFood)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	food, err := controller.usecase.InsertFood(insertFood.ToDomainInsertFood())
	return controllers.SuccessResponse(c, response.FromDomainInputFood(food))
}

func (controller *FoodController) GetAllFood(c echo.Context) error {
	food, err := controller.usecase.GetAllFoods()

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "error binding", err)
	}

	return controllers.SuccessResponse(c, response.FromDomainListFood(food))
}

func (controller *FoodController) DeleteFood(c echo.Context) error {
	param := c.Param("foodId")
	foodID, _ := strconv.Atoi(param)
	food, err := controller.usecase.DeleteFood(uint(foodID))
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "error binding", err)
	}

	return controllers.SuccessResponse(c, food)
}
