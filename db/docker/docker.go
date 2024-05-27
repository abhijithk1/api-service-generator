package docker

import (
	"strings"

	"github.com/abhijithk1/api-service-generator/common"
)

var postgresRun = `docker run --name postgres_db -p 5432:5432 -e POSTGRES_USER=%s -e POSTGRES_PASSWORD=%s POSTGRES_DB=%s -d postgres`
func RunPostgresContainer() error{
	dockerCMds := strings.Split(postgresRun, " ")

	_, err := common.ExecuteCmds(dockerCMds[0], dockerCMds[1:])
	if err != nil {
		return err
	}

	return nil
}