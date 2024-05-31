package mocks

import "github.com/stretchr/testify/mock"

// Mock for ExecuteCmds
type MockCmdsExecutor struct {
	mock.Mock
}

func NewMockCmdsExecutor() *MockCmdsExecutor {
	return new(MockCmdsExecutor)
}

func (m *MockCmdsExecutor) ExecuteCmds(cmdStr string, cmdArgs []string, workDir string) ([]byte, error) {
	args := m.Called(cmdStr, cmdArgs, workDir)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockCmdsExecutor) CreateDirectory(path string) error {
	args := m.Called(path)
	return args.Error(0)
}

func (m *MockCmdsExecutor) CreateFileAndItsContent(fileName string, fileData interface{}, content string) error {
	args := m.Called(fileName, fileData, content)
	return args.Error(0)
}
