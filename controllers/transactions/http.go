package controllers

import (
	"foodcal/business/transactions"
	"foodcal/controllers"
	"foodcal/controllers/transactions/request"
	"foodcal/controllers/transactions/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	usecase transactions.TransUsecaseInterface
}

func NewTransactionController(uc transactions.TransUsecaseInterface) *TransactionController {
	return &TransactionController{
		usecase: uc,
	}
}

func (controller *TransactionController) Transactions(c echo.Context) error {
	trans := request.Transactions{}
	err := c.Bind(&trans)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}

	transaction, err := controller.usecase.Transactions(id, trans.ToDomainTransaction())
	return controllers.SuccessResponse(c, response.FromDomainTransaction(transaction))
}

func (controller *TransactionController) HistoryTrans(c echo.Context) error {
	transaction, err := controller.usecase.HistoryTrans()

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "error binding", err)
	}

	return controllers.SuccessResponse(c, response.FromDomainHistoryTrans(transaction))
}
