package mocks

import "github.com/stretchr/testify/mock"

type MockCommon struct {
	mock.Mock
}

func NewMockCommon() *MockCommon {
	return new(MockCommon)
}

func (m *MockCommon) MarshalYAML(v interface{}) ([]byte, error) {
	args := m.Called(v)
	return nil, args.Error(1)
}