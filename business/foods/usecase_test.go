package foods_test

import (
	"foodcal/business/foods"
	_foodMock "foodcal/business/foods/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var FoodRepoInterfaceMock _foodMock.FoodRepoInterface
var foodUsecase foods.FoodUsecaseInterface
var foodDomain foods.Domain
var updateFoodDomain foods.Domain

func TestMain(m *testing.M) {
	foodUsecase = foods.NewUseCase(&FoodRepoInterfaceMock, 2)
	foodDomain = foods.Domain{
		Id:          1,
		Name:        "ayam goreng",
		Cost:        "Rp.20.000",
		Description: "ayam goreng rekomendasi",
		Category:    "makanan berat",
		Restaurant:  "KFC",
		Calorie:     "234 calories",
	}

	m.Run()
}

func TestInsertFood(t *testing.T) {
	t.Run("Success Insert", func(t *testing.T) {
		FoodRepoInterfaceMock.On("InsertFood",
			mock.Anything).Return(foodDomain, nil).Once()

		var requestInsertDomain = foods.Domain{
			Name:        "ayam goreng",
			Cost:        "Rp.20.000",
			Description: "ayam goreng rekomendasi",
			Category:    "makanan berat",
			Restaurant:  "KFC",
			Calorie:     "234",
		}

		domain, err := foodUsecase.InsertFood(&requestInsertDomain)

		assert.Nil(t, err)
		assert.Equal(t, foodDomain, domain)

	})
}

func TestDeleteFood(t *testing.T) {
	t.Run("Success delete", func(t *testing.T) {
		FoodRepoInterfaceMock.On("DeleteFood",
			mock.AnythingOfType("uint")).Return("Product deleted", nil).Once()

		domain, err := foodUsecase.DeleteFood(5)

		assert.Nil(t, err)
		assert.Equal(t, "Product deleted", domain)
	})

	t.Run("Unsuccess delete", func(t *testing.T) {
		FoodRepoInterfaceMock.On("DeleteFood",
			mock.AnythingOfType("uint")).Return("", nil).Once()

		domain, err := foodUsecase.DeleteFood(5)

		assert.Nil(t, err)
		assert.Equal(t, "", domain)
	})
}

func TestUpdateFood(t *testing.T) {
	t.Run("Success Update", func(t *testing.T) {
		FoodRepoInterfaceMock.On("UpdateFood", mock.AnythingOfType("uint"),
			mock.Anything).Return(updateFoodDomain, nil).Once()

		var requestUpdateDomain = foods.Domain{
			Name:        "pizza",
			Cost:        "Rp.150.000",
			Description: "saus keju mozarella",
			Category:    "makanan berat",
			Restaurant:  "PHD",
			Calorie:     "314",
		}

		domain, err := foodUsecase.UpdateFood(5, &requestUpdateDomain)

		assert.Nil(t, err)
		assert.Equal(t, updateFoodDomain, domain)

	})
}
