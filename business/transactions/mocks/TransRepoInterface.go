// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	transactions "foodcal/business/transactions"

	mock "github.com/stretchr/testify/mock"
)

// TransRepoInterface is an autogenerated mock type for the TransRepoInterface type
type TransRepoInterface struct {
	mock.Mock
}

// HistoryTrans provides a mock function with given fields:
func (_m *TransRepoInterface) HistoryTrans() ([]transactions.Domain, error) {
	ret := _m.Called()

	var r0 []transactions.Domain
	if rf, ok := ret.Get(0).(func() []transactions.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]transactions.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Transactions provides a mock function with given fields: domain
func (_m *TransRepoInterface) Transactions(domain *transactions.Domain) (transactions.Domain, error) {
	ret := _m.Called(domain)

	var r0 transactions.Domain
	if rf, ok := ret.Get(0).(func(*transactions.Domain) transactions.Domain); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Get(0).(transactions.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*transactions.Domain) error); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
