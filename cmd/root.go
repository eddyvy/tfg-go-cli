package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "Application for creating a REST API from a Postgresql Database",
	Long: `This applications is a tool created as a final degree project
with the aim of creating a new project with the use of a CLI.`,
	Args: cobra.ExactArgs(1),
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
