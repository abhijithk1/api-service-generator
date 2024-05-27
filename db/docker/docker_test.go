package docker

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/abhijithk1/api-service-generator/common"
	"github.com/abhijithk1/api-service-generator/mocks"
	"github.com/abhijithk1/api-service-generator/models"
	"github.com/stretchr/testify/assert"
)

func TestRunDockerContainer_Success(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	dbInput := models.DBInputs{
		PsqlUser:     "root",
		PsqlPassword: "password",
		DBName:       "postgres",
	}

	runCmd := fmt.Sprintf("docker run --name postgres_db -p 5432:5432 -e POSTGRES_USER=%s -e POSTGRES_PASSWORD=%s POSTGRES_DB=%s -d postgres", "root", "password", "postgres")

	cmdSplits := strings.Split(runCmd, " ")

	cmdStr := cmdSplits[0]
	cmdArgs := cmdSplits[1:]

	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs).Return([]byte(""), nil)

	err := RunPostgresContainer(dbInput)
	assert.NoError(t, err)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestRunDockerContainer_Error(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	dbInput := models.DBInputs{
		PsqlUser:     "root",
		PsqlPassword: "password",
		DBName:       "postgres",
	}

	runCmd := fmt.Sprintf("docker run --name postgres_db -p 5432:5432 -e POSTGRES_USER=%s -e POSTGRES_PASSWORD=%s POSTGRES_DB=%s -d postgres", "root", "password", "postgres")

	cmdSplits := strings.Split(runCmd, " ")

	cmdStr := cmdSplits[0]
	cmdArgs := cmdSplits[1:]

	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs).Return([]byte(""), errors.New("error in running docker run"))

	err := RunPostgresContainer(dbInput)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in running docker run")

	mockCmdsExecutor.AssertExpectations(t)
}
