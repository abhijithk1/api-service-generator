package cleanup

import (
	"errors"
	"testing"

	"github.com/abhijithk1/api-service-generator/common"
	"github.com/abhijithk1/api-service-generator/mocks"
	"github.com/stretchr/testify/assert"
)

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

func TestRemoveDockerContainer_Success(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	containerName := "container"
	cmdStr := "docker"
	cmdArgs1 := []string{"kill", containerName}
	cmdArgs2 := []string{"rm", containerName}
	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs1, ".").Return([]byte(""), nil)
	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs2, ".").Return([]byte(""), nil)

	err := removeDockerContainer(containerName)
	assert.NoError(t, err)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestRemoveDockerContainer_KillContainerError(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	containerName := "container"
	cmdStr := "docker"
	cmdArgs1 := []string{"kill", containerName}
	cmdArgs2 := []string{"rm", containerName}
	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs1, ".").Return([]byte(""), nil)
	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs2, ".").Return([]byte(""), errors.New("error in removing the container"))

	err := removeDockerContainer(containerName)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in removing the container")

	mockCmdsExecutor.AssertExpectations(t)
}

func TestRemoveDockerContainer_StopContainerError(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	containerName := "container"
	cmdStr := "docker"
	cmdArgs1 := []string{"kill", containerName}
	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs1, ".").Return([]byte(""), errors.New("error in stopping the container"))

	err := removeDockerContainer(containerName)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in stopping the container")

	mockCmdsExecutor.AssertExpectations(t)
}

func TestCleanUp_WithoutContainerName(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	wrkDir := "dir"
	containerName := ""
	cmdStr := "rm"
	cmdArgs := []string{"-rf", wrkDir}
	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs, ".").Return([]byte(""), nil)

	CleanUp(wrkDir, containerName)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestCleanUp_WithoutContainerName_Error(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	wrkDir := "dir"
	containerName := ""
	cmdStr := "rm"
	cmdArgs := []string{"-rf", wrkDir}
	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs, ".").Return([]byte(""), errors.New("error in removing the directory"))

	CleanUp(wrkDir, containerName)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestCleanUp_WithContainerName(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	wrkDir := "dir"
	containerName := "containerName"
	cmdStr := "rm"
	cmdArgs := []string{"-rf", wrkDir}
	cmdStr2 := "docker"
	cmdArgs1 := []string{"kill", containerName}
	cmdArgs2 := []string{"rm", containerName}

	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs, ".").Return([]byte(""), nil)
	mockCmdsExecutor.On("ExecuteCmds", cmdStr2, cmdArgs1, ".").Return([]byte(""), nil)
	mockCmdsExecutor.On("ExecuteCmds", cmdStr2, cmdArgs2, ".").Return([]byte(""), nil)
	
	CleanUp(wrkDir, containerName)

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
	cmdArgs1 := []string{"kill", containerName}
	cmdArgs2 := []string{"rm", containerName}

	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs, ".").Return([]byte(""), nil)
	mockCmdsExecutor.On("ExecuteCmds", cmdStr2, cmdArgs1, ".").Return([]byte(""), nil)
	mockCmdsExecutor.On("ExecuteCmds", cmdStr2, cmdArgs2, ".").Return([]byte(""), errors.New("error in removing docker container"))
	
	CleanUp(wrkDir, containerName)

	mockCmdsExecutor.AssertExpectations(t)
}
