package response

import (
	"foodcal/business/transactions"
	"time"
)

type TransResponse struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IdUser    int       `json:"id_user"`
	IdFood    int       `json:"id_food"`
	Status    string    `json:"status"`
}

func FromDomainTransaction(domain transactions.Domain) TransResponse {
	return TransResponse{
		Id:        domain.Id,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		IdUser:    domain.IdUser,
		IdFood:    domain.IdFood,
		Status:    domain.Status,
	}
}

func FromDomainHistoryTrans(domain []transactions.Domain) []TransResponse {
	var history []TransResponse
	for _, v := range domain {
		history = append(history, FromDomainTransaction(v))
	}

	return history
}
