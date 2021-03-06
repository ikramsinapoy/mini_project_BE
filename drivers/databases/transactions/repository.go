package transactions

import (
	"foodcal/business/transactions"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(gormDb *gorm.DB) transactions.TransRepoInterface {
	return &TransactionRepository{
		db: gormDb,
	}
}

func (repo *TransactionRepository) Transactions(domain *transactions.Domain) (transactions.Domain, error) {
	userDb := FromDomain(*domain)
	userDb.Status = "Transaksi berhasil"

	err := repo.db.Preload("Food").Preload("User").Create(&userDb).Error
	if err != nil {
		return transactions.Domain{}, err
	}

	return userDb.ToDomain(), nil
}

func (repo *TransactionRepository) HistoryTrans() ([]transactions.Domain, error) {
	var history []Transaction
	err := repo.db.Find(&history).Error

	if err != nil {
		return []transactions.Domain{}, err
	}
	return toDomainHistoryTrans(history), nil
}
