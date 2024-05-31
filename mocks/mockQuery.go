package mocks

import (
	"github.com/abhijithk1/api-service-generator/models"
	"github.com/stretchr/testify/mock"
)

type MockQuery struct {
	mock.Mock
}

func NewMockQuery() *MockQuery {
	return new(MockQuery)
}

func (m * MockQuery) SetTableQuery(initSchema models.InitSchema) error {
	args := m.Called(initSchema)
	return args.Error(0)
}