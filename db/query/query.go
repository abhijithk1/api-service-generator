package query

import (
	"github.com/abhijithk1/api-service-generator/common"
	"github.com/abhijithk1/api-service-generator/models"
)

var queryDirectoryPath = "/pkg/db/query/"

type QueryInterface interface {
	SetTableQuery(initSchema models.InitSchema) (err error)
}

var DefaultQueryClient QueryInterface = &QueryClient{}

type QueryClient struct{}

var table_sql = `-- Generated using API Service Generator

-- name: List{{.TableName}} :many
SELECT * FROM {{.TableName}};
`

func (q * QueryClient) SetTableQuery(initSchema models.InitSchema) (err error) {
	tableQueryFileName := initSchema.WrkDir + queryDirectoryPath + initSchema.TableName + ".sql"
	err = common.CreateFileAndItsContent(tableQueryFileName, initSchema, table_sql)
	if err != nil {
		return
	}
	return nil
}

func SetTableQuery(initSchema models.InitSchema) (err error) {
	return DefaultQueryClient.SetTableQuery(initSchema)
}
