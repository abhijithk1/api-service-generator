package docker

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

func TestRunDockerContainer_Success(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	postgresRun = `docker run postgres`

	cmdStr := "docker"
	cmdArgs := []string{"run", "postgres"}

	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs).Return([]byte(""), nil)

	err := RunPostgresContainer()
	assert.NoError(t, err)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestRunDockerContainer_Error(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	postgresRun = `docker run postgres`

	cmdStr := "docker"
	cmdArgs := []string{"run", "postgres"}

	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs).Return([]byte(""), errors.New("error in running docker run"))

	err := RunPostgresContainer()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in running docker run")

	mockCmdsExecutor.AssertExpectations(t)
}