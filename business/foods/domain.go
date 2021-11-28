package foods

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string
	Cost        string
	Description string
	Category    string
	Restaurant  string
}

type FoodUsecaseInterface interface {
	InsertFood(domain *Domain) (Domain, error)
	GetAllFoods() ([]Domain, error)
	DeleteFood(id uint) (string, error)
	UpdateFood(id uint, domain *Domain) (Domain, error)
}

type FoodRepoInterface interface {
	InsertFood(domain *Domain) (Domain, error)
	GetAllFoods() ([]Domain, error)
	DeleteFood(id uint) (string, error)
	UpdateFood(id uint, domain *Domain) (Domain, error)
}
