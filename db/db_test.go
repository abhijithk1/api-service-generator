package db

import (
	"errors"
	"os"
	"testing"

	"github.com/abhijithk1/api-service-generator/common"
	"github.com/abhijithk1/api-service-generator/db/docker"
	"github.com/abhijithk1/api-service-generator/db/migrations"
	"github.com/abhijithk1/api-service-generator/db/query"
	"github.com/abhijithk1/api-service-generator/mocks"
	"github.com/abhijithk1/api-service-generator/models"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestInitialiseSQLC_Success(t *testing.T) {
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

func TestInitialiseSQLC_Error(t *testing.T) {
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

func TestGenerateSQLC_Success(t *testing.T) {
	mockExec := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockExec

	cmdStr := "sqlc"
	cmdArgs := []string{"generate"}

	wrkDir := "file"

	mockExec.On("ExecuteCmds", cmdStr, cmdArgs, wrkDir).Return([]byte(""), nil)

	err := generateSQLC(wrkDir)
	assert.NoError(t, err)
	mockExec.AssertExpectations(t)
}

func TestGenerateSQLC_Error(t *testing.T) {
	mockExec := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockExec

	cmdStr := "sqlc"
	cmdArgs := []string{"generate"}

	wrkDir := "file"

	mockExec.On("ExecuteCmds", cmdStr, cmdArgs, wrkDir).Return([]byte(""), errors.New("error in generating sqlc code"))

	err := generateSQLC(wrkDir)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in generating sqlc code")

	mockExec.AssertExpectations(t)
}

func TestEditSQLCYaml_Success(t *testing.T) {
	mockExec := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockExec
	wrkDir := "new-dir"
	sqlcYaml := models.SQLCYAML{
		Version: "1",
		Packages: []models.Packages{
			{
				Name:          "db",
				Path:          "./pkg/db",
				Schema:        "./pkg/db/migrations",
				Queries:       "./pkg/db/query/",
				Engine:        "postgresql",
				EmitInterface: false,
			},
		},
	}
	sqlcYamlMarshal, _ := yaml.Marshal(sqlcYaml)
	fileName := wrkDir + sqlcFileName
	mockExec.On("CreateFileAndItsContent", fileName, nil, string(sqlcYamlMarshal)).Return(nil)

	err := editSQLCYAML(wrkDir)
	assert.NoError(t, err)

	mockExec.AssertExpectations(t)
}

func TestEditSQLCYaml_ErrorWritingFile(t *testing.T) {
	mockExec := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockExec
	wrkDir := "new-dir"
	sqlcYaml := models.SQLCYAML{
		Version: "1",
		Packages: []models.Packages{
			{
				Name:          "db",
				Path:          "./pkg/db",
				Schema:        "./pkg/db/migrations",
				Queries:       "./pkg/db/query/",
				Engine:        "postgresql",
				EmitInterface: false,
			},
		},
	}

	sqlcYamlMarshal, _ := yaml.Marshal(sqlcYaml)
	fileName := wrkDir + sqlcFileName
	mockExec.On("CreateFileAndItsContent", fileName, nil, string(sqlcYamlMarshal)).Return(errors.New("error in editing sqlc.yaml"))

	err := editSQLCYAML(wrkDir)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error writing modified sqlc.yaml")

	mockExec.AssertExpectations(t)
}

func TestEditSQLCYaml_ErrorMarshalling(t *testing.T) {
	mockCommon := mocks.NewMockCommon()
	common.MarshalYAML = mockCommon.MarshalYAML

	wrkDir := "new-dir"
	sqlcYaml := models.SQLCYAML{
		Version: "1",
		Packages: []models.Packages{
			{
				Name:          "db",
				Path:          "./pkg/db",
				Schema:        "./pkg/db/migrations",
				Queries:       "./pkg/db/query/",
				Engine:        "postgresql",
				EmitInterface: false,
			},
		},
	}

	mockCommon.On("MarshalYAML", sqlcYaml).Return([]byte(""), errors.New("marshal error"))
	err := editSQLCYAML(wrkDir)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in marshalling the sqlc.yaml content")

	mockCommon.AssertExpectations(t)
}

func TestRunSQLC_Success(t *testing.T) {
	mockExec := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockExec

	cmdStr := "sqlc"
	cmdArgs := []string{"init"}

	wrkDir := "file"

	sqlcYaml := models.SQLCYAML{
		Version: "1",
		Packages: []models.Packages{
			{
				Name:          "db",
				Path:          "./pkg/db",
				Schema:        "./pkg/db/migrations",
				Queries:       "./pkg/db/query/",
				Engine:        "postgresql",
				EmitInterface: false,
			},
		},
	}
	sqlcYamlMarshal, _ := yaml.Marshal(sqlcYaml)
	fileName := wrkDir + sqlcFileName
	cmdStr1 := "sqlc"
	cmdArgs1 := []string{"generate"}
	common.MarshalYAML = yaml.Marshal

	mockExec.On("ExecuteCmds", cmdStr, cmdArgs, wrkDir).Return([]byte(""), nil)
	mockExec.On("CreateFileAndItsContent", fileName, nil, string(sqlcYamlMarshal)).Return(nil)
	mockExec.On("ExecuteCmds", cmdStr1, cmdArgs1, wrkDir).Return([]byte(""), nil)


	err := runSQLC(wrkDir)
	assert.NoError(t, err)

	mockExec.AssertExpectations(t)
}

func TestRunSQLC_GenerateError(t *testing.T) {
	mockExec := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockExec

	cmdStr := "sqlc"
	cmdArgs := []string{"init"}

	wrkDir := "file"

	sqlcYaml := models.SQLCYAML{
		Version: "1",
		Packages: []models.Packages{
			{
				Name:          "db",
				Path:          "./pkg/db",
				Schema:        "./pkg/db/migrations",
				Queries:       "./pkg/db/query/",
				Engine:        "postgresql",
				EmitInterface: false,
			},
		},
	}
	sqlcYamlMarshal, _ := yaml.Marshal(sqlcYaml)
	fileName := wrkDir + sqlcFileName
	cmdStr1 := "sqlc"
	cmdArgs1 := []string{"generate"}
	common.MarshalYAML = yaml.Marshal

	mockExec.On("ExecuteCmds", cmdStr, cmdArgs, wrkDir).Return([]byte(""), nil)
	mockExec.On("CreateFileAndItsContent", fileName, nil, string(sqlcYamlMarshal)).Return(nil)
	mockExec.On("ExecuteCmds", cmdStr1, cmdArgs1, wrkDir).Return([]byte(""), errors.New("error in generating sqlc code"))


	err := runSQLC(wrkDir)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in generating sqlc code")

	mockExec.AssertExpectations(t)
}

func TestRunSQLC_EditSQLCError(t *testing.T) {
	mockExec := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockExec

	cmdStr := "sqlc"
	cmdArgs := []string{"init"}

	wrkDir := "file"

	sqlcYaml := models.SQLCYAML{
		Version: "1",
		Packages: []models.Packages{
			{
				Name:          "db",
				Path:          "./pkg/db",
				Schema:        "./pkg/db/migrations",
				Queries:       "./pkg/db/query/",
				Engine:        "postgresql",
				EmitInterface: false,
			},
		},
	}
	sqlcYamlMarshal, _ := yaml.Marshal(sqlcYaml)
	fileName := wrkDir + sqlcFileName

	mockExec.On("ExecuteCmds", cmdStr, cmdArgs, wrkDir).Return([]byte(""), nil)
	mockExec.On("CreateFileAndItsContent", fileName, nil, string(sqlcYamlMarshal)).Return(errors.New("error in editing sqlc.yaml"))


	err := runSQLC(wrkDir)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in editing sqlc.yaml")

	mockExec.AssertExpectations(t)
}

func TestRunSQLC_InitialiseSQLCError(t *testing.T) {
	mockExec := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockExec

	cmdStr := "sqlc"
	cmdArgs := []string{"init"}

	wrkDir := "file"

	common.MarshalYAML = yaml.Marshal

	mockExec.On("ExecuteCmds", cmdStr, cmdArgs, wrkDir).Return([]byte(""), errors.New("error in initialising sqlc.yaml"))

	err := runSQLC(wrkDir)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in initialising sqlc.yaml")

	mockExec.AssertExpectations(t)
}

func TestSetup_Default(t *testing.T) {
	dbInputs := models.DBInputs{
		DBMS: "mysql",
	}
	Setup(dbInputs)
}

func TestSetup_Postgres(t *testing.T) {

	//mock docker client
	mockDocker := mocks.NewMockDocker()
	docker.DefaultDockerClient = mockDocker

	dbInputs := models.DBInputs{
		DBMS: "postgres",
		DBName: "database",
		WrkDir: "dir",
		ContainerName: "postgres_db",
		ContainerPort: 6432,
		PsqlUser: "user",
		PsqlPassword: "password",
		TableName: "table1",
	}
	
	mockDocker.On("RunContainer", dbInputs).Return(errors.New("error in running container"))

	Setup(dbInputs)

	mockDocker.AssertExpectations(t)

}

func TestConnectDB(t *testing.T) {
	mockExec := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockExec

	dbConnection := models.DBConnection{
		GoModule: "example",
		WrkDir: "dir",
	}

	fileName := dbConnection.WrkDir + connectionPath + "connection.go"
	mockExec.On("CreateFileAndItsContent", fileName, dbConnection, connection).Return(nil)

	err := connectDb(dbConnection)
	assert.NoError(t, err)

	mockExec.AssertExpectations(t)
}

func TestMainTest(t *testing.T) {
	mockExec := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockExec
	wrkDir := "dir"

	mainTestFileName := wrkDir + connectionPath + "main_test.go"
	mockExec.On("CreateFileAndItsContent", mainTestFileName, nil, mainTestContent).Return(nil)

	err := mainTest(wrkDir)
	assert.NoError(t, err)

	mockExec.AssertExpectations(t)
}

func TestSetupPostgres_Success(t *testing.T) {
	//mock cmd executor
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	//mock docker client
	mockDocker := mocks.NewMockDocker()
	docker.DefaultDockerClient = mockDocker

	//mock migration client
	mockMigration := mocks.NewMockMigration()
	migrations.DefaultMigrationClient = mockMigration

	//mock query client
	mockQuery := mocks.NewMockQuery()
	query.DefaultQueryClient = mockQuery

	dbInputs := models.DBInputs{
		DBMS: "postgres",
		DBName: "database",
		WrkDir: "dir",
		ContainerName: "postgres_db",
		ContainerPort: 6432,
		PsqlUser: "user",
		PsqlPassword: "password",
		TableName: "table1",
		GoModule: "example",
	}

	initSchema := models.InitSchema{
		TableName: dbInputs.TableName,
		WrkDir: dbInputs.WrkDir,
	}

	cmdStr := "sqlc"
	cmdArgs := []string{"init"}

	sqlcYaml := models.SQLCYAML{
		Version: "1",
		Packages: []models.Packages{
			{
				Name:          "db",
				Path:          "./pkg/db",
				Schema:        "./pkg/db/migrations",
				Queries:       "./pkg/db/query/",
				Engine:        "postgresql",
				EmitInterface: false,
			},
		},
	}
	sqlcYamlMarshal, _ := yaml.Marshal(sqlcYaml)
	sqlcFileName := dbInputs.WrkDir + sqlcFileName
	cmdStr1 := "sqlc"
	cmdArgs1 := []string{"generate"}
	common.MarshalYAML = yaml.Marshal

	dbConnection := models.DBConnection{
		GoModule: "example",
		WrkDir: "dir",
	}

	fileName := dbInputs.WrkDir + connectionPath + "connection.go"
	
	mainTestFileName := dbInputs.WrkDir + connectionPath + "main_test.go"
	
	mockDocker.On("RunContainer", dbInputs).Return(nil)
	mockMigration.On("Migration", dbInputs, initSchema).Return(nil)
	mockQuery.On("SetTableQuery", initSchema).Return(nil)
	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs, dbInputs.WrkDir).Return([]byte(""), nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", sqlcFileName, nil, string(sqlcYamlMarshal)).Return(nil)
	mockCmdsExecutor.On("ExecuteCmds", cmdStr1, cmdArgs1, dbInputs.WrkDir).Return([]byte(""), nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", fileName, dbConnection, connection).Return(nil)
	mockMigration.On("RunMigration", dbInputs).Return(nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", mainTestFileName, nil, mainTestContent).Return(nil)

	setupPostgres(dbInputs)

	mockCmdsExecutor.AssertExpectations(t)
	mockDocker.AssertExpectations(t)
	mockMigration.AssertExpectations(t)
	mockQuery.AssertExpectations(t)

}

func TestSetupPostgres_MainTestError(t *testing.T) {
	//mock cmd executor
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	//mock docker client
	mockDocker := mocks.NewMockDocker()
	docker.DefaultDockerClient = mockDocker

	//mock migration client
	mockMigration := mocks.NewMockMigration()
	migrations.DefaultMigrationClient = mockMigration

	//mock query client
	mockQuery := mocks.NewMockQuery()
	query.DefaultQueryClient = mockQuery

	dbInputs := models.DBInputs{
		DBMS: "postgres",
		DBName: "database",
		WrkDir: "dir",
		ContainerName: "postgres_db",
		ContainerPort: 6432,
		PsqlUser: "user",
		PsqlPassword: "password",
		TableName: "table1",
		GoModule: "example",
	}

	initSchema := models.InitSchema{
		TableName: dbInputs.TableName,
		WrkDir: dbInputs.WrkDir,
	}

	cmdStr := "sqlc"
	cmdArgs := []string{"init"}

	sqlcYaml := models.SQLCYAML{
		Version: "1",
		Packages: []models.Packages{
			{
				Name:          "db",
				Path:          "./pkg/db",
				Schema:        "./pkg/db/migrations",
				Queries:       "./pkg/db/query/",
				Engine:        "postgresql",
				EmitInterface: false,
			},
		},
	}
	sqlcYamlMarshal, _ := yaml.Marshal(sqlcYaml)
	sqlcFileName := dbInputs.WrkDir + sqlcFileName
	cmdStr1 := "sqlc"
	cmdArgs1 := []string{"generate"}
	common.MarshalYAML = yaml.Marshal
	dbConnection := models.DBConnection{
		GoModule: "example",
		WrkDir: "dir",
	}

	fileName := dbInputs.WrkDir + connectionPath + "connection.go"
	
	mainTestFileName := dbInputs.WrkDir + connectionPath + "main_test.go"
	
	mockDocker.On("RunContainer", dbInputs).Return(nil)
	mockMigration.On("Migration", dbInputs, initSchema).Return(nil)
	mockQuery.On("SetTableQuery", initSchema).Return(nil)
	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs, dbInputs.WrkDir).Return([]byte(""), nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", sqlcFileName, nil, string(sqlcYamlMarshal)).Return(nil)
	mockCmdsExecutor.On("ExecuteCmds", cmdStr1, cmdArgs1, dbInputs.WrkDir).Return([]byte(""), nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", fileName, dbConnection, connection).Return(nil)
	mockMigration.On("RunMigration", dbInputs).Return(nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", mainTestFileName, nil, mainTestContent).Return(errors.New("error in writing main_test.go"))

	setupPostgres(dbInputs)

	mockCmdsExecutor.AssertExpectations(t)
	mockDocker.AssertExpectations(t)
	mockMigration.AssertExpectations(t)
	mockQuery.AssertExpectations(t)

}

func TestSetupPostgres_RunMigrationError(t *testing.T) {
	//mock cmd executor
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	//mock docker client
	mockDocker := mocks.NewMockDocker()
	docker.DefaultDockerClient = mockDocker

	//mock migration client
	mockMigration := mocks.NewMockMigration()
	migrations.DefaultMigrationClient = mockMigration

	//mock query client
	mockQuery := mocks.NewMockQuery()
	query.DefaultQueryClient = mockQuery

	dbInputs := models.DBInputs{
		DBMS: "postgres",
		DBName: "database",
		WrkDir: "dir",
		ContainerName: "postgres_db",
		ContainerPort: 6432,
		PsqlUser: "user",
		PsqlPassword: "password",
		TableName: "table1",
		GoModule: "example",
	}

	initSchema := models.InitSchema{
		TableName: dbInputs.TableName,
		WrkDir: dbInputs.WrkDir,
	}

	cmdStr := "sqlc"
	cmdArgs := []string{"init"}

	sqlcYaml := models.SQLCYAML{
		Version: "1",
		Packages: []models.Packages{
			{
				Name:          "db",
				Path:          "./pkg/db",
				Schema:        "./pkg/db/migrations",
				Queries:       "./pkg/db/query/",
				Engine:        "postgresql",
				EmitInterface: false,
			},
		},
	}
	sqlcYamlMarshal, _ := yaml.Marshal(sqlcYaml)
	sqlcFileName := dbInputs.WrkDir + sqlcFileName
	cmdStr1 := "sqlc"
	cmdArgs1 := []string{"generate"}
	common.MarshalYAML = yaml.Marshal
	dbConnection := models.DBConnection{
		GoModule: "example",
		WrkDir: "dir",
	}

	fileName := dbInputs.WrkDir + connectionPath + "connection.go"
	
	mockDocker.On("RunContainer", dbInputs).Return(nil)
	mockMigration.On("Migration", dbInputs, initSchema).Return(nil)
	mockQuery.On("SetTableQuery", initSchema).Return(nil)
	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs, dbInputs.WrkDir).Return([]byte(""), nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", sqlcFileName, nil, string(sqlcYamlMarshal)).Return(nil)
	mockCmdsExecutor.On("ExecuteCmds", cmdStr1, cmdArgs1, dbInputs.WrkDir).Return([]byte(""), nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", fileName, dbConnection, connection).Return(nil)
	mockMigration.On("RunMigration", dbInputs).Return(errors.New("error running migration"))

	setupPostgres(dbInputs)

	mockCmdsExecutor.AssertExpectations(t)
	mockDocker.AssertExpectations(t)
	mockMigration.AssertExpectations(t)
	mockQuery.AssertExpectations(t)

}

func TestSetupPostgres_ConnectionDBError(t *testing.T) {
	//mock cmd executor
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	//mock docker client
	mockDocker := mocks.NewMockDocker()
	docker.DefaultDockerClient = mockDocker

	//mock migration client
	mockMigration := mocks.NewMockMigration()
	migrations.DefaultMigrationClient = mockMigration

	//mock query client
	mockQuery := mocks.NewMockQuery()
	query.DefaultQueryClient = mockQuery

	dbInputs := models.DBInputs{
		DBMS: "postgres",
		DBName: "database",
		WrkDir: "dir",
		ContainerName: "postgres_db",
		ContainerPort: 6432,
		PsqlUser: "user",
		PsqlPassword: "password",
		TableName: "table1",
		GoModule: "example",
	}

	initSchema := models.InitSchema{
		TableName: dbInputs.TableName,
		WrkDir: dbInputs.WrkDir,
	}

	cmdStr := "sqlc"
	cmdArgs := []string{"init"}

	sqlcYaml := models.SQLCYAML{
		Version: "1",
		Packages: []models.Packages{
			{
				Name:          "db",
				Path:          "./pkg/db",
				Schema:        "./pkg/db/migrations",
				Queries:       "./pkg/db/query/",
				Engine:        "postgresql",
				EmitInterface: false,
			},
		},
	}
	sqlcYamlMarshal, _ := yaml.Marshal(sqlcYaml)
	sqlcFileName := dbInputs.WrkDir + sqlcFileName
	cmdStr1 := "sqlc"
	cmdArgs1 := []string{"generate"}
	common.MarshalYAML = yaml.Marshal
	dbConnection := models.DBConnection{
		GoModule: "example",
		WrkDir: "dir",
	}

	fileName := dbInputs.WrkDir + connectionPath + "connection.go"
	
	mockDocker.On("RunContainer", dbInputs).Return(nil)
	mockMigration.On("Migration", dbInputs, initSchema).Return(nil)
	mockQuery.On("SetTableQuery", initSchema).Return(nil)
	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs, dbInputs.WrkDir).Return([]byte(""), nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", sqlcFileName, nil, string(sqlcYamlMarshal)).Return(nil)
	mockCmdsExecutor.On("ExecuteCmds", cmdStr1, cmdArgs1, dbInputs.WrkDir).Return([]byte(""), nil)
	mockCmdsExecutor.On("CreateFileAndItsContent", fileName, dbConnection, connection).Return(errors.New("Error in connecting DB"))

	setupPostgres(dbInputs)

	mockCmdsExecutor.AssertExpectations(t)
	mockDocker.AssertExpectations(t)
	mockMigration.AssertExpectations(t)
	mockQuery.AssertExpectations(t)

}

func TestSetupPostgres_RunSQLCError(t *testing.T) {
	//mock cmd executor
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	//mock docker client
	mockDocker := mocks.NewMockDocker()
	docker.DefaultDockerClient = mockDocker

	//mock migration client
	mockMigration := mocks.NewMockMigration()
	migrations.DefaultMigrationClient = mockMigration

	//mock query client
	mockQuery := mocks.NewMockQuery()
	query.DefaultQueryClient = mockQuery

	dbInputs := models.DBInputs{
		DBMS: "postgres",
		DBName: "database",
		WrkDir: "dir",
		ContainerName: "postgres_db",
		ContainerPort: 6432,
		PsqlUser: "user",
		PsqlPassword: "password",
		TableName: "table1",
	}

	initSchema := models.InitSchema{
		TableName: dbInputs.TableName,
		WrkDir: dbInputs.WrkDir,
	}

	cmdStr := "sqlc"
	cmdArgs := []string{"init"}
	
	mockDocker.On("RunContainer", dbInputs).Return(nil)
	mockMigration.On("Migration", dbInputs, initSchema).Return(nil)
	mockQuery.On("SetTableQuery", initSchema).Return(nil)
	mockCmdsExecutor.On("ExecuteCmds", cmdStr, cmdArgs, dbInputs.WrkDir).Return([]byte(""), errors.New("Error in running sqlc"))

	setupPostgres(dbInputs)

	mockCmdsExecutor.AssertExpectations(t)
	mockDocker.AssertExpectations(t)
	mockMigration.AssertExpectations(t)
	mockQuery.AssertExpectations(t)

}

func TestSetupPostgres_SetTableQueryError(t *testing.T) {
	//mock docker client
	mockDocker := mocks.NewMockDocker()
	docker.DefaultDockerClient = mockDocker

	//mock migration client
	mockMigration := mocks.NewMockMigration()
	migrations.DefaultMigrationClient = mockMigration

	//mock query client
	mockQuery := mocks.NewMockQuery()
	query.DefaultQueryClient = mockQuery

	dbInputs := models.DBInputs{
		DBMS: "postgres",
		DBName: "database",
		WrkDir: "dir",
		ContainerName: "postgres_db",
		ContainerPort: 6432,
		PsqlUser: "user",
		PsqlPassword: "password",
		TableName: "table1",
	}

	initSchema := models.InitSchema{
		TableName: dbInputs.TableName,
		WrkDir: dbInputs.WrkDir,
	}
	
	mockDocker.On("RunContainer", dbInputs).Return(nil)
	mockMigration.On("Migration", dbInputs, initSchema).Return(nil)
	mockQuery.On("SetTableQuery", initSchema).Return(errors.New("error in setting table query"))

	setupPostgres(dbInputs)

	mockDocker.AssertExpectations(t)
	mockMigration.AssertExpectations(t)
	mockQuery.AssertExpectations(t)

}

func TestSetupPostgres_SMigrationError(t *testing.T) {
	//mock docker client
	mockDocker := mocks.NewMockDocker()
	docker.DefaultDockerClient = mockDocker

	//mock migration client
	mockMigration := mocks.NewMockMigration()
	migrations.DefaultMigrationClient = mockMigration

	dbInputs := models.DBInputs{
		DBMS: "postgres",
		DBName: "database",
		WrkDir: "dir",
		ContainerName: "postgres_db",
		ContainerPort: 6432,
		PsqlUser: "user",
		PsqlPassword: "password",
		TableName: "table1",
	}

	initSchema := models.InitSchema{
		TableName: dbInputs.TableName,
		WrkDir: dbInputs.WrkDir,
	}
	
	mockDocker.On("RunContainer", dbInputs).Return(nil)
	mockMigration.On("Migration", dbInputs, initSchema).Return(errors.New("error in migrating"))

	setupPostgres(dbInputs)

	mockDocker.AssertExpectations(t)
	mockMigration.AssertExpectations(t)

}

func TestSetupPostgres_RunningContainer(t *testing.T) {
	//mock docker client
	mockDocker := mocks.NewMockDocker()
	docker.DefaultDockerClient = mockDocker

	dbInputs := models.DBInputs{
		DBMS: "postgres",
		DBName: "database",
		WrkDir: "dir",
		ContainerName: "postgres_db",
		ContainerPort: 6432,
		PsqlUser: "user",
		PsqlPassword: "password",
		TableName: "table1",
	}
	
	mockDocker.On("RunContainer", dbInputs).Return(errors.New("error in running container"))

	setupPostgres(dbInputs)

	mockDocker.AssertExpectations(t)

}
