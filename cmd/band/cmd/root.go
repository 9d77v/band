/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// params
var (
	pkgDir  string
	appName string
	service string
	entity  string
	idType  string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "band",
	Short: "A code generater for micro services",
	Long:  `A code generater for micro services`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	initCommand.Flags().StringVarP(&pkgDir, "pkg", "p", "", "go package name")
	initCommand.Flags().StringVarP(&appName, "app", "a", "", "app name")
	rootCmd.AddCommand(initCommand)
	serviceCommand.Flags().StringVarP(&service, "service", "s", "", "service name")
	serviceCommand.Flags().StringVarP(&entity, "entity", "e", "", "entity name")
	serviceCommand.Flags().StringVarP(&idType, "id_type", "t", "uint", "id type")
	rootCmd.AddCommand(serviceCommand)
}
