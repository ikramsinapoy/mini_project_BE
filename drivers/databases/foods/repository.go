package foods

import (
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

func (repo *FoodRepository) UpdateFood(id uint, domain *foods.Domain) (foods.Domain, error) {
	foodDb := FromDomain(*domain)

	food := repo.db.Where("id = ?", id).Updates(&foodDb)
	if food.Error != nil {
		return foods.Domain{}, food.Error
	}

	return foodDb.ToDomain(), nil
}
