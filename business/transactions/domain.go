package transactions

import "time"

type Domain struct {
	Id        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	IdUser    int
	IdFood    int
	Status    bool
}

type TransUsecaseInterface interface {
	Transactions(userId int, domain *Domain) (Domain, error)
	HistoryTrans() ([]Domain, error)
}

type TransRepoInterface interface {
	Transactions(userId int, domain *Domain) (Domain, error)
	HistoryTrans() ([]Domain, error)
}
