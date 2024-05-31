package docker

import (
	"fmt"
	"strings"

	"github.com/abhijithk1/api-service-generator/common"
	"github.com/abhijithk1/api-service-generator/models"
)

var PostgresRun = `docker run --name %s -p %d:5432 -e POSTGRES_USER=%s -e POSTGRES_PASSWORD=%s -e POSTGRES_DB=%s -d postgres`

type DockerContainer interface {
	RunContainer(dbInputs models.DBInputs) error
}

var DefaultDockerClient DockerContainer = &DockerClient{}

type DockerClient struct{}

func (d *DockerClient)RunContainer(dbInputs models.DBInputs) error {
	runCmd := fmt.Sprintf(PostgresRun, dbInputs.ContainerName, dbInputs.ContainerPort, dbInputs.PsqlUser, dbInputs.PsqlPassword, strings.ToLower(dbInputs.DBName))
	dockerCMds := strings.Split(runCmd, " ")

	fmt.Println("\n\nRunning command: ", runCmd)
	
	output, err := common.ExecuteCmds(dockerCMds[0], dockerCMds[1:], ".")
	if err != nil {
		if strings.Contains(string(output), `The container name "/postgres_db" is already in use by container`) {
			fmt.Printf("\nOutput : %s\n", output)
			return err
		}
		fmt.Printf("\nError running command: %s\nOutput: %s\n", err, output)
		return err
	}

	fmt.Println("\n\nSuccessfully started the PostgreSQL container")
	return nil
}


func RunContainer(dbInputs models.DBInputs) error {
	return DefaultDockerClient.RunContainer(dbInputs)
}