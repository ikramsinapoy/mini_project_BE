package foods

import (
	"foodcal/business/foods"
	"time"

	"gorm.io/gorm"
)

type Food struct {
	Id          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string
	Cost        string
	Description string
	Category    string
	Restaurant  string
}

func (food *Food) ToDomain() foods.Domain {
	return foods.Domain{
		Id:          food.Id,
		CreatedAt:   food.CreatedAt,
		UpdatedAt:   food.UpdatedAt,
		DeletedAt:   food.DeletedAt,
		Name:        food.Name,
		Cost:        food.Cost,
		Description: food.Description,
		Category:    food.Category,
		Restaurant:  food.Restaurant,
	}
}

func FromDomain(domain foods.Domain) Food {
	return Food{
		Id:          domain.Id,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		DeletedAt:   domain.DeletedAt,
		Name:        domain.Name,
		Cost:        domain.Cost,
		Description: domain.Description,
		Category:    domain.Category,
		Restaurant:  domain.Restaurant,
	}
}
func (food Food) toDomainListFood() foods.Domain {
	return foods.Domain{
		Id:          food.Id,
		CreatedAt:   food.CreatedAt,
		UpdatedAt:   food.UpdatedAt,
		DeletedAt:   food.DeletedAt,
		Name:        food.Name,
		Cost:        food.Cost,
		Description: food.Description,
		Category:    food.Category,
		Restaurant:  food.Restaurant,
	}
}

func toDomainMenu(food []Food) []foods.Domain {
	menu := []foods.Domain{}

	for _, value := range food {
		menu = append(menu, value.toDomainListFood())
	}

	return menu
}
