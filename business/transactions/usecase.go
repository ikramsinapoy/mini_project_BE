package transactions

import (
	// "foodcal/drivers/databases/foods"
	// "foodcal/business/foods"
	"time"
)

type TransUseCases struct {
	repo TransRepoInterface
}

func NewUseCase(transRepo TransRepoInterface, contextTimeout time.Duration) TransUsecaseInterface {
	return &TransUseCases{
		repo: transRepo,
	}
}

func (usecase *TransUseCases) Transactions(domain *Domain) (Domain, error) {
	trans, err := usecase.repo.Transactions(domain)

	if err != nil {
		return Domain{}, err
	}

	return trans, nil
}

func (usecase *TransUseCases) HistoryTrans() ([]Domain, error) {
	food, err := usecase.repo.HistoryTrans()

	if err != nil {
		return []Domain{}, err
	}
	return food, nil
}
