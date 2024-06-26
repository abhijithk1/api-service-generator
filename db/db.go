package db

import (
	"fmt"

	"github.com/abhijithk1/api-service-generator/common"
	"github.com/abhijithk1/api-service-generator/db/docker"
	"github.com/abhijithk1/api-service-generator/db/migrations"
	"github.com/abhijithk1/api-service-generator/db/query"
	"github.com/abhijithk1/api-service-generator/models"
)

var (
	sqlcFileName   = "/sqlc.yaml"
	connectionPath = "/pkg/db/"
)

func runSQLC(driver, wrkDir string) (err error) {
	err = initialiseSQLC(wrkDir)
	if err != nil {
		return
	}
	fmt.Println("\n\n*** Successfully initialised SLQC `postgres_db` ***")

	err = editSQLCYAML(driver, wrkDir)
	if err != nil {
		return
	}
	fmt.Println("\n\n*** Updated sqlc file ***")

	err = generateSQLC(wrkDir)
	if err != nil {
		return
	}
	fmt.Println("\n\n*** Successfully Generated SQLC ***")

	return nil
}

func initialiseSQLC(wrkDir string) (err error) {
	_, err = common.ExecuteCmds("sqlc", []string{"init"}, wrkDir)
	if err != nil {
		return
	}

	return nil
}

func generateSQLC(wrkDir string) (err error) {
	_, err = common.ExecuteCmds("sqlc", []string{"generate"}, wrkDir)
	if err != nil {
		return
	}

	return nil
}

func editSQLCYAML(driver, wrkDir string) error {
	sqlcYaml := models.SQLCYAML{
		Version: "1",
		Packages: []models.Packages{
			{
				Name:          "db",
				Path:          "./pkg/db",
				Schema:        "./pkg/db/migrations",
				Queries:       "./pkg/db/query/",
				EmitInterface: false,
			},
		},
	}

	sqlcEngine(driver, &sqlcYaml)

	sqlcYamlMarshal, err := common.MarshalYAML(sqlcYaml)
	if err != nil {
		return fmt.Errorf("error in marshalling the sqlc.yaml content")
	}

	fileName := wrkDir + sqlcFileName
	err = common.CreateFileAndItsContent(fileName, nil, string(sqlcYamlMarshal))
	if err != nil {
		return fmt.Errorf("error writing modified sqlc.yaml: %w", err)
	}

	return nil
}

// func Setup(dbInputs models.DBInputs) error{
// 	switch dbInputs.DBMS {
// 	case "postgres":
// 		return setupPostgres(dbInputs)
// 	default:
// 		fmt.Printf("The DBMS %s is not Supported currently", dbInputs.DBMS)
// 	}

// 	return nil
// }

func Setup(dbInputs models.DBInputs) (err error) {
	fmt.Println("*** Database setup ***")

	fmt.Println("\n*** Running Postgres in Docker Container ***")
	err = docker.RunContainer(dbInputs)
	if err != nil {
		fmt.Println("\nError : ", err)
		return
	}

	fmt.Println("\n\n*** Postgres is Successfully Running in Docker Container `postgres_db` ***")

	initSchema := models.InitSchema{
		TableName: dbInputs.TableName,
		WrkDir:    dbInputs.WrkDir,
	}

	err = migrations.Migration(dbInputs, initSchema)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	fmt.Println("\n\n*** Successfully Migrated ***")

	err = query.SetTableQuery(initSchema)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	fmt.Println("\n\n*** Query are successfully written ***")

	err = runSQLC(dbInputs.DBMS, dbInputs.WrkDir)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	fmt.Println("\n\n*** Successfully setup SQLC ***")

	err = connectDb(dbInputs)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("\n\n*** Successfully setup db connection file ***")

	err = migrations.RunMigration(dbInputs)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("\n\n*** Successfully setup migration file ***")

	err = mainTest(dbInputs.WrkDir)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	return nil
}

const connection = `// Generated BY API Service Generator

package db

import (
	"database/sql"
	util "{{.GoModule}}/{{.WrkDir}}/utils"

	"github.com/IBM/alchemy-logging/src/go/alog"
	_ "{{.DriverPackage}}"
)

var ch = alog.UseChannel("MAIN")

func GetConnection() *sql.DB {

	driver := util.GetAppConfig().DBDriver
	source := util.GetAppConfig().DBSource
	conn, err := sql.Open(driver, source)
	if err != nil {
		panic("unable to open database connection")
	}
	if conn.Ping() != nil {
		panic("unable to ping database")
	} else {
		ch.Log(alog.INFO, "connected to Database")
	}
	return conn
}
`

func connectDb(dbConnection models.DBInputs) error {
	fileName := dbConnection.WrkDir + connectionPath + "connection.go"
	return common.CreateFileAndItsContent(fileName, dbConnection, connection)

}

const mainTestContent = `/*
Generated By API Generator
*/

package db

import (
	"os"
	"testing"
)

var TestQueries *Queries

func TestMain(m *testing.M) {

	conn := GetConnection()
	TestQueries = New(conn)

	os.Exit(m.Run())
}
`

func mainTest(wrkDir string) error {
	fileName := wrkDir + connectionPath + "main_test.go"
	return common.CreateFileAndItsContent(fileName, nil, mainTestContent)
}

func sqlcEngine(driver string, sqlc *models.SQLCYAML) {
	switch driver {
	case "postgres":
		sqlc.Packages[0].Engine = "postgresql"
	case "mysql":
		sqlc.Packages[0].Engine = "mysql"
	}
}
