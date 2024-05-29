package cmd

import (
	"fmt"
	"os"

	"github.com/eddyvy/tfg-go-cli/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new project with a REST API from a Postgresql Database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		conf, err := internal.ReadFlagsConfig(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = internal.ConfigureDatabase(conf)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// os.Stdout.Sync()
		fmt.Println("Creating project...")

		// fmt.Printf("%+v\n", conf)
		// fmt.Printf("%+v\n", conf.ProjectConfig)
		// fmt.Printf("%+v\n", conf.DatabaseConfig)
		// for _, table := range conf.DatabaseConfig.Tables {
		// 	fmt.Printf("%+v\n", table)
		// 	for _, col := range table.Columns {
		// 		fmt.Printf("%+v\n", col)
		// 	}
		// }

		err = internal.CreateNewProject(conf)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
	Args: cobra.RangeArgs(0, 1),
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
	viper.SetDefault("db_schema", "public")
	viper.SetDefault("db_ssl", "disable")
	viper.BindPFlag("db_type", cmd.PersistentFlags().Lookup("db_type"))
	viper.BindPFlag("db_host", cmd.PersistentFlags().Lookup("db_host"))
	viper.BindPFlag("db_port", cmd.PersistentFlags().Lookup("db_port"))
	viper.BindPFlag("db_database", cmd.PersistentFlags().Lookup("db_database"))
	viper.BindPFlag("db_schema", cmd.PersistentFlags().Lookup("db_schema"))
	viper.BindPFlag("db_user", cmd.PersistentFlags().Lookup("db_user"))
	viper.BindPFlag("db_pass", cmd.PersistentFlags().Lookup("db_pass"))
	viper.BindPFlag("db_ssl", cmd.PersistentFlags().Lookup("db_ssl"))
}
