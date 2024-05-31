package mocks

import (
	"github.com/abhijithk1/api-service-generator/models"
	"github.com/stretchr/testify/mock"
)

type MockDocker struct {
	mock.Mock
}

func NewMockDocker() *MockDocker {
	return new(MockDocker)
}

func (m * MockDocker) RunContainer(dbInputs models.DBInputs) error {
	args := m.Called(dbInputs)
	return args.Error(0)
}