package db

import (
	"fmt"
	"os"

	"github.com/abhijithk1/api-service-generator/common"
	"github.com/abhijithk1/api-service-generator/db/docker"
	"github.com/abhijithk1/api-service-generator/db/migrations"
	"github.com/abhijithk1/api-service-generator/db/query"
	"github.com/abhijithk1/api-service-generator/models"
	"gopkg.in/yaml.v3"
)

var sqlcFileName = "sqlc.yaml"

type SQLCYAML struct {
	Version  string     `yaml:"version"`
	Packages []Packages `yaml:"packages"`
}

type Packages struct {
	Name          string `yaml:"name"`
	Path          string `yaml:"path"`
	Queries       string `yaml:"queries"`
	Schema        string `yaml:"schema"`
	Engine        string `yaml:"engine"`
	EmitInterface bool   `yaml:"emit_interface"`
}

func runSQLC(wrkDir string) (err error) {
	err = initialiseSQLC(wrkDir)
	if err != nil {
		return
	}

	err = editSQLCYAML(wrkDir)
	if err != nil {
		return
	}

	err = generateSQLC(wrkDir)
	if err != nil {
		return
	}

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

func editSQLCYAML(wrkDir string) error {

	sqlcYaml := SQLCYAML{
		Version: "1",
		Packages: []Packages{
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

	sqlcYamlMarshal, err := yaml.Marshal(sqlcYaml)
	if err != nil {
		return fmt.Errorf("error in marshalling the sqlc.yaml content")
	}

	sqlcFileName = wrkDir + sqlcFileName
	// Write the modified content back to the file
	err = os.WriteFile(sqlcFileName, sqlcYamlMarshal, 0644)
	if err != nil {
		return fmt.Errorf("error writing modified sqlc.yaml: %w", err)
	}

	return nil
}

func Setup(dbInputs models.DBInputs) {
	switch dbInputs.DBMS {
	case "postgres":
		setupPostgres(dbInputs)
	}
}

func setupPostgres(dbInputs models.DBInputs) {
	fmt.Println("*** Database setup ***")

	fmt.Println("\n*** Running Postgres in Docker Container ***")
	err := docker.RunPostgresContainer(dbInputs)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	fmt.Println("\n*** Postgres is Successfully Running in Docker Container `postgres_db` ***")

	initSchema := models.InitSchema{
		TableName: dbInputs.DBName,
	}

	err = migrations.PostgresMigration(dbInputs, initSchema)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	err = query.SetTableQuery(initSchema)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	err = runSQLC(dbInputs.WrkDir)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
}
