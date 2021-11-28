package transactions

import "time"

type TransUseCases struct {
	repo TransRepoInterface
}

func NewUseCase(transRepo TransRepoInterface, contextTimeout time.Duration) TransUsecaseInterface {
	return &TransUseCases{
		repo: transRepo,
	}
}

func (usecase *TransUseCases) Transactions(idUser int, domain *Domain) (Domain, error) {
	trans, err := usecase.repo.Transactions(idUser, domain)
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
