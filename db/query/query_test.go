package query

import (
	"errors"
	"testing"

	"github.com/abhijithk1/api-service-generator/common"
	"github.com/abhijithk1/api-service-generator/mocks"
	"github.com/abhijithk1/api-service-generator/models"
	"github.com/stretchr/testify/assert"
)

func TestSetTableQuery(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	fileName := "test/pkg/db/query/dummy.sql"
	initSchema := models.InitSchema{
		TableName: "dummy",
		WrkDir:    "test",
	}

	query_sql := `
	-- name: Listdummy :many
		SELECT * FROM dummy;
	`

	table_sql = query_sql

	mockCmdsExecutor.On("CreateFileAndItsContent", fileName, initSchema, query_sql).Return(nil)

	err := SetTableQuery(initSchema)
	assert.NoError(t, err)

	mockCmdsExecutor.AssertExpectations(t)
}

func TestSetTableQuery_Error(t *testing.T) {
	mockCmdsExecutor := mocks.NewMockCmdsExecutor()
	common.DefaultExecutor = mockCmdsExecutor

	fileName := "test/pkg/db/query/dummy.sql"
	initSchema := models.InitSchema{
		TableName: "dummy",
		WrkDir:    "test",
	}
	query_sql := `
	-- name: Listdummy :many
		SELECT * FROM dummy;
	`

	table_sql = query_sql

	mockCmdsExecutor.On("CreateFileAndItsContent", fileName, initSchema, query_sql).Return(errors.New("error in creating query file and content"))

	err := SetTableQuery(initSchema)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error in creating query file and content")

	mockCmdsExecutor.AssertExpectations(t)
}
