package request

import "foodcal/business/transactions"

type Transactions struct {
	IdFood int `json:"id_food"`
}

func (food Transactions) ToDomainTransaction() *transactions.Domain {
	return &transactions.Domain{
		IdFood: food.IdFood,
	}
}
