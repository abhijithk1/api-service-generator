package db

import (
	"errors"
	"os"
	"testing"

	"github.com/abhijithk1/api-service-generator/common"
	"github.com/abhijithk1/api-service-generator/mocks"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestGenerateSQLC_Success(t *testing.T) {
	mockExec := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockExec

	cmdStr := "sqlc"
	cmdArgs := []string{"init"}

	wrkDir := "file"

	mockExec.On("ExecuteCmds", cmdStr, cmdArgs, wrkDir).Return([]byte(""), nil)

	err := initialiseSQLC(wrkDir)
	assert.NoError(t, err)
	mockExec.AssertExpectations(t)
}

func TestGenerateSQLC_Error(t *testing.T) {
	mockExec := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockExec

	cmdStr := "sqlc"
	cmdArgs := []string{"init"}

	wrkDir := "file"

	mockExec.On("ExecuteCmds", cmdStr, cmdArgs, wrkDir).Return([]byte(""), errors.New("error in sqlc initialisation"))

	err := initialiseSQLC(wrkDir)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in sqlc initialisation")
	mockExec.AssertExpectations(t)
}

func TestEditSQLCYaml_Success(t *testing.T) {

}

func TestEditSQLCYaml_ErrorMarshalling(t *testing.T) {

}

func TestEditSQLCYaml_ErrorWritingFile(t *testing.T) {

}
