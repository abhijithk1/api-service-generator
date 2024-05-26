/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	Name         string
	DBMS         string
	DBName       string
	PsqlUser     string
	PsqlPassword string
	TableName    string
	APIName      string
)

// generateTemplateCmd represents the generateTemplate command
var generateTemplateCmd = &cobra.Command{
	Use:   "generateTemplate",
	Short: "Generate Template",
	Long:  `Command that generates the API service template`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generateTemplate called.......")
		Name, _ = cmd.Flags().GetString("name")
		if Name == "" {
			fmt.Println("Name should be provided")
			fmt.Println("\nUsage: api-service-generator generateTemplate --name <name>")
			return
		}

		fmt.Printf("Name is : %s\n", Name)

		// Prompt the user for additional input
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("DBMS (currently supports only Postgres): ")
		DBMS, _ = reader.ReadString('\n')
		DBMS = strings.ToLower(strings.TrimSpace(DBMS))
		if DBMS == "" {
			fmt.Println("DBMS value is empty, by default using Postgres")
			DBMS = "postgres"
		}

		fmt.Print("Enter the POSTGRES_USER: ")
		PsqlUser, _ = reader.ReadString('\n')
		PsqlUser = strings.ToLower(strings.TrimSpace(PsqlUser))
		if PsqlUser == "" {
			fmt.Println("Database value is empty, by default using the value as 'postgres'")
			PsqlUser = "postgres"
		}

		fmt.Print("Enter the POSTGRES_PASSWORD: ")
		PsqlPassword, _ = reader.ReadString('\n')
		PsqlPassword = strings.ToLower(strings.TrimSpace(PsqlPassword))
		if PsqlPassword == "" {
			fmt.Println("Database value is empty, by default using value as 'password'")
			PsqlPassword = "password"
		}

		fmt.Print("Enter the Name of the Database: ")
		DBName, _ = reader.ReadString('\n')
		DBName = strings.ToLower(strings.TrimSpace(DBName))
		if DBName == "" {
			fmt.Println("Database value is empty, by default using the value of POSTGRES_USER")
			DBName = PsqlUser
		}

		fmt.Print("\nEnter a Table Name: ")
		TableName, _ = reader.ReadString('\n')
		TableName = strings.TrimSpace(TableName)
		if TableName == "" {
			fmt.Println("Table name is empty, by default using the name 'api_table'")
			TableName = "api_table"
		}

		fmt.Print("\nEnter a API Group: ")
		APIName, _ = reader.ReadString('\n')
		APIName = strings.TrimSpace(APIName)
		if APIName == "" {
			fmt.Println("API Group is empty, by default using the name 'dummy'")
			APIName = "dummy"
		}
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
