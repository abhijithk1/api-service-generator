package docker

import (
	"fmt"
	"strings"

	"github.com/abhijithk1/api-service-generator/common"
	"github.com/abhijithk1/api-service-generator/models"
)

var (
	PostgresRun = `docker run --name %s -p %d:5432 -e POSTGRES_USER=%s -e POSTGRES_PASSWORD=%s -e POSTGRES_DB=%s -v pgdata:/var/lib/postgresql/data -d postgres`
	MySqlRun = `docker run --name %s -p %d:3306 -e MYSQL_ROOT_PASSWORD=%s -e MYSQL_USER=%s -e MYSQL_PASSWORD=%s -e MYSQL_DATABASE=%s -v mysql_data:/var/lib/mysql -d mysql`
)

type DockerContainer interface {
	RunContainer(dbInputs models.DBInputs) error
}

var DefaultDockerClient DockerContainer = &DockerClient{}

type DockerClient struct{}

func (d *DockerClient)RunContainer(dbInputs models.DBInputs) error {
	var runCmd string
	switch dbInputs.DBMS {
	case "postgres":
		runCmd = runPostgresContainer(dbInputs)
	case "mysql":
		runCmd = runMySqlContainer(dbInputs)
	default:
		fmt.Println("Driver not supported")
		return fmt.Errorf("driver not supported")
	}

	dockerCMds := strings.Split(runCmd, " ")
	fmt.Println("\n\nRunning command: ", runCmd)
	
	output, err := common.ExecuteCmds(dockerCMds[0], dockerCMds[1:], ".")
	if err != nil {
		if strings.Contains(string(output), `already in use by container`) {
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

func runPostgresContainer(dbInputs models.DBInputs) string {
	return fmt.Sprintf(PostgresRun, dbInputs.ContainerName, dbInputs.ContainerPort, dbInputs.Postgres.PsqlUser, dbInputs.Postgres.PsqlPassword, strings.ToLower(dbInputs.DBName))
}

func runMySqlContainer(dbInputs models.DBInputs) string {
	return fmt.Sprintf(MySqlRun, dbInputs.ContainerName, dbInputs.ContainerPort, dbInputs.MySQL.MysqlRootPassword, dbInputs.MySQL.MysqlUser, dbInputs.MySQL.MysqlPassword, strings.ToLower(dbInputs.DBName))
}
