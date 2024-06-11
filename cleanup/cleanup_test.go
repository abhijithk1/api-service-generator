package cleanup

import (
	"errors"
	"os"
	"testing"

	"github.com/abhijithk1/api-service-generator/common"
	"github.com/abhijithk1/api-service-generator/mocks"
	"github.com/stretchr/testify/assert"
)

func TestMain (m *testing.M) {
	os.Exit(m.Run())
}

func TestRemoveDirectory_Success(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	wrkDir := "dir"
	cmdStr := "rm"
	cmdArgs := []string{"-rf", wrkDir}
	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs, ".").Return([]byte(""), nil)

	err := removeDirectory(wrkDir)
	assert.NoError(t, err)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestRemoveDirectory_Error(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	wrkDir := "dir"
	cmdStr := "rm"
	cmdArgs := []string{"-rf", wrkDir}
	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs, ".").Return([]byte(""), errors.New("error in removing directory"))

	err := removeDirectory(wrkDir)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in removing directory")

	mockCmdsExecutor.AssertExpectations(t)
}

func TestRemoveDockerContainer_SuccessPostgres(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	containerName := "container"
	volume := "pgdata"
	driver := "postgres"
	cmdStr := "docker"
	cmdArgs1 := []string{"rm", "-f", containerName}
	cmdArgs2 := []string{"volume", "rm", volume}
	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs1, ".").Return([]byte(""), nil)
	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs2, ".").Return([]byte(""), nil)

	err := removeDockerContainer(containerName, driver)
	assert.NoError(t, err)

	mockCmdsExecutor.AssertExpectations(t)
}
func TestRemoveDockerContainer_SuccessMySQL(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	containerName := "container"
	volume := "mysql_data"
	driver := "mysql"
	cmdStr := "docker"
	cmdArgs1 := []string{"rm", "-f", containerName}
	cmdArgs2 := []string{"volume", "rm", volume}
	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs1, ".").Return([]byte(""), nil)
	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs2, ".").Return([]byte(""), nil)

	err := removeDockerContainer(containerName, driver)
	assert.NoError(t, err)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestRemoveDockerContainer_ContainerVolumeError(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	containerName := "container"
	volume := "pgdata"
	driver := "postgres"
	cmdStr := "docker"
	cmdArgs1 := []string{"rm", "-f", containerName}
	cmdArgs2 := []string{"volume", "rm", volume}
	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs1, ".").Return([]byte(""), nil)
	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs2, ".").Return([]byte(""), errors.New("error in removing the container volume"))

	err := removeDockerContainer(containerName, driver)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in removing the container volume")

	mockCmdsExecutor.AssertExpectations(t)
}

func TestRemoveDockerContainer_ContainerError(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	containerName := "container"
	driver := "postgres"
	cmdStr := "docker"
	cmdArgs1 := []string{"rm", "-f", containerName}
	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs1, ".").Return([]byte(""), errors.New("error in stopping the container"))

	err := removeDockerContainer(containerName, driver)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in stopping the container")

	mockCmdsExecutor.AssertExpectations(t)
}

func TestCleanUp_WithoutContainerName(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	wrkDir := "dir"
	containerName := ""
	driver := "postgres"
	cmdStr := "rm"
	cmdArgs := []string{"-rf", wrkDir}
	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs, ".").Return([]byte(""), nil)

	CleanUp(wrkDir, containerName, driver)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestCleanUp_WithoutContainerName_Error(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	wrkDir := "dir"
	containerName := ""
	driver := "postgres"
	cmdStr := "rm"
	cmdArgs := []string{"-rf", wrkDir}
	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs, ".").Return([]byte(""), errors.New("error in removing the directory"))

	CleanUp(wrkDir, containerName, driver)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestCleanUp_WithContainerName(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	wrkDir := "dir"
	containerName := "containerName"
	cmdStr := "rm"
	cmdArgs := []string{"-rf", wrkDir}
	volume := "pgdata"
	driver := "postgres"
	cmdStr2 := "docker"
	cmdArgs1 := []string{"rm", "-f", containerName}
	cmdArgs2 := []string{"volume", "rm", volume}

	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs, ".").Return([]byte(""), nil)
	mockCmdsExecutor.On("ExecuteCmds", cmdStr2, cmdArgs1, ".").Return([]byte(""), nil)
	mockCmdsExecutor.On("ExecuteCmds", cmdStr2, cmdArgs2, ".").Return([]byte(""), nil)
	
	CleanUp(wrkDir, containerName, driver)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestCleanUp_WithContainerName_Error(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	wrkDir := "dir"
	containerName := "containerName"
	cmdStr := "rm"
	cmdArgs := []string{"-rf", wrkDir}
	cmdStr2 := "docker"
	volume := "pgdata"
	driver := "postgres"
	cmdArgs1 := []string{"rm", "-f", containerName}
	cmdArgs2 := []string{"volume", "rm", volume}

	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs, ".").Return([]byte(""), nil)
	mockCmdsExecutor.On("ExecuteCmds", cmdStr2, cmdArgs1, ".").Return([]byte(""), nil)
	mockCmdsExecutor.On("ExecuteCmds", cmdStr2, cmdArgs2, ".").Return([]byte(""), errors.New("error in removing docker container"))
	
	CleanUp(wrkDir, containerName, driver)

	mockCmdsExecutor.AssertExpectations(t)
}
