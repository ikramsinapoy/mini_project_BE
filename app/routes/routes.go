package routes

import (
	foodController "foodcal/controllers/foods"
	userController "foodcal/controllers/users"

	"github.com/labstack/echo/v4"
)

type RouteControllerList struct {
	UserController userController.UserController
	FoodController foodController.FoodController
}

func (controller RouteControllerList) RouteRegister(c *echo.Echo) {
	users := c.Group("/user")
	users.POST("", controller.UserController.Register)
	users.POST("/login", controller.UserController.Login)

	foods := c.Group("/food")
	foods.POST("", controller.FoodController.InsertFood)
	foods.GET("/menu", controller.FoodController.GetAllFood)
}
