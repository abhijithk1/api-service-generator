package mocks

import "github.com/stretchr/testify/mock"

// Mock for ExecuteCmds
type MockCmdsExecutor struct {
	mock.Mock
}

func NewMockCmdsExecutor() *MockCmdsExecutor {
	return new(MockCmdsExecutor)
}

func (m *MockCmdsExecutor) ExecuteCmds(cmdStr string, cmdArgs []string) ([]byte, error) {
	args := m.Called(cmdStr, cmdArgs)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockCmdsExecutor) CreateDirectory(path string) error {
	args := m.Called(path)
	return args.Error(0)
}

// func (m *MockExecutor) ExecuteGoMod(path, name string) error {
// 	args := m.Called(path, name)
// 	return args.Error(0)
// }

// func (m *MockExecutor) ExecuteGoGets() error {
// 	args := m.Called()
// 	return args.Error(0)
// }