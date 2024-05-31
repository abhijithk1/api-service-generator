package migrations

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
	"github.com/stretchr/testify/require"
)

func TestMain (m *testing.M) {
	os.Exit(m.Run())
}
func TestInitialiseMigration_Success(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()	
	common.DefaultExecutor = mockCmdsExecutor

	runCmd := "migrate create -ext sql -dir file/pkg/db/migrations -seq init_schema"
	cmdSplits := strings.Split(runCmd, " ")

	cmdStr := cmdSplits[0]
	cmdArgs := cmdSplits[1:]

	wrkDir := "file"

	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs, ".").Return([]byte(""), nil)

	err := initialiseMigration(wrkDir)
	assert.NoError(t, err)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestInitialiseMigration_Error(t *testing.T) {
	mockCmdsExec := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExec

	wrkDir := "newfile"
	runCmd := "migrate create -ext sql -dir newfile/pkg/db/migrations -seq init_schema"

	cmdSplits := strings.Split(runCmd, " ")

	cmdStr := cmdSplits[0]
	cmdArgs := cmdSplits[1:]


	mockCmdsExec.On("ExecuteCmds", cmdStr, cmdArgs, ".").Return([]byte(""), errors.New("error in initialising"))

	err := initialiseMigration(wrkDir)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in initialising")

	mockCmdsExec.AssertExpectations(t)
}

func TestRunMigration(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	dbInputs := models.DBInputs{
		DBMS: "postgres",
		PsqlUser: "root",
		PsqlPassword: "password",
		DBName: "database",
		WrkDir: "file",
	}

	psqlUrl := fmt.Sprintf(PostgresqlUrl, dbInputs.DBMS, dbInputs.PsqlUser, dbInputs.PsqlPassword, dbInputs.DBName)

	fileName := dbInputs.WrkDir + migrationUpFilePath + "migrate.go"

	mockCmdsExecutor.On("CreateFileAndItsContent", fileName, models.Migration{DatabaseURL: psqlUrl}, migrateUp_content).Return(nil)
	err := RunMigration(dbInputs)

	require.NoError(t, err)

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
	mockCmdsExecutor.On("CreateFileAndItsContent", migrationDirectoryPath+fileName, initSchema, init_schema_up).Return(errors.New("error in writing init schema up sql file"))
	err = writeSchemaUpFile(initSchema)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in writing init schema up sql file")

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
	mockCmdsExecutor.On("CreateFileAndItsContent", migrationDirectoryPath+fileName, initSchema, init_schema_down).Return(errors.New("error in writing init schema down sql file"))
	err = writeSchemaDownFile(initSchema)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in writing init schema down sql file")

	mockCmdsExecutor.AssertExpectations(t)
}

func TestMigration_Success(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	runCmd := "migrate create -ext sql -dir wrkdir/pkg/db/migrations -seq init_schema"
	cmdSplits := strings.Split(runCmd, " ")
	cmdStr := cmdSplits[0]
	cmdArgs := cmdSplits[1:]

	fileName := "/init_schema_up.sql"
	fileName1 := "/init_schema_down.sql"
	migrationUpFileName = fileName
	migrationDownFileName = fileName1

	schemaUp := `CREATE TABLE IF NOT EXISTS dummy(
		id PRIMARY KEY,
		name VARCHAR(255)
	);`

	schemaDown := `DROP TABLE IF EXISTS dummy;`

	init_schema_up = schemaUp
	init_schema_down = schemaDown

	initSchema := models.InitSchema{
		TableName: "dummy",
	}

	dbInput := models.DBInputs{
		PsqlUser:     "root",
		PsqlPassword: "password",
		DBName:       "postgres",
		WrkDir:       "wrkdir",
	}

	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs, ".").Return([]byte(""), nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", migrationDirectoryPath+fileName, initSchema, init_schema_up).Return(nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", migrationDirectoryPath+fileName1, initSchema, init_schema_down).Return(nil)

	err := Migration(dbInput, initSchema)
	assert.NoError(t, err)
	mockCmdsExecutor.AssertExpectations(t)

}

func TestMigration_InitialiseError(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	runCmd := "migrate create -ext sql -dir wrkdir1/pkg/db/migrations -seq init_schema"
	cmdSplits := strings.Split(runCmd, " ")
	cmdStr := cmdSplits[0]
	cmdArgs := cmdSplits[1:]

	initSchema := models.InitSchema{
		TableName: "dummy",
	}

	dbInput := models.DBInputs{
		PsqlUser:     "root",
		PsqlPassword: "password",
		DBName:       "postgres",
		WrkDir:       "wrkdir1",
	}

	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs, ".").Return([]byte(""), errors.New("error in creating migration"))

	err := Migration(dbInput, initSchema)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in creating migration")

	mockCmdsExecutor.AssertExpectations(t)

}

func TestMigration_WriteSchemaUpFIleError(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	runCmd := "migrate create -ext sql -dir wrkdir2/pkg/db/migrations -seq init_schema"
	cmdSplits := strings.Split(runCmd, " ")
	cmdStr := cmdSplits[0]
	cmdArgs := cmdSplits[1:]

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

	dbInput := models.DBInputs{
		PsqlUser:     "root",
		PsqlPassword: "password",
		DBName:       "postgres",
		WrkDir:       "wrkdir2",
	}

	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs, ".").Return([]byte(""), nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", migrationDirectoryPath+fileName, initSchema, init_schema_up).Return(errors.New("error writing init schema up sql file"))

	err := Migration(dbInput, initSchema)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error writing init schema up sql file")

	mockCmdsExecutor.AssertExpectations(t)

}

func TestMigration_WriteSchemaDownFileError(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	runCmd := "migrate create -ext sql -dir wrkdir3/pkg/db/migrations -seq init_schema"
	cmdSplits := strings.Split(runCmd, " ")
	cmdStr := cmdSplits[0]
	cmdArgs := cmdSplits[1:]

	fileName := "/init_schema_up.sql"
	fileName1 := "/init_schema_down.sql"
	migrationUpFileName = fileName
	migrationDownFileName = fileName1

	schemaUp := `CREATE TABLE IF NOT EXISTS dummy(
		id PRIMARY KEY,
		name VARCHAR(255)
	);`

	schemaDown := `DROP TABLE IF EXISTS dummy;`

	init_schema_up = schemaUp
	init_schema_down = schemaDown

	initSchema := models.InitSchema{
		TableName: "dummy",
	}

	dbInput := models.DBInputs{
		PsqlUser:     "root",
		PsqlPassword: "password",
		DBName:       "postgres",
		WrkDir:       "wrkdir3",
	}

	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs, ".").Return([]byte(""), nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", migrationDirectoryPath+fileName, initSchema, init_schema_up).Return(nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", migrationDirectoryPath+fileName1, initSchema, init_schema_down).Return(errors.New("error in writing init schema down sql file"))

	err := Migration(dbInput, initSchema)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in writing init schema down sql file")

	mockCmdsExecutor.AssertExpectations(t)

}
