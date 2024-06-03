package mw

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/abhijithk1/api-service-generator/common"
	"github.com/abhijithk1/api-service-generator/mocks"
	"github.com/abhijithk1/api-service-generator/models"
	"github.com/stretchr/testify/assert"
)

func TestMain (m *testing.M) {
	os.Exit(m.Run())
}

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
	apiInputs := models.APIInputs{
		WrkDir: "dir",
		APIGroup: "dummy",
		GoModule: "example",
	}
	fileName := fmt.Sprintf(AuthPath, apiInputs.WrkDir) + "auth.go"
	mockCmdsExecutor.On("CreateFileAndItsContent", fileName,apiInputs, authMiddleWare).Return(nil)

	err := createAuthMiddleWare(apiInputs)
	assert.NoError(t, err)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestSetupMiddleware_Success(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	apiInputs := models.APIInputs{
		WrkDir: "dir",
		APIGroup: "dummy",
		GoModule: "example",
	}

	fileCorsName := fmt.Sprintf(CorsPath, apiInputs.WrkDir) + "cors.go"
	fileAuthName := fmt.Sprintf(AuthPath, apiInputs.WrkDir) + "auth.go"
	
	mockCmdsExecutor.On("CreateFileAndItsContent", fileCorsName, nil, corsMiddleWare).Return(nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", fileAuthName, apiInputs, authMiddleWare).Return(nil)

	err := SetupMiddleWare(apiInputs)
	assert.NoError(t, err)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestSetupMiddleware_AuthMiddleWareError(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	apiInputs := models.APIInputs{
		WrkDir: "dir",
		APIGroup: "dummy",
		GoModule: "example",
	}

	fileCorsName := fmt.Sprintf(CorsPath, apiInputs.WrkDir) + "cors.go"
	fileAuthName := fmt.Sprintf(AuthPath, apiInputs.WrkDir) + "auth.go"
	
	mockCmdsExecutor.On("CreateFileAndItsContent", fileCorsName, nil, corsMiddleWare).Return(nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", fileAuthName, apiInputs, authMiddleWare).Return(errors.New("error in creating auth middleware"))

	err := SetupMiddleWare(apiInputs)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in creating auth middleware")

	mockCmdsExecutor.AssertExpectations(t)
}

func TestSetupMiddleware_CorsMiddleWareError(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	apiInputs := models.APIInputs{
		WrkDir: "dir",
		APIGroup: "dummy",
		GoModule: "example",
	}
	fileCorsName := fmt.Sprintf(CorsPath, apiInputs.WrkDir) + "cors.go"
	
	mockCmdsExecutor.On("CreateFileAndItsContent", fileCorsName, nil, corsMiddleWare).Return(errors.New("error in creating cors middleware"))

	err := SetupMiddleWare(apiInputs)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in creating cors middleware")

	mockCmdsExecutor.AssertExpectations(t)
}
