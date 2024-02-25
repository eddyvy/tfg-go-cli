package cmd

import (
	"fmt"
	"os"

	helloworld "github.com/eddyvy/tfg-go-cli/hello-world"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hellogo [project name]",
	Short: "Application for creating a hello world",
	Long: `This applications is a tool created as a final degree project
with the aim of creating a new project with the use of a CLI.
This application, at this moment, only creates a hello world project.`,
	Run:  run,
	Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hellogo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func run(cmd *cobra.Command, args []string) {
	arg := args[0]
	err := helloworld.CreateHelloWorld(arg)

	if err != nil {
		fmt.Println("Error creating project:", err)
	} else {
		fmt.Println("Hello world app created successfully!")
	}
}
