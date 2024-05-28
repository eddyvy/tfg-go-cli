package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

func createProject(cfg *GlobalConfig, tables []string) error {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		os.Exit(1)
	}
	projectDir := filepath.Join(currentDir, cfg.ProjectConfig.Name)
	// templatesDir := filepath.Join(currentDir, "templates")

	projectExists, err := tfgExists(projectDir)
	if err != nil {
		return err
	}

	if projectExists {
		// TODO: Implement project update
		return nil
	} else {
		err := os.MkdirAll(projectDir, 0755)
		if err != nil {
			return err
		}
		return nil
	}
}
