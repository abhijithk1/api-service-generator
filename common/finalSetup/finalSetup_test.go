package finalsetup

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/abhijithk1/api-service-generator/common"
	"github.com/abhijithk1/api-service-generator/mocks"
	"github.com/abhijithk1/api-service-generator/models"
	"github.com/stretchr/testify/assert"
)

func TestMain (m *testing.M) {
	os.Exit(m.Run())
}

func TestCreateMainFile(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	apiInputs := models.APIInputs{
		WrkDir: "dir",
		APIGroup: "dummy",
		APIGroupTitle: "Dummy",
		GoModule: "example",
	}
	filename := apiInputs.WrkDir + "/main.go"

	mockCmdsExecutor.On("CreateFileAndItsContent", filename, apiInputs, mainContent).Return(nil)
	err := createMainFile(apiInputs)
	assert.NoError(t, err)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestCreateMakeFile(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	wrkDir := "dir"
	filename := wrkDir + "/Makefile"

	mockCmdsExecutor.On("CreateFileAndItsContent", filename, nil, makeFileContent).Return(nil)

	err := createMakeFile(wrkDir)
	assert.NoError(t, err)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestCreateENVFile_Postgres(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	dbInputs := models.DBInputs{
		Postgres: models.PostgresDriver{
			PsqlUser: "user",
			PsqlPassword: "password",
		},
		ContainerPort: 6432,
		DBMS: "postgres",
		DBName: "data_db",
		WrkDir: "dir",
	}
	filename := dbInputs.WrkDir + "/app.env"

	newEnvFile := fmt.Sprintf(envFile, PostgresDBSource)

	mockCmdsExecutor.On("CreateFileAndItsContent", filename, dbInputs, newEnvFile).Return(nil)

	err := createENVFile(dbInputs)
	assert.NoError(t, err)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestCreateENVFile_MySql(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	dbInputs := models.DBInputs{
		MySQL: models.MySQLDriver{
			MysqlRootPassword: "secret",
			MysqlUser: "user",
			MysqlPassword: "password",
		},
		ContainerPort: 6432,
		DBMS: "mysql",
		DBName: "data_db",
		WrkDir: "dir",
	}
	filename := dbInputs.WrkDir + "/app.env"

	newEnvFile := fmt.Sprintf(envFile, MysqlDBSource)

	mockCmdsExecutor.On("CreateFileAndItsContent", filename, dbInputs, newEnvFile).Return(nil)

	err := createENVFile(dbInputs)
	assert.NoError(t, err)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestCreateENVFile_Default(t *testing.T) {

	dbInputs := models.DBInputs{
		DBMS: "another driver",
	}

	err := createENVFile(dbInputs)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "driver not supported")
}

func TestFinalSetup_Success(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	apiInputs := models.APIInputs{
		WrkDir: "dir",
		APIGroup: "dummy",
		APIGroupTitle: "Dummy",
		GoModule: "example",
	}
	dbInputs := models.DBInputs{
		Postgres: models.PostgresDriver{		
			PsqlPassword: "password",
			PsqlUser: "user",
		},
		ContainerPort: 6432,
		DBMS: "postgres",
		DBName: "data_db",
		WrkDir: "dir",
	}
	mainFilename := apiInputs.WrkDir + "/main.go"
	makeFilename := apiInputs.WrkDir + "/Makefile"
	envFilename := dbInputs.WrkDir + "/app.env"
	httpFileName := apiInputs.WrkDir + "/api.http"

	newEnvFile := fmt.Sprintf(envFile, PostgresDBSource)
	
	mockCmdsExecutor.On("CreateFileAndItsContent", mainFilename, apiInputs, mainContent).Return(nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", envFilename, dbInputs, newEnvFile).Return(nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", makeFilename, nil, makeFileContent).Return(nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", httpFileName, apiInputs, api_HTTP).Return(nil)

	FinalSetup(apiInputs, dbInputs)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestFinalSetup_APIHTTPFileError(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	apiInputs := models.APIInputs{
		WrkDir: "dir",
		APIGroup: "dummy",
		APIGroupTitle: "Dummy",
		GoModule: "example",
	}
	dbInputs := models.DBInputs{
		Postgres: models.PostgresDriver{		
			PsqlUser: "user",
			PsqlPassword: "password",
		},
		ContainerPort: 6432,
		DBMS: "postgres",
		DBName: "data_db",
		WrkDir: "dir",
	}
	mainFilename := apiInputs.WrkDir + "/main.go"
	makeFilename := apiInputs.WrkDir + "/Makefile"
	envFilename := dbInputs.WrkDir + "/app.env"
	httpFileName := apiInputs.WrkDir + "/api.http"

	newEnvFile := fmt.Sprintf(envFile, PostgresDBSource)
	
	mockCmdsExecutor.On("CreateFileAndItsContent", mainFilename, apiInputs, mainContent).Return(nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", envFilename, dbInputs, newEnvFile).Return(nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", makeFilename, nil, makeFileContent).Return(nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", httpFileName, apiInputs, api_HTTP).Return(errors.New("error in creating api.http"))

	FinalSetup(apiInputs, dbInputs)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestFinalSetup_MakefileError(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	apiInputs := models.APIInputs{
		WrkDir: "dir",
		APIGroup: "dummy",
		APIGroupTitle: "Dummy",
		GoModule: "example",
	}
	dbInputs := models.DBInputs{
		Postgres: models.PostgresDriver{		
			PsqlUser: "user",
			PsqlPassword: "password",
		},
		ContainerPort: 6432,
		DBMS: "postgres",
		DBName: "data_db",
		WrkDir: "dir",
	}
	mainFilename := apiInputs.WrkDir + "/main.go"
	makeFilename := apiInputs.WrkDir + "/Makefile"
	envFilename := dbInputs.WrkDir + "/app.env"

	newEnvFile := fmt.Sprintf(envFile, PostgresDBSource)

	mockCmdsExecutor.On("CreateFileAndItsContent", mainFilename, apiInputs, mainContent).Return(nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", envFilename, dbInputs, newEnvFile).Return(nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", makeFilename, nil, makeFileContent).Return(errors.New("error in creating MakeFile"))

	FinalSetup(apiInputs, dbInputs)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestFinalSetup_EnvFileError(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	apiInputs := models.APIInputs{
		WrkDir: "dir",
		APIGroup: "dummy",
		APIGroupTitle: "Dummy",
		GoModule: "example",
	}
	dbInputs := models.DBInputs{
		Postgres: models.PostgresDriver{	
			PsqlUser: "user",
			PsqlPassword: "password",
		},
		ContainerPort: 6432,
		DBMS: "postgres",
		DBName: "data_db",
		WrkDir: "dir",
	}
	mainFilename := apiInputs.WrkDir + "/main.go"
	envFilename := dbInputs.WrkDir + "/app.env"

	newEnvFile := fmt.Sprintf(envFile,PostgresDBSource)

	mockCmdsExecutor.On("CreateFileAndItsContent", mainFilename, apiInputs, mainContent).Return(nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", envFilename, dbInputs, newEnvFile).Return(errors.New("error in creating app.env"))

	FinalSetup(apiInputs, dbInputs)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestFinalSetup_MainFileError(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	apiInputs := models.APIInputs{
		WrkDir: "dir",
		APIGroup: "dummy",
		APIGroupTitle: "Dummy",
		GoModule: "example",
	}

	dbInputs := models.DBInputs{}
	filename := apiInputs.WrkDir + "/main.go"

	mockCmdsExecutor.On("CreateFileAndItsContent", filename, apiInputs, mainContent).Return(errors.New("error in creating main.go"))

	FinalSetup(apiInputs, dbInputs)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestCreateAPIHTTPFile(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	apiInputs := models.APIInputs{
		WrkDir: "dir",
		APIGroup: "dummy",
		APIGroupTitle: "Dummy",
		GoModule: "example",
	}

	fileName := apiInputs.WrkDir + "/api.http"
	mockCmdsExecutor.On("CreateFileAndItsContent", fileName, apiInputs, api_HTTP).Return(nil)

	err := createAPIHTTPFile(apiInputs)

	assert.NoError(t, err)

	mockCmdsExecutor.AssertExpectations(t)

}
