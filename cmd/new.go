package cmd

import (
	"fmt"
	"os"

	"github.com/eddyvy/tfg-go-cli/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type CliNewI interface {
	ReadFlagsConfig(projectName string) (*internal.GlobalConfig, error)
	ConfigureDatabase(conf *internal.GlobalConfig) error
	CreateNewProject(conf *internal.GlobalConfig) error
	FormatProject(conf *internal.GlobalConfig) error
	TidyProject(conf *internal.GlobalConfig) error
	GoModDownloadProject(conf *internal.GlobalConfig) error
}

type CliNew struct{}

func (c *CliNew) ReadFlagsConfig(projectName string) (*internal.GlobalConfig, error) {
	return internal.ReadFlagsConfig(projectName)
}

func (c *CliNew) ConfigureDatabase(conf *internal.GlobalConfig) error {
	return internal.ConfigureDatabase(conf)
}

func (c *CliNew) CreateNewProject(conf *internal.GlobalConfig) error {
	return internal.CreateNewProject(conf)
}

func (c *CliNew) FormatProject(conf *internal.GlobalConfig) error {
	return internal.FormatProject(conf)
}

func (c *CliNew) TidyProject(conf *internal.GlobalConfig) error {
	return internal.TidyProject(conf)
}

func (c *CliNew) GoModDownloadProject(conf *internal.GlobalConfig) error {
	return internal.GoModDownloadProject(conf)
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new project with a REST API from a Postgresql Database",
	Run:   runNew(&CliNew{}),
	Args:  cobra.RangeArgs(0, 1),
}

func init() {
	rootCmd.AddCommand(newCmd)
	initNewFlags(newCmd)
}

func initNewFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("project_name", "n", "", "project name")
	cmd.PersistentFlags().StringP("project_base", "", "", "project base")
	viper.BindPFlag("project_name", cmd.PersistentFlags().Lookup("project_name"))
	viper.BindPFlag("project_base", cmd.PersistentFlags().Lookup("project_base"))

	cmd.PersistentFlags().StringP("db_type", "", "", "database type (e.g. postgresql) ONLY postgresql is supported")
	cmd.PersistentFlags().StringP("db_host", "", "", "database host (e.g. 127.0.0.1)")
	cmd.PersistentFlags().StringP("db_port", "", "", "database port (e.g. 5432)")
	cmd.PersistentFlags().StringP("db_database", "", "", "database name")
	cmd.PersistentFlags().StringP("db_schema", "", "", "schema name")
	cmd.PersistentFlags().StringP("db_user", "", "", "database user")
	cmd.PersistentFlags().StringP("db_pass", "", "", "database password")
	cmd.PersistentFlags().StringP("db_ssl", "", "", "ssl mode enabled")
	viper.SetDefault("db_type", "postgresql")
	viper.BindPFlag("db_type", cmd.PersistentFlags().Lookup("db_type"))
	viper.BindPFlag("db_host", cmd.PersistentFlags().Lookup("db_host"))
	viper.BindPFlag("db_port", cmd.PersistentFlags().Lookup("db_port"))
	viper.BindPFlag("db_database", cmd.PersistentFlags().Lookup("db_database"))
	viper.BindPFlag("db_schema", cmd.PersistentFlags().Lookup("db_schema"))
	viper.BindPFlag("db_user", cmd.PersistentFlags().Lookup("db_user"))
	viper.BindPFlag("db_pass", cmd.PersistentFlags().Lookup("db_pass"))
	viper.BindPFlag("db_ssl", cmd.PersistentFlags().Lookup("db_ssl"))
}

func runNew(cliN CliNewI) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		projectName := ""
		if len(args) == 1 {
			projectName = args[0]
		}

		conf, err := cliN.ReadFlagsConfig(projectName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = cliN.ConfigureDatabase(conf)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Creating project...")

		err = cliN.CreateNewProject(conf)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = cliN.FormatProject(conf)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = cliN.TidyProject(conf)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = cliN.GoModDownloadProject(conf)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Project created successfully!")
		fmt.Println("Run \"cd " + conf.ProjectConfig.ProjectDir + "\" to navigate to the project folder.")
		fmt.Println("Run \"go run .\" to start your new server! 🚀")
	}
}
