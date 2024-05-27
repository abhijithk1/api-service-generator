package docker

import (
	"fmt"
	"strings"

	"github.com/abhijithk1/api-service-generator/common"
	"github.com/abhijithk1/api-service-generator/models"
)

var postgresRun = `docker run --name postgres_db -p 5432:5432 -e POSTGRES_USER=%s -e POSTGRES_PASSWORD=%s POSTGRES_DB=%s -d postgres`
func RunPostgresContainer(dbInputs models.DBInputs) error{
	postgresRun = fmt.Sprintf(postgresRun, dbInputs.PsqlUser, dbInputs.PsqlPassword, dbInputs.DBName)
	dockerCMds := strings.Split(postgresRun, " ")

	_, err := common.ExecuteCmds(dockerCMds[0], dockerCMds[1:])
	if err != nil {
		return err
	}

	return nil
}