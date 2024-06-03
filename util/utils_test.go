package util

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/abhijithk1/api-service-generator/common"
	"github.com/abhijithk1/api-service-generator/mocks"
	"github.com/stretchr/testify/assert"
)

func TestMain (m *testing.M) {
	os.Exit(m.Run())
}
func TestCreateConfigFile(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	wrkDir := "dir"

	filePathName := fmt.Sprintf(UtilPath, wrkDir) + "config.go"
	mockCmdsExecutor.On("CreateFileAndItsContent", filePathName, nil, configContent).Return(nil)
	
	err := createConfigFile(wrkDir)
	
	assert.NoError(t, err)
	mockCmdsExecutor.AssertExpectations(t)
}

func TestCreateUtilsFile(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	wrkDir := "dir"

	filePathName := fmt.Sprintf(UtilPath, wrkDir) + "utils.go"
	mockCmdsExecutor.On("CreateFileAndItsContent", filePathName, nil, utilContent).Return(nil)
	
	err := createUtilsFile(wrkDir)
	
	assert.NoError(t, err)
	mockCmdsExecutor.AssertExpectations(t)
}

func TestSetupUtils_Success(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	wrkDir := "dir"

	fileConfigPathName := fmt.Sprintf(UtilPath, wrkDir) + "config.go"
	fileUtilsPathName := fmt.Sprintf(UtilPath, wrkDir) + "utils.go"
	
	mockCmdsExecutor.On("CreateFileAndItsContent", fileConfigPathName, nil, configContent).Return(nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", fileUtilsPathName, nil, utilContent).Return(nil)

	SetUtils(wrkDir)
	mockCmdsExecutor.AssertExpectations(t)
}

func TestSetupUtils_CreateUtilsError(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	wrkDir := "dir"

	fileConfigPathName := fmt.Sprintf(UtilPath, wrkDir) + "config.go"
	fileUtilsPathName := fmt.Sprintf(UtilPath, wrkDir) + "utils.go"
	
	mockCmdsExecutor.On("CreateFileAndItsContent", fileConfigPathName, nil, configContent).Return(nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", fileUtilsPathName, nil, utilContent).Return(errors.New("error in creating the utils.go"))

	SetUtils(wrkDir)
	mockCmdsExecutor.AssertExpectations(t)
}

func TestSetupUtils_CreateConfigError(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	wrkDir := "dir"

	fileConfigPathName := fmt.Sprintf(UtilPath, wrkDir) + "config.go"
	
	mockCmdsExecutor.On("CreateFileAndItsContent", fileConfigPathName, nil, configContent).Return(errors.New("error in creating the config.go"))

	SetUtils(wrkDir)
	mockCmdsExecutor.AssertExpectations(t)
}