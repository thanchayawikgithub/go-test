package main

import (
	"errors"
	"fmt"

	"github.com/stretchr/testify/mock"
)

func main() {
	customerRepo := &CustomerRepositoryMock{}
	customerRepo.On("GetCustomer", 1).Return("Bond", 20, nil)
	customerRepo.On("GetCustomer", 2).Return("", 0, errors.New("customer not found"))

	name, age, err := customerRepo.GetCustomer(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(name, age)
}

type CustomerRepositoryMock struct {
	mock.Mock
}

func (m *CustomerRepositoryMock) GetCustomer(id int) (name string, age int, err error) {
	args := m.Called(id)
	return args.String(0), args.Int(1), args.Error(2)
}
