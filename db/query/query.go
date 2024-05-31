package query

import (
	"github.com/abhijithk1/api-service-generator/common"
	"github.com/abhijithk1/api-service-generator/models"
)

var queryDirectoryPath = "/pkg/db/query/"

var table_sql = `-- Generated using API Service Generator

-- name: List{{.TableName}} :many
SELECT * FROM {{.TableName}};
`

func SetTableQuery(initSchema models.InitSchema) (err error) {
	tableQueryFileName := initSchema.WrkDir + queryDirectoryPath + initSchema.TableName + ".sql"
	err = common.CreateFileAndItsContent(tableQueryFileName, initSchema, table_sql)
	if err != nil {
		return
	}
	return nil
}
