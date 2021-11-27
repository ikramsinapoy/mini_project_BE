package request

import "foodcal/business/foods"

type Foods struct {
	Name        string `json:"name"`
	Cost        string `json:"cost"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Restaurant  string `json:"restaurant"`
}

func (food Foods) ToDomainInsertFood() *foods.Domain {
	return &foods.Domain{
		Name:        food.Name,
		Cost:        food.Cost,
		Description: food.Description,
		Category:    food.Category,
		Restaurant:  food.Restaurant,
	}
}
