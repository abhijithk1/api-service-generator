package api

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

func TestCreateAPIGroup(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor
	apiInputs := models.APIInputs{
		WrkDir: "dir",
		APIGroup: "dummy",
		APIGroupTitle: "Dummy",
		TableName: "table_name",
		TableNameTitle: "TableName",
	}

	filePath := fmt.Sprintf(APIFilePath, apiInputs.WrkDir,apiInputs.APIGroup)
	mockCmdsExecutor.On("CreateDirectory",filePath).Return(nil)

	err := createApiGroup(apiInputs)
	
	assert.NoError(t, err)
	mockCmdsExecutor.AssertExpectations(t)
}

func TestCreateControllerFile(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor
	apiInputs := models.APIInputs{
		WrkDir: "dir",
		APIGroup: "dummy",
		APIGroupTitle: "Dummy",
		TableName: "table_name",
		TableNameTitle: "TableName",
	}

	fileName := fmt.Sprintf(APIFilePath, apiInputs.WrkDir,apiInputs.APIGroup) + "controller.go"
	mockCmdsExecutor.On("CreateFileAndItsContent", fileName, apiInputs, controllerContent).Return(nil)

	err := createControllerFile(apiInputs)
	assert.NoError(t, err)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestCreateServiceFile(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor
	apiInputs := models.APIInputs{
		WrkDir: "dir",
		APIGroup: "dummy",
		APIGroupTitle: "Dummy",
		TableName: "table_name",
		TableNameTitle: "TableName",
	}

	fileName := fmt.Sprintf(APIFilePath, apiInputs.WrkDir,apiInputs.APIGroup) + "service.go"
	mockCmdsExecutor.On("CreateFileAndItsContent", fileName, apiInputs, serviceContent).Return(nil)

	err := createServiceFile(apiInputs)
	assert.NoError(t, err)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestSetup_Success(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor
	apiInputs := models.APIInputs{
		WrkDir: "dir",
		APIGroup: "dummy",
		APIGroupTitle: "Dummy",
		TableName: "table_name",
		TableNameTitle: "TableName",
	}

	filePath := fmt.Sprintf(APIFilePath, apiInputs.WrkDir,apiInputs.APIGroup)
	fileControllerName := fmt.Sprintf(APIFilePath, apiInputs.WrkDir,apiInputs.APIGroup) + "controller.go"
	fileServiceName := fmt.Sprintf(APIFilePath, apiInputs.WrkDir,apiInputs.APIGroup) + "service.go"
	
	mockCmdsExecutor.On("CreateDirectory",filePath).Return(nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", fileControllerName, apiInputs, controllerContent).Return(nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", fileServiceName, apiInputs, serviceContent).Return(nil)

	Setup(apiInputs)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestSetup_ServiceFileError(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor
	apiInputs := models.APIInputs{
		WrkDir: "dir",
		APIGroup: "dummy",
		APIGroupTitle: "Dummy",
		TableName: "table_name",
		TableNameTitle: "TableName",
	}

	filePath := fmt.Sprintf(APIFilePath, apiInputs.WrkDir,apiInputs.APIGroup)
	fileControllerName := fmt.Sprintf(APIFilePath, apiInputs.WrkDir,apiInputs.APIGroup) + "controller.go"
	fileServiceName := fmt.Sprintf(APIFilePath, apiInputs.WrkDir,apiInputs.APIGroup) + "service.go"
	
	mockCmdsExecutor.On("CreateDirectory",filePath).Return(nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", fileControllerName, apiInputs, controllerContent).Return(nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", fileServiceName, apiInputs, serviceContent).Return(errors.New("error in creating service.go"))

	Setup(apiInputs)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestSetup_ControllerFileError(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor
	apiInputs := models.APIInputs{
		WrkDir: "dir",
		APIGroup: "dummy",
		APIGroupTitle: "Dummy",
		TableName: "table_name",
		TableNameTitle: "TableName",
	}

	filePath := fmt.Sprintf(APIFilePath, apiInputs.WrkDir,apiInputs.APIGroup)
	fileControllerName := fmt.Sprintf(APIFilePath, apiInputs.WrkDir,apiInputs.APIGroup) + "controller.go"
	
	mockCmdsExecutor.On("CreateDirectory",filePath).Return(nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", fileControllerName, apiInputs, controllerContent).Return(errors.New("error in creating controller.go"))

	Setup(apiInputs)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestSetup_CreateAPIGroupError(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor
	apiInputs := models.APIInputs{
		WrkDir: "dir",
		APIGroup: "dummy",
		APIGroupTitle: "Dummy",
		TableName: "table_name",
		TableNameTitle: "TableName",
	}

	filePath := fmt.Sprintf(APIFilePath, apiInputs.WrkDir,apiInputs.APIGroup)
	
	mockCmdsExecutor.On("CreateDirectory",filePath).Return(errors.New("error in creating API Group directory"))

	Setup(apiInputs)

	mockCmdsExecutor.AssertExpectations(t)
}
