package api

import (
	"fmt"

	"github.com/abhijithk1/api-service-generator/common"
	"github.com/abhijithk1/api-service-generator/models"
)

var (
	APIFilePath = "%s/api/v1/%s/"
)

func Setup(apiInputs models.APIInputs) {
	apiInputs.APIGroupTitle = common.ToCamelCase(apiInputs.APIGroup)
	apiInputs.TableNameTitle = common.ToCamelCase(apiInputs.TableName)

	err := createApiGroup(apiInputs)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	err = createControllerFile(apiInputs)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	err = createServiceFile(apiInputs)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}

func createApiGroup(apiInputs models.APIInputs) error {
	filePath := fmt.Sprintf(APIFilePath, apiInputs.WrkDir,apiInputs.APIGroup)
	return common.CreateDirectory(filePath)
}

const controllerContent = `// Generated By API Service Generator
package {{.APIGroup}}

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type {{.APIGroupTitle}}Resource struct {
	service Service
}


func ResourceHandler(r *gin.RouterGroup, service Service) {
	resource := New{{.APIGroupTitle}}Resource(service)
	
	r.GET("/{{.APIGroup}}", resource.Get{{.APIGroupTitle}})
}

func New{{.APIGroupTitle}}Resource(service Service){{.APIGroupTitle}}Resource {
	return {{.APIGroupTitle}}Resource{service}
}

func (r *{{.APIGroupTitle}}Resource) Get{{.APIGroupTitle}}(c *gin.Context) {
	{{.TableNameTitle}}, err := r.service.Get{{.APIGroupTitle}}(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, {{.TableNameTitle}})
}

`

func createControllerFile(apiInputs models.APIInputs) error {
	fileName := fmt.Sprintf(APIFilePath, apiInputs.WrkDir,apiInputs.APIGroup) + "controller.go"
	return common.CreateFileAndItsContent(fileName, apiInputs, controllerContent)
}

const serviceContent = `// Generated By API Service Generator
package {{.APIGroup}}

import (
	"context"
	"example/{{.WrkDir}}/pkg/db"
)

type Service interface {
	Get{{.APIGroupTitle}}(ctx context.Context) ([]db.{{.TableNameTitle}}, error)
}

type {{.APIGroupTitle}}Service struct {
	DBConn *db.Queries
}

func New{{.APIGroupTitle}}Service(DBConn *db.Queries) {{.APIGroupTitle}}Service {
	return {{.APIGroupTitle}}Service{DBConn}
}

func (s *{{.APIGroupTitle}}Service) Get{{.APIGroupTitle}}(ctx context.Context) ([]db.{{.TableNameTitle}}, error) {
	return s.DBConn.List{{.TableName}}(ctx)
}

`

func createServiceFile(apiInputs models.APIInputs) error {
	fileName := fmt.Sprintf(APIFilePath, apiInputs.WrkDir,apiInputs.APIGroup) + "service.go"
	return common.CreateFileAndItsContent(fileName, apiInputs, serviceContent)
}