package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
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

	prompt := promptui.Prompt{HideEntered: true}

	for cfg.ProjectName == "" {
		prompt.Label = "Please enter a project name (e.g. my_app)"
		cfg.ProjectName, err = prompt.Run()
		if err != nil {
			return nil, err
		}
	}
	fmt.Println("Projec name:", cfg.ProjectName)

	for cfg.ProjectBase == "" {
		prompt.Label = "Please enter a project base (e.g. github.com/spf13/)"
		cfg.ProjectBase, err = prompt.Run()
		if err != nil {
			return nil, err
		}

		runes := []rune(cfg.ProjectBase)
		lastChar := runes[len(runes)-1]
		if lastChar != '/' {
			cfg.ProjectBase += "/"
		}
	}
	fmt.Println("Project base:", cfg.ProjectBase)

	return &cfg, nil
}
