package migrations

import (
	"fmt"
	"strings"

	"github.com/abhijithk1/api-service-generator/common"
	"github.com/abhijithk1/api-service-generator/models"
)

var (
	MigrateCreateTemplate  = `migrate create -ext sql -dir pkg/db/migrations -seq init_schema`
	MigrateUp              = `migrate -path pkg/db/migrations -database "%s://%s:%s@localhost:5432/%s?sslmode=disable" -verbose up`
	MigrateDown            = `migrate -path pkg/db/migrations -database "%s://%s:%s@localhost:5432/%s?sslmode=disable" -verbose down`
	migrationDirectoryPath = "/pkg/db/migrations"
	migrationUpFileName    = "/000001_init_schema_up.sql"
	migrationDownFileName  = "/000001_init_schema_down.sql"
)

func PostgresMigration(dbInputs models.DBInputs, initSchema models.InitSchema) (err error) {
	err = initialiseMigration(dbInputs.WrkDir)
	if err != nil {
		return
	}

	err = writeSchemaUpFile(initSchema)
	if err != nil {
		return
	}

	err = writeSchemaDownFile(initSchema)
	if err != nil {
		return
	}

	err = migrationUp(dbInputs)
	if err != nil {
		return
	}

	return nil
}

func initialiseMigration(wrkDir string) (err error) {
	cmdSplits := strings.Split(MigrateCreateTemplate, " ")

	_, err = common.ExecuteCmds(cmdSplits[0], cmdSplits[1:], wrkDir)
	if err != nil {
		return
	}

	return nil
}

func migrationUp(dbInputs models.DBInputs) error {
	MigrateUp = fmt.Sprintf(MigrateUp, "postgresql", dbInputs.PsqlUser, dbInputs.PsqlPassword, dbInputs.DBName)
	migrateCmds := strings.Split(MigrateUp, " ")

	_, err := common.ExecuteCmds(migrateCmds[0], migrateCmds[1:], dbInputs.WrkDir)
	if err != nil {
		return err
	}

	return nil
}

var init_schema_up = `
/*
Generated using API Service Generator
*/

CREATE TABLE IF NOT EXISTS {{.TableName}} (
	id PRIMARY KEY,
	name VARCHAR(255)
	);
`

var init_schema_down = `
/*
Generated using API Service Generator
*/

DROP TABLE IF EXISTS {{.TableName}};
`

func writeSchemaUpFile(initSchema models.InitSchema) error {
	return common.CreateFileAndItsContent(migrationDirectoryPath+migrationUpFileName, initSchema, init_schema_up)
}

func writeSchemaDownFile(initSchema models.InitSchema) error {
	return common.CreateFileAndItsContent(migrationDirectoryPath+migrationDownFileName, initSchema, init_schema_down)
}
