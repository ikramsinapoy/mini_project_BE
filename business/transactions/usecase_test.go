package transactions_test

import (
	"foodcal/business/transactions"
	_transMock "foodcal/business/transactions/mocks"
	"testing"
	"time"

	// "github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/mock"
)

var TransRepoInterfaceMock _transMock.TransRepoInterface
var transUsecase transactions.TransUsecaseInterface
var transDomain transactions.Domain

func TestMain(m *testing.M) {
	transUsecase = transactions.NewUseCase(&TransRepoInterfaceMock, 2)
	transDomain = transactions.Domain{
		Id:        1,
		IdUser:    4,
		IdFood:    5,
		Status:    "Transaksi berhasil",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	m.Run()
}

func TestTransactions(t *testing.T) {

	// t.Run("Successfull Transaction", func(t *testing.T) {
	// 	TransRepoInterfaceMock.On("Transaction",
	// 		mock.AnythingOfType("transactions.Domain")).Return(transDomain, nil).Once()

	// 	var requestTransDomain = transactions.Domain{
	// 		Id:     1,
	// 		IdUser: 4,
	// 		IdFood: 5,
	// 	}

	// 	domain, err := transUsecase.Transactions(&requestTransDomain)

	// 	assert.Nil(t, err)
	// 	assert.Equal(t, transDomain, domain)
	// })

	// t.Run("Unsuccessfull Transaction", func(t *testing.T) {
	// 	TransRepoInterfaceMock.On("Transaction",
	// 		mock.AnythingOfType("transactions.Domain")).Return(-1, nil).Once()

	// 	var requestTransDomain = &transactions.Domain{
	// 		Id:     1,
	// 		IdUser: 4,
	// 		IdFood: 5,
	// 		Status: "Transaksi berhasil",
	// 	}

	// 	domain, err := transUsecase.Transactions(requestTransDomain)

	// 	assert.NotNil(t, err)
	// 	assert.NotEqual(t, transDomain, domain)
	// })
}

func TestHistoryTrans(t *testing.T) {
	t.Run("Valid test", func(t *testing.T) {
		TransRepoInterfaceMock.On("HistoryTrans").Return([]transactions.Domain{transDomain}, nil).Once()

		trans, err := transUsecase.HistoryTrans()
		assert.Nil(t, err)
		assert.Contains(t, trans, transDomain)
	})

	t.Run("Invalid test", func(t *testing.T) {
		TransRepoInterfaceMock.On("HistoryTrans").Return([]transactions.Domain{}, assert.AnError).Once()

		trans, err := transUsecase.HistoryTrans()
		assert.NotNil(t, err)
		assert.NotContains(t, trans, transDomain)
	})
}
