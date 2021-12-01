package users_test

import (
	"context"
	"foodcal/app/middlewares"
	"foodcal/business/users"
	_userMock "foodcal/business/users/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var UserRepoInterfaceMock _userMock.UserRepoInterface
var userUsecase users.UserUsecaseInterface
var userDomain users.Domain

// func TestMain(m *testing.T){

// }
func TestMain(m *testing.M) {
	userUsecase = users.NewUseCase(&UserRepoInterfaceMock, 2, &middlewares.ConfigJWT{})
	userDomain = users.Domain{
		Id:       1,
		Username: "ikramsinapoy",
		Address:  "jl.brigjen katamso",
		Phone:    "085298984475",
		Email:    "ikramsinapoy8@gmail.com",
		Password: "123",
		Dob:      "05-08-2001",
		Token:    "",
	}

	m.Run()
}

func TestRegister(t *testing.T) {
	t.Run("Success Register", func(t *testing.T) {
		UserRepoInterfaceMock.On("Register",
			mock.Anything).Return(userDomain, nil).Once()

		var requestLoginDomain = users.Domain{
			Id:       1,
			Username: "ikramsinapoy",
			Address:  "jl.brigjen katamso",
			Phone:    "085298984475",
			Email:    "ikramsinapoy8@gmail.com",
			Password: "123",
			Dob:      "05-08-2001",
		}
		domain, err := userUsecase.Register(&requestLoginDomain)

		assert.Nil(t, err)
		assert.NotEmpty(t, domain)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Success Login", func(t *testing.T) {
		UserRepoInterfaceMock.On("Login",
			mock.Anything,
			mock.Anything).Return(userDomain, nil).Once()

		var requestLoginDomain = users.Domain{
			Email:    "ikramsinapoy@gmail.com",
			Password: "123",
		}
		domain, err := userUsecase.Login(requestLoginDomain, context.Background())

		assert.Nil(t, err)
		assert.NotEmpty(t, domain)
	})

	t.Run("Unsuccessful Login", func(t *testing.T) {
		UserRepoInterfaceMock.On("Login",
			mock.Anything,
			mock.Anything).Return(users.Domain{}, assert.AnError).Once()

		var requestLoginDomain = users.Domain{
			Email:    "ikramsinapoy@gmail.com",
			Password: "123",
		}
		domain, err := userUsecase.Login(requestLoginDomain, context.Background())

		assert.NotNil(t, err)
		assert.Empty(t, domain)
	})
}
