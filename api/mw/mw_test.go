package mw

import (
	"errors"
	"fmt"
	"testing"

	"github.com/abhijithk1/api-service-generator/common"
	"github.com/abhijithk1/api-service-generator/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateCorsMiddleware(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	wrkDir := "dir"
	fileName := fmt.Sprintf(CorsPath, wrkDir) + "cors.go"
	mockCmdsExecutor.On("CreateFileAndItsContent", fileName, nil, corsMiddleWare).Return(nil)

	err := createCorsMiddleWare(wrkDir)
	assert.NoError(t, err)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestCreateAuthMiddleware(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor
	wrkDir := "dir"
	fileName := fmt.Sprintf(AuthPath, wrkDir) + "auth.go"
	mockCmdsExecutor.On("CreateFileAndItsContent", fileName, nil, authMiddleWare).Return(nil)

	err := createAuthMiddleWare(wrkDir)
	assert.NoError(t, err)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestSetupMiddleware_Success(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	wrkDir := "dir"
	fileCorsName := fmt.Sprintf(CorsPath, wrkDir) + "cors.go"
	fileAuthName := fmt.Sprintf(AuthPath, wrkDir) + "auth.go"
	
	mockCmdsExecutor.On("CreateFileAndItsContent", fileCorsName, nil, corsMiddleWare).Return(nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", fileAuthName, nil, authMiddleWare).Return(nil)

	err := SetupMiddleWare(wrkDir)
	assert.NoError(t, err)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestSetupMiddleware_AuthMiddleWareError(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	wrkDir := "dir"
	fileCorsName := fmt.Sprintf(CorsPath, wrkDir) + "cors.go"
	fileAuthName := fmt.Sprintf(AuthPath, wrkDir) + "auth.go"
	
	mockCmdsExecutor.On("CreateFileAndItsContent", fileCorsName, nil, corsMiddleWare).Return(nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", fileAuthName, nil, authMiddleWare).Return(errors.New("error in creating auth middleware"))

	err := SetupMiddleWare(wrkDir)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in creating auth middleware")

	mockCmdsExecutor.AssertExpectations(t)
}

func TestSetupMiddleware_CorsMiddleWareError(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	wrkDir := "dir"
	fileCorsName := fmt.Sprintf(CorsPath, wrkDir) + "cors.go"
	
	mockCmdsExecutor.On("CreateFileAndItsContent", fileCorsName, nil, corsMiddleWare).Return(errors.New("error in creating cors middleware"))

	err := SetupMiddleWare(wrkDir)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in creating cors middleware")

	mockCmdsExecutor.AssertExpectations(t)
}
