package internal

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"golang.org/x/tools/imports"
)

func FormatProject(cfg *GlobalConfig) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	projectDir := filepath.Join(currentDir, cfg.ProjectConfig.Name)

	err = filepath.Walk(projectDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".go" {
			return nil
		}

		return formatFile(path)
	})

	if err != nil {
		return err
	}

	return nil
}

func TidyProject(cfg *GlobalConfig) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	projectDir := filepath.Join(currentDir, cfg.ProjectConfig.Name)

	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = projectDir
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func formatFile(path string) error {
	src, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	out, err := imports.Process(path, src, nil)
	if err != nil {
		return err
	}

	return os.WriteFile(path, out, os.ModePerm)
}
