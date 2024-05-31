/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/abhijithk1/api-service-generator/common"
	"github.com/abhijithk1/api-service-generator/db"
	"github.com/abhijithk1/api-service-generator/models"
	"github.com/spf13/cobra"
)

// generateTemplateCmd represents the generateTemplate command
var generateTemplateCmd = &cobra.Command{
	Use:   "generateTemplate",
	Short: "Generate Template",
	Long:  `Command that generates the API service template`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generateTemplate called.......")
		dbInputs := models.DBInputs{}
		apiInputs := models.APIInputs{}

		dbInputs.WrkDir, _ = cmd.Flags().GetString("name")
		if dbInputs.WrkDir == "" {
			fmt.Println("Name should be provided")
			fmt.Println("\nUsage: api-service-generator generateTemplate --name <name>")
			return
		}

		apiInputs.WrkDir = dbInputs.WrkDir
		fmt.Printf("Name is : %s\n", dbInputs.WrkDir)

		// Prompt the user for additional input
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("DBMS (currently supports only Postgres): ")
		dbInputs.DBMS, _ = reader.ReadString('\n')
		dbInputs.DBMS = strings.ToLower(strings.TrimSpace(dbInputs.DBMS))
		if dbInputs.DBMS == "" {
			fmt.Println("DBMS value is empty, by default using Postgres")
			dbInputs.DBMS = "postgres"
		}

		fmt.Print("Enter the name for the Postgres Docker container: ")
		dbInputs.ContainerName, _ = reader.ReadString('\n')
		dbInputs.ContainerName = strings.ToLower(strings.TrimSpace(dbInputs.ContainerName))
		if dbInputs.ContainerName == "" {
			fmt.Println("Container name value is empty, by default using postgres_db")
			dbInputs.ContainerName = "postgres_db"
		}

		fmt.Print("Enter the name for the Postgres Docker container port: ")
		port, _ := reader.ReadString('\n')
		dbInputs.ContainerPort, _ = strconv.Atoi(strings.TrimSpace(port))
		if dbInputs.ContainerPort == 0 {
			fmt.Println("Container port value is empty, by default using 6432")
			dbInputs.ContainerPort = 6432
		}

		fmt.Print("Enter the POSTGRES_USER: ")
		dbInputs.PsqlUser, _ = reader.ReadString('\n')
		dbInputs.PsqlUser = strings.ToLower(strings.TrimSpace(dbInputs.PsqlUser))
		if dbInputs.PsqlUser == "" {
			fmt.Println("Database value is empty, by default using the value as 'postgres'")
			dbInputs.PsqlUser = "postgres"
		}

		fmt.Print("Enter the POSTGRES_PASSWORD: ")
		dbInputs.PsqlPassword, _ = reader.ReadString('\n')
		dbInputs.PsqlPassword = strings.ToLower(strings.TrimSpace(dbInputs.PsqlPassword))
		if dbInputs.PsqlPassword == "" {
			fmt.Println("Database value is empty, by default using value as 'password'")
			dbInputs.PsqlPassword = "password"
		}

		fmt.Print("Enter the Name of the Database: ")
		dbInputs.DBName, _ = reader.ReadString('\n')
		dbInputs.DBName = strings.ToLower(strings.TrimSpace(dbInputs.DBName))
		if dbInputs.DBName == "" {
			fmt.Println("Database value is empty, by default using the value of POSTGRES_USER")
			dbInputs.DBName = dbInputs.PsqlUser
		}

		fmt.Print("Enter a Table Name: ")
		dbInputs.TableName, _ = reader.ReadString('\n')
		dbInputs.TableName = strings.TrimSpace(dbInputs.TableName)
		if dbInputs.TableName == "" {
			fmt.Println("Table name is empty, by default using the name 'api_table'")
			dbInputs.TableName = "api_table"
		}

		fmt.Print("Enter a API Group: ")
		apiInputs.APIGroup, _ = reader.ReadString('\n')
		apiInputs.APIGroup = strings.TrimSpace(apiInputs.APIGroup)
		if apiInputs.APIGroup == "" {
			fmt.Println("API Group is empty, by default using the name 'dummy'")
			apiInputs.APIGroup = "dummy"
		}

		common.Initialise("example", dbInputs.WrkDir)
		db.Setup(dbInputs)
	},
}

func init() {
	generateTemplateCmd.Flags().StringP("name", "n", "", "Name of the API Service that needs to be generated.")
	rootCmd.AddCommand(generateTemplateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateTemplateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateTemplateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
