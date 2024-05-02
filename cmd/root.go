package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ProjectConfig struct {
	ProjectName string `mapstructure:"project_name"`
	ProjectBase string `mapstructure:"project_base"`
}

var rootCmd = &cobra.Command{
	Use:   "tfg [project name]",
	Short: "Application for creating a REST API from a Postgresql Database",
	Long: `This applications is a tool created as a final degree project
with the aim of creating a new project with the use of a CLI.`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.Unmarshal(&cfg)

		for cfg.ProjectName == "" {
			fmt.Println("Please enter a project name (e.g. my_app):")
			fmt.Scanln(&cfg.ProjectName)
		}

		for cfg.ProjectBase == "" {
			fmt.Println("Please enter a project base (e.g. github.com/spf13/):")
			fmt.Scanln(&cfg.ProjectBase)
		}

		fmt.Println(cfg)

		fmt.Println("Creating project...")
	},
	Args: cobra.MatchAll(cobra.ArbitraryArgs),
}

func init() {
	rootCmd.PersistentFlags().StringP("project_base", "pb", "", "base project directory eg. github.com/myaccount/")
	rootCmd.PersistentFlags().StringP("project_name", "pn", "", "project name")
	viper.BindPFlag("project_base", rootCmd.PersistentFlags().Lookup("project_base"))
	viper.BindPFlag("project_name", rootCmd.PersistentFlags().Lookup("project_name"))
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
