package request

import "foodcal/business/transactions"

type Transactions struct {
	Id          uint   `json:"id"`
	IdFood      int    `json:"id_food"`
	IdUser      int    `json:"id_user"`
	ProductName string `json:"product_name"`
}

func (trans Transactions) ToDomainTransaction() *transactions.Domain {
	return &transactions.Domain{
		IdFood: trans.IdFood,
		IdUser: trans.IdUser,
	}
}
