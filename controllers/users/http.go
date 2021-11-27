package controllers

import (
	"foodcal/business/users"
	"foodcal/controllers"
	"foodcal/controllers/users/request"
	"foodcal/controllers/users/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	usecase users.UserUsecaseInterface
}

func NewUserController(uc users.UserUsecaseInterface) *UserController {
	return &UserController{
		usecase: uc,
	}
}
func (controller *UserController) Register(c echo.Context) error {
	// ctx := c.Request().Context()
	var userRegister request.Users
	err := c.Bind(&userRegister)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	user, err := controller.usecase.Register(userRegister.ToDomainRegister())
	return controllers.SuccessResponse(c, response.FromDomainRegister(user))
}

func (controller *UserController) Login(c echo.Context) error {
	ctx := c.Request().Context()
	var userLogin request.UserLogin
	err := c.Bind(&userLogin)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	user, err := controller.usecase.Login(*userLogin.ToDomainLogin(), ctx)
	return controllers.SuccessResponse(c, response.FromDomainLogin(user))
}

// func (controller *UserController) GetAllUsers(c echo.Context) error {
// 	return controllers.SuccessResponse(c, response.UserResponse{})
// }
