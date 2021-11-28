package transactions

import (
	"foodcal/business/transactions"
	"time"
)

type Transaction struct {
	Id        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	IdUser    int
	IdFood    int
	Status    bool
}

func (transaction *Transaction) ToDomain() transactions.Domain {
	return transactions.Domain{
		Id:        transaction.Id,
		CreatedAt: transaction.CreatedAt,
		UpdatedAt: transaction.UpdatedAt,
		IdUser:    transaction.IdUser,
		IdFood:    transaction.IdFood,
		Status:    transaction.Status,
	}
}

func FromDomain(domain transactions.Domain) Transaction {
	return Transaction{
		Id:        domain.Id,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		IdUser:    domain.IdUser,
		IdFood:    domain.IdFood,
		Status:    domain.Status,
	}
}

func toDomainHistoryTrans(transaction []Transaction) []transactions.Domain {
	history := []transactions.Domain{}

	for _, v := range transaction {
		history = append(history, v.ToDomain())
	}

	return history
}
