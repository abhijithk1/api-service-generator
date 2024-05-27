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
		ContainerName: "postgres_db",
		PsqlUser:     "root",
		PsqlPassword: "password",
		DBName:       "postgres",
		WrkDir:       "wrkdir",
	}

	runCmd := fmt.Sprintf(PostgresRun, "postgres_db", "root", "password", "postgres")

	cmdSplits := strings.Split(runCmd, " ")

	cmdStr := cmdSplits[0]
	cmdArgs := cmdSplits[1:]

	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs, ".").Return([]byte(""), nil)

	err := RunPostgresContainer(dbInput)
	assert.NoError(t, err)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestRunDockerContainer_Error(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	dbInput := models.DBInputs{
		ContainerName: "postgres_db",
		PsqlUser:     "root",
		PsqlPassword: "password",
		DBName:       "postgres",
		WrkDir:       "wrkdir",
	}

	runCmd := fmt.Sprintf(PostgresRun, "postgres_db", "root", "password", "postgres")

	cmdSplits := strings.Split(runCmd, " ")

	cmdStr := cmdSplits[0]
	cmdArgs := cmdSplits[1:]

	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs, ".").Return([]byte(""), errors.New("error in running docker run"))

	err := RunPostgresContainer(dbInput)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in running docker run")

	mockCmdsExecutor.AssertExpectations(t)
}

func TestRunDockerContainer_SpecialCaseError(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	dbInput := models.DBInputs{
		ContainerName: "postgres_db",
		PsqlUser:     "root",
		PsqlPassword: "password",
		DBName:       "postgres",
		WrkDir:       "wrkdir",
	}

	runCmd := fmt.Sprintf(PostgresRun, "postgres_db", "root", "password", "postgres")

	cmdSplits := strings.Split(runCmd, " ")

	cmdStr := cmdSplits[0]
	cmdArgs := cmdSplits[1:]

	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs, ".").Return([]byte(`The container name "/postgres_db" is already in use by container`), errors.New("error in running docker run"))

	err := RunPostgresContainer(dbInput)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in running docker run")

	mockCmdsExecutor.AssertExpectations(t)
}
