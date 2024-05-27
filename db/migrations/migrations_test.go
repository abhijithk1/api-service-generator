package migrations

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

func TestInitialiseMigration_Success(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	runCmd := "migrate create -ext sql -dir pkg/db/migrations -seq init_schema"
	cmdSplits := strings.Split(runCmd, " ")

	cmdStr := cmdSplits[0]
	cmdArgs := cmdSplits[1:]

	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs).Return([]byte(""), nil)

	err := initialiseMigration()
	assert.NoError(t, err)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestInitialiseMigration_Error(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	runCmd := "migrate create -ext sql -dir pkg/db/migrations -seq init_schema"

	cmdSplits := strings.Split(runCmd, " ")

	cmdStr := cmdSplits[0]
	cmdArgs := cmdSplits[1:]

	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs).Return([]byte(""), errors.New("error in initialising"))

	err := initialiseMigration()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in initialising")

	mockCmdsExecutor.AssertExpectations(t)
}

func TestMigrationUp_Success(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	dbInput := models.DBInputs{
		PsqlUser:     "root",
		PsqlPassword: "password",
		DBName:       "postgres",
	}

	runCmd := fmt.Sprintf(MigrateUp, "postgresql", "root", "password", "postgres")

	cmdSplits := strings.Split(runCmd, " ")

	cmdStr := cmdSplits[0]
	cmdArgs := cmdSplits[1:]

	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs).Return([]byte(""), nil)

	err := migrationUp(dbInput)
	assert.NoError(t, err)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestMigrationUp_Error(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	dbInput := models.DBInputs{
		PsqlUser:     "root",
		PsqlPassword: "password",
		DBName:       "postgres",
	}

	runCmd := fmt.Sprintf(MigrateUp, "postgresql", "root", "password", "postgres")

	cmdSplits := strings.Split(runCmd, " ")

	cmdStr := cmdSplits[0]
	cmdArgs := cmdSplits[1:]

	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs).Return([]byte(""), errors.New("error in migrating"))

	err := migrationUp(dbInput)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in migrating")

	mockCmdsExecutor.AssertExpectations(t)
}

func TestWriteSchemaUpFile(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	fileName := "/init_schema_up.sql"
	migrationUpFileName = fileName
	schemaUp := `CREATE TABLE IF NOT EXISTS dummy(
		id PRIMARY KEY,
		name VARCHAR(255)
	);`

	init_schema_up = schemaUp


	initSchema := models.InitSchema{
		TableName: "dummy",
	}

	mockCmdsExecutor.On("CreateFileAndItsContent", migrationDirectoryPath+fileName, initSchema, init_schema_up).Return(nil)

	err := writeSchemaUpFile(initSchema)
	assert.NoError(t, err)
	mockCmdsExecutor.AssertExpectations(t)


	//Error case
	mockCmdsExecutor = mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor
	mockCmdsExecutor.On("CreateFileAndItsContent", migrationDirectoryPath+fileName, initSchema, init_schema_up).Return(errors.New("error in migrating up"))
	err = writeSchemaUpFile(initSchema)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in migrating up")

	mockCmdsExecutor.AssertExpectations(t)
}

func TestWriteSchemaDownFile(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	fileName := "/init_schema_down.sql"
	migrationDownFileName = fileName
	schemaDown := `DROP TABLE IF EXISTS dummy;`

	init_schema_down = schemaDown


	initSchema := models.InitSchema{
		TableName: "dummy",
	}

	mockCmdsExecutor.On("CreateFileAndItsContent", migrationDirectoryPath+fileName, initSchema, init_schema_down).Return(nil)

	err := writeSchemaDownFile(initSchema)
	assert.NoError(t, err)

	mockCmdsExecutor.AssertExpectations(t)

	//Error case
	mockCmdsExecutor = mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor
	mockCmdsExecutor.On("CreateFileAndItsContent", migrationDirectoryPath+fileName, initSchema, init_schema_down).Return(errors.New("error in migrating down"))
	err = writeSchemaDownFile(initSchema)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in migrating down")

	mockCmdsExecutor.AssertExpectations(t)
}