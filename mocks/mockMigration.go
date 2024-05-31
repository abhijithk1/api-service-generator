package mocks

import (
	"github.com/abhijithk1/api-service-generator/models"
	"github.com/stretchr/testify/mock"
)

type MockMigration struct {
	mock.Mock
}

func NewMockMigration() *MockMigration {
	return new(MockMigration)
}

func (m *MockMigration) Migration(dbInputs models.DBInputs, initSchema models.InitSchema) (err error) {
	args := m.Called(dbInputs, initSchema)
	return args.Error(0)
}

func (m *MockMigration) RunMigration(dbInputs models.DBInputs) error {
	args := m.Called(dbInputs)
	return args.Error(0)
}