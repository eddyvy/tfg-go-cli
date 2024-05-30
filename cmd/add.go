package cmd

import (
	"fmt"
	"os"

	"github.com/eddyvy/tfg-go-cli/internal"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds to an existinbg TFG project a REST API endpoint from a Postgresql Database Table",
	Run: func(cmd *cobra.Command, args []string) {
		projectRelPath := ""
		if len(args) == 1 {
			projectRelPath = args[0]
		}

		conf, err := internal.ReadYamlConfig(projectRelPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = internal.ConfigureDatabaseForUpdate(conf)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Updating project...")

		err = internal.UpdateProject(conf)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = internal.FormatProject(conf)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = internal.TidyProject(conf)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Project updated successfully!")
	},
	Args: cobra.RangeArgs(0, 1),
}

func init() {
	rootCmd.AddCommand(addCmd)
}
