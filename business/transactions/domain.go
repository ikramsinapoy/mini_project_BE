package transactions

import (
	"time"
)

type Domain struct {
	Id        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	IdUser    int
	IdFood    int
	Status    string
}

type TransUsecaseInterface interface {
	Transactions(domain *Domain) (Domain, error)
	HistoryTrans() ([]Domain, error)
}

type TransRepoInterface interface {
	Transactions(domain *Domain) (Domain, error)
	HistoryTrans() ([]Domain, error)
}
