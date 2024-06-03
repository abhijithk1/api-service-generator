package cleanup

import (
	"fmt"

	"github.com/abhijithk1/api-service-generator/common"
)

func CleanUp(wrkDir, containerName string) {
	err := removeDirectory(wrkDir)
	if err != nil {
		fmt.Println("Error : ", err)
	}
	if containerName != "" {
		err = removeDockerContainer(containerName)
		if err != nil {
			fmt.Println("Error : ", err)
		}
	}
}

func removeDirectory(wrkDir string) error {
	cmdStr := "rm"
	cmdArgs := []string{"-rf", wrkDir}
	output, err := common.ExecuteCmds(cmdStr, cmdArgs, ".")
	if err != nil {
		fmt.Printf("\nError running command: %s\nOutput: %s\n", err, output)
		return err
	}

	fmt.Printf("\n\nSuccessfully removed the directory: %s\n", wrkDir)
	return nil
}

func removeDockerContainer(containerName string) error{
	cmdStr := "docker"
	cmdArgs1 := []string{"kill", containerName}
	cmdArgs2 := []string{"rm", containerName}

	output, err := common.ExecuteCmds(cmdStr, cmdArgs1, ".")
	if err != nil {
		fmt.Printf("\nError running command: %s\nOutput: %s\n", err, output)
		return err
	}

	fmt.Printf("\n\nSuccessfully stopped the container: %s\n", containerName)

	output, err = common.ExecuteCmds(cmdStr, cmdArgs2, ".")
	if err != nil {
		fmt.Printf("\nError running command: %s\nOutput: %s\n", err, output)
		return err
	}

	fmt.Printf("\n\nSuccessfully removed the container: %s\n", containerName)
	
	return nil
}