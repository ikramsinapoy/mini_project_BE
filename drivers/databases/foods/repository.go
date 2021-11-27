package foods

import (
	"fmt"
	"foodcal/business/foods"

	"gorm.io/gorm"
)

type FoodRepository struct {
	db *gorm.DB
}

func NewFoodRepository(gormDb *gorm.DB) foods.FoodRepoInterface {
	return &FoodRepository{
		db: gormDb,
	}
}

func (repo *FoodRepository) InsertFood(domain *foods.Domain) (foods.Domain, error) {
	userDb := FromDomain(*domain)
	err := repo.db.Create(&userDb).Error
	if err != nil {
		return foods.Domain{}, err
	}

	return userDb.ToDomain(), nil
}

func (repo *FoodRepository) GetAllFoods() ([]foods.Domain, error) {
	var food []Food
	err := repo.db.Find(&food).Error

	fmt.Println(food)

	if err != nil {
		return []foods.Domain{}, err
	}
	return toDomainMenu(food), nil
}

func (repo *FoodRepository) DeleteFood(id uint) (string, error) {
	rec := Food{}
	err := repo.db.Delete(&rec, "id = ?", id).Error

	if err != nil {
		return "", err
	}
	return "Product Deleted", nil
}
