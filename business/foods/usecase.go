package foods

import (
	"time"
)

type FoodUseCases struct {
	repo FoodRepoInterface
}

func NewUseCase(foodRepo FoodRepoInterface, contextTimeout time.Duration) FoodUsecaseInterface {
	return &FoodUseCases{
		repo: foodRepo,
	}
}

func (usecase *FoodUseCases) InsertFood(domain *Domain) (Domain, error) {
	user, err := usecase.repo.InsertFood(domain)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (usecase *FoodUseCases) GetAllFoods() ([]Domain, error) {
	food, err := usecase.repo.GetAllFoods()

	if err != nil {
		return []Domain{}, err
	}
	return food, nil
}
