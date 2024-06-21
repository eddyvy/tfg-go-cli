package cmd

import (
	"fmt"
	"os"

	"github.com/eddyvy/tfg-go-cli/internal"
	"github.com/spf13/cobra"
)

type CliAddI interface {
	ReadYamlConfig(projectRelPath string) (*internal.GlobalConfig, error)
	ConfigureDatabaseForUpdate(conf *internal.GlobalConfig) error
	UpdateProject(conf *internal.GlobalConfig) error
	FormatProject(conf *internal.GlobalConfig) error
	TidyProject(conf *internal.GlobalConfig) error
}

type CliAdd struct{}

func (c *CliAdd) ReadYamlConfig(projectRelPath string) (*internal.GlobalConfig, error) {
	return internal.ReadYamlConfig(projectRelPath)
}

func (c *CliAdd) ConfigureDatabaseForUpdate(conf *internal.GlobalConfig) error {
	return internal.ConfigureDatabaseForUpdate(conf)
}

func (c *CliAdd) UpdateProject(conf *internal.GlobalConfig) error {
	return internal.UpdateProject(conf)
}

func (c *CliAdd) FormatProject(conf *internal.GlobalConfig) error {
	return internal.FormatProject(conf)
}

func (c *CliAdd) TidyProject(conf *internal.GlobalConfig) error {
	return internal.TidyProject(conf)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds to an existinbg TFG project a REST API endpoint from a Postgresql Database Table",
	Run:   runAdd(&CliAdd{}),
	Args:  cobra.RangeArgs(0, 1),
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func runAdd(cliAdd CliAddI) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		projectRelPath := ""
		if len(args) == 1 {
			projectRelPath = args[0]
		}

		conf, err := cliAdd.ReadYamlConfig(projectRelPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = cliAdd.ConfigureDatabaseForUpdate(conf)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Updating project...")

		err = cliAdd.UpdateProject(conf)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = cliAdd.FormatProject(conf)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = cliAdd.TidyProject(conf)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Project updated successfully!")
	}
}
