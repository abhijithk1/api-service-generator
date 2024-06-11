package docker

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/abhijithk1/api-service-generator/common"
	"github.com/abhijithk1/api-service-generator/mocks"
	"github.com/abhijithk1/api-service-generator/models"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestRunContainer_Postgres(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	dbInput := models.DBInputs{
		DBMS:          "postgres",
		ContainerName: "postgres_db",
		ContainerPort: 6432,
		Postgres: models.PostgresDriver{
			PsqlUser:      "root",
			PsqlPassword:  "password",
		},
		DBName:        "postgres",
		WrkDir:        "wrkdir",
	}

	runCmd := fmt.Sprintf(PostgresRun, "postgres_db", 6432, "root", "password", "postgres")

	cmdSplits := strings.Split(runCmd, " ")

	cmdStr := cmdSplits[0]
	cmdArgs := cmdSplits[1:]

	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs, ".").Return([]byte(""), nil)

	err := RunContainer(dbInput)
	assert.NoError(t, err)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestRunContainer_Error(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	dbInput := models.DBInputs{
		DBMS:          "postgres",
		ContainerName: "postgres_db",
		ContainerPort: 6432,
		Postgres: models.PostgresDriver{		
			PsqlUser:      "root",
			PsqlPassword:  "password",
		},
		DBName:        "postgres",
		WrkDir:        "wrkdir",
	}

	runCmd := fmt.Sprintf(PostgresRun, "postgres_db", 6432, "root", "password", "postgres")

	cmdSplits := strings.Split(runCmd, " ")

	cmdStr := cmdSplits[0]
	cmdArgs := cmdSplits[1:]

	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs, ".").Return([]byte(""), errors.New("error in running docker run"))

	err := RunContainer(dbInput)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in running docker run")

	mockCmdsExecutor.AssertExpectations(t)
}

func TestRunContainer_SpecialCaseError(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	dbInput := models.DBInputs{
		DBMS:          "postgres",
		ContainerName: "postgres_db",
		ContainerPort: 6432,
		Postgres: models.PostgresDriver{		
			PsqlUser:      "root",
			PsqlPassword:  "password",
		},
		DBName:        "postgres",
		WrkDir:        "wrkdir",
	}

	runCmd := fmt.Sprintf(PostgresRun, "postgres_db", 6432, "root", "password", "postgres")

	cmdSplits := strings.Split(runCmd, " ")

	cmdStr := cmdSplits[0]
	cmdArgs := cmdSplits[1:]

	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs, ".").Return([]byte(`already in use by container`), errors.New("error in running docker run"))

	err := RunContainer(dbInput)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in running docker run")

	mockCmdsExecutor.AssertExpectations(t)
}

func TestRunPostgresContainer(t *testing.T) {

	dbInput := models.DBInputs{
		DBMS:          "postgres",
		ContainerName: "postgres_db",
		ContainerPort: 6432,
		Postgres: models.PostgresDriver{	
			PsqlUser:      "root",
			PsqlPassword:  "password",
		},
		DBName:        "postgres",
		WrkDir:        "wrkdir",
	}

	runCmd := fmt.Sprintf(PostgresRun, "postgres_db", 6432, "root", "password", "postgres")

	str := runPostgresContainer(dbInput)
	assert.Contains(t, runCmd, str)
} 

func TestRunMySQLContainer(t *testing.T) {

	dbInput := models.DBInputs{
		DBMS:          "mysql",
		ContainerName: "mysql_db",
		ContainerPort: 3306,
		MySQL: models.MySQLDriver{
			MysqlRootPassword: "secret",
			MysqlUser: "root",
			MysqlPassword: "password",
		},
		DBName:        "mysql",
		WrkDir:        "wrkdir",
	}

	runCmd := fmt.Sprintf(MySqlRun, "mysql_db", 3306, "secret", "root", "password", "mysql")

	str := runMySqlContainer(dbInput)
	assert.Contains(t, runCmd, str)
} 

func TestRunContainer_MYSQL(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	dbInput := models.DBInputs{
		DBMS:          "mysql",
		ContainerName: "mysql_db",
		ContainerPort: 3306,
		MySQL: models.MySQLDriver{
			MysqlRootPassword: "secret",
			MysqlUser: "root",
			MysqlPassword: "password",
		},
		DBName:        "mysql",
		WrkDir:        "wrkdir",
	}

	runCmd := fmt.Sprintf(MySqlRun, "mysql_db", 3306, "secret", "root", "password", "mysql")

	cmdSplits := strings.Split(runCmd, " ")

	cmdStr := cmdSplits[0]
	cmdArgs := cmdSplits[1:]

	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs, ".").Return([]byte(""), nil)

	err := RunContainer(dbInput)
	assert.NoError(t, err)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestRunContainer_Default(t *testing.T) {
	dbInput := models.DBInputs{
		DBMS:          "anotherDriver",
	}

	err := RunContainer(dbInput)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "driver not supported")
}
