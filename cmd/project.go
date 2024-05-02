package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ProjectConfig struct {
	ProjectName string `mapstructure:"project_name"`
	ProjectBase string `mapstructure:"project_base"`
}

func initProjectFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("project_name", "n", "", "project name")
	cmd.PersistentFlags().StringP("project_base", "", "", "project base")
	viper.BindPFlag("project_name", cmd.PersistentFlags().Lookup("project_name"))
	viper.BindPFlag("project_base", cmd.PersistentFlags().Lookup("project_base"))
}

func readProjectConfig() (*ProjectConfig, error) {
	var cfg ProjectConfig

	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	for cfg.ProjectName == "" {
		fmt.Println("Please enter a project name (e.g. my_app):")
		fmt.Scanln(&cfg.ProjectName)
	}

	for cfg.ProjectBase == "" {
		fmt.Println("Please enter a project base (e.g. github.com/spf13/):")
		fmt.Scanln(&cfg.ProjectBase)
	}

	return &cfg, nil
}
