package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	ProjectName string `mapstructure:"projectname"`
	ProjectBase string `mapstructure:"projectbase"`
}

var cfg Config

var rootCmd = &cobra.Command{
	Use:   "tfg [project name]",
	Short: "Application for creating a REST API from a Postgresql Database",
	Long: `This applications is a tool created as a final degree project
with the aim of creating a new project with the use of a CLI.`,
	Run: func(cmd *cobra.Command, args []string) {
		// arg := args[0]
		// err := helloworld.CreateHelloWorld(arg)

		// if err != nil {
		// 	fmt.Println("Error creating project:", err)
		// } else {
		// 	fmt.Println("Hello world app created successfully!")
		// }

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
	rootCmd.PersistentFlags().StringP("projectname", "n", "", "project name")
	rootCmd.PersistentFlags().StringP("projectbase", "b", "", "base project directory eg. github.com/spf13/")
	viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
	viper.BindPFlag("projectname", rootCmd.PersistentFlags().Lookup("projectname"))
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
