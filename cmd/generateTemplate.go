/*
Copyright © 2024 ABHIJITH K abhijith0807@gmail.com
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/abhijithk1/api-service-generator/api"
	"github.com/abhijithk1/api-service-generator/api/mw"
	"github.com/abhijithk1/api-service-generator/cleanup"
	"github.com/abhijithk1/api-service-generator/common"
	finalsetup "github.com/abhijithk1/api-service-generator/common/finalSetup"
	"github.com/abhijithk1/api-service-generator/db"
	"github.com/abhijithk1/api-service-generator/models"
	"github.com/abhijithk1/api-service-generator/util"
	"github.com/spf13/cobra"
)

// generateTemplateCmd represents the generateTemplate command
var generateTemplateCmd = &cobra.Command{
	Use:   "go-template",
	Short: "Generate Template",
	Long:  `Command that generates the API service template`,
	Run: runGenerateTemplate,
}

func init() {
	generateTemplateCmd.Flags().StringP("name", "n", "", "Name of the API Service that needs to be generated.")
	rootCmd.AddCommand(generateTemplateCmd)
}

func runGenerateTemplate(cmd *cobra.Command, args []string) {
	dbInputs := models.DBInputs{}
	apiInputs := models.APIInputs{}

	dbInputs.WrkDir, _ = cmd.Flags().GetString("name")
	if dbInputs.WrkDir == "" {
		fmt.Println("Name should be provided")
		fmt.Println("\nUsage: api-service-generator generateTemplate --name <name>")
		return
	}
	apiInputs.WrkDir = dbInputs.WrkDir

	reader := bufio.NewReader(os.Stdin)
	dbInputs.DBMS = promptForInput(reader, "Enter the Database Driver (currently supports only Postgres): ", "postgres", common.IsValidString)
	dbInputs.ContainerName = promptForInput(reader, "Enter the name for the Postgres Docker container: ", "postgres_db", common.IsValidString)
	dbInputs.ContainerPort = promptForInt(reader, "Enter the name for the Postgres Docker container port: ", 6432)
	dbInputs.PsqlUser = promptForInput(reader, "Enter the POSTGRES_USER: ", "postgres", common.IsValidString)
	dbInputs.PsqlPassword = promptForInput(reader, "Enter the POSTGRES_PASSWORD: ", "password", common.IsValidString)
	dbInputs.DBName = promptForInput(reader, "Enter the Name of the Database: ", dbInputs.PsqlUser, common.IsValidString)
	dbInputs.TableName = promptForInput(reader, "Enter a Table Name: ", "api_table", common.IsValidString)
	apiInputs.TableName = dbInputs.TableName
	apiInputs.APIGroup = promptForInput(reader, "Enter an API Group: ", "dummy", common.IsValidString)
	apiInputs.GoModule = promptForInput(reader, "Enter a Go Module Base Path: ", "example", func(s string) bool {return true})
	dbInputs.GoModule = apiInputs.GoModule

	steps := []func() error{
		func() error { return common.Initialise(apiInputs.GoModule, dbInputs.WrkDir) },
		func() error { return db.Setup(dbInputs) },
		func() error { return api.Setup(apiInputs) },
		func() error { return mw.SetupMiddleWare(apiInputs) },
		func() error { return util.SetUtils(apiInputs.WrkDir) },
		func() error { return finalsetup.FinalSetup(apiInputs, dbInputs) },
	}

	for _, step := range steps {
		if err := step(); err != nil {
			fmt.Printf("Error: Setup step failed: %v", err)
			cleanup.CleanUp(dbInputs.WrkDir, dbInputs.ContainerName)
			return
		}
	}

	fmt.Println("\n\n Successfully generated API service Template...")
	fmt.Println("\n\n Happy Coding...")
}

func promptForInput(reader *bufio.Reader, prompt, defaultValue string, validationFunc func(string) bool) string {
	for {
		fmt.Print("\n" + prompt)
		input, _ := reader.ReadString('\n')
		input = strings.ToLower(strings.TrimSpace(input))
		if input == "" {
			fmt.Printf("%s value is empty, by default using %s\n", prompt, defaultValue)
			return defaultValue
		}
		if validationFunc(input) {
			return input
		}
		fmt.Printf("Invalid input! %s contains invalid characters. Please try again.\n", prompt)
	}
}

func promptForInt(reader *bufio.Reader, prompt string, defaultValue int) int {
	fmt.Print("\n" + prompt)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		fmt.Printf("%s value is empty, by default using %d\n", prompt, defaultValue)
		return defaultValue
	}
	value, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("%s value is invalid, by default using %d\n", prompt, defaultValue)
		return defaultValue
	}
	return value
}
