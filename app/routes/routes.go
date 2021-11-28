package routes

import (
	foodController "foodcal/controllers/foods"
	transactionController "foodcal/controllers/transactions"
	userController "foodcal/controllers/users"

	"github.com/labstack/echo/v4"
)

type RouteControllerList struct {
	UserController        userController.UserController
	FoodController        foodController.FoodController
	TransactionController transactionController.TransactionController
}

func (controller RouteControllerList) RouteRegister(c *echo.Echo) {
	users := c.Group("/user")
	users.POST("", controller.UserController.Register)
	users.POST("/login", controller.UserController.Login)

	foods := c.Group("/food")
	foods.POST("", controller.FoodController.InsertFood)
	foods.GET("/menu", controller.FoodController.GetAllFood)
	foods.DELETE("/:foodId", controller.FoodController.DeleteFood)
	foods.PUT("/update/:foodId", controller.FoodController.UpdateFood)

	transaction := c.Group("/transaction")
	transaction.POST("", controller.TransactionController.Transactions)
	transaction.GET("/history", controller.TransactionController.HistoryTrans)
}
