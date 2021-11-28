package response

import (
	"foodcal/business/foods"
	"time"

	"gorm.io/gorm"
)

type FoodInputResponse struct {
	Id          uint           `json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt"`
	Name        string         `json:"name"`
	Cost        string         `json:"cost"`
	Description string         `json:"description"`
	Category    string         `json:"category"`
	Restaurant  string         `json:"restaurant"`
}
type FoodResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Cost        string `json:"cost"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Restaurant  string `json:"restaurant"`
}

func FromDomainInputFood(domain foods.Domain) FoodInputResponse {
	return FoodInputResponse{
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

func FromDomainUpdateFood(domain foods.Domain) FoodInputResponse {
	return FoodInputResponse{
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

func FromDomainAllFood(domain foods.Domain) FoodResponse {
	return FoodResponse{
		Id:          domain.Id,
		Name:        domain.Name,
		Cost:        domain.Cost,
		Description: domain.Description,
		Category:    domain.Category,
		Restaurant:  domain.Restaurant,
	}
}

func FromDomainListFood(domain []foods.Domain) []FoodResponse {
	var menu []FoodResponse
	for _, value := range domain {
		menu = append(menu, FromDomainAllFood(value))
	}

	return menu
}
