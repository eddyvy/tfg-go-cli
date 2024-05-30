package internal

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-yaml/yaml"
)

func CreateNewProject(cfg *GlobalConfig) error {
	err := createProjectFolder(cfg.ProjectConfig.Name)
	if err != nil {
		return err
	}

	err = createTfgYaml(cfg)
	if err != nil {
		RemoveAll(cfg)
		return err
	}

	err = ExecuteTemplatesBase(cfg)
	if err != nil {
		RemoveAll(cfg)
		return err
	}

	err = ExecuteTemplatesResources(cfg)
	if err != nil {
		RemoveAll(cfg)
		return err
	}

	err = UpdateRouter(cfg.DatabaseConfig.Tables, cfg)
	if err != nil {
		RemoveAll(cfg)
		return err
	}

	return nil
}

func RemoveAll(cfg *GlobalConfig) error {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return err
	}

	projectDir := filepath.Join(currentDir, cfg.ProjectConfig.Name)

	fmt.Println("Coming back changes...")
	err = os.RemoveAll(projectDir)
	if err != nil {
		fmt.Println("Error removing directory:", err)
		return err
	}

	return nil
}

func createProjectFolder(projectName string) error {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		os.Exit(1)
	}
	projectDir := filepath.Join(currentDir, projectName)

	_, err = os.Stat(projectDir)

	if err == nil {
		return fmt.Errorf("a folder with the same name already exists")
	} else if !os.IsNotExist(err) {
		return err
	}

	err = os.MkdirAll(projectDir, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func createTfgYaml(cfg *GlobalConfig) error {
	tfgYmlBytes, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	tfgYmlPath := filepath.Join(cfg.ProjectConfig.Name, cfg.ConfigFile)
	return os.WriteFile(tfgYmlPath, tfgYmlBytes, os.ModePerm)
}
