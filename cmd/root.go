/*
Copyright © 2024 ABHIJITH K abhijith0807@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "api-service-generator",
	Short: "API Service Generator",
	Long: `A Golang CLI application that generates a Golang API service application template.
			It Uses Gin framework for the REST APIs
			SQLC for Database connection to PostgresQL

		This application generates a boiler template of an API service, where the user can build their own custom API service on top of it.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// alog.Config(alog.INFO, alog.ChannelMap{
	// 	"generator":    alog.INFO,
	// 	// "steps":   alog.INFO,
	// 	// "monitor": alog.DEBUG,
	// })
	// cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.api-service-generator.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
