package cmd

import (
	"testing"

	"github.com/eddyvy/tfg-go-cli/cmd/mocks"
	"github.com/eddyvy/tfg-go-cli/internal"
)

func TestRunNew(t *testing.T) {
	fakeConfig := &internal.GlobalConfig{
		ConfigFile: "config.yaml",
		Version:    "1.0.0",
		ProjectConfig: &internal.ProjectConfig{
			Name:       "testproject",
			Base:       "github.com/eddyvy",
			ProjectDir: "/home/user/projects",
		},
		DatabaseConfig: &internal.DatabaseConfig{
			Host:           "localhost",
			Port:           "5432",
			User:           "postgres",
			Type:           "postgresql",
			Database:       "testdb",
			Schema:         "public",
			Password:       "password",
			SSL:            "disable",
			Tables:         nil,
			UpdatingTables: nil,
		},
	}

	t.Run("New command works correctly", func(t *testing.T) {
		cliN := mocks.NewCliNewI(t)
		runFunc := runNew(cliN)

		cliN.On("ReadFlagsConfig", "").Return(fakeConfig, nil)
		cliN.On("ConfigureDatabase", fakeConfig).Return(nil)
		cliN.On("CreateNewProject", fakeConfig).Return(nil)
		cliN.On("FormatProject", fakeConfig).Return(nil)
		cliN.On("TidyProject", fakeConfig).Return(nil)
		cliN.On("GoModDownloadProject", fakeConfig).Return(nil)

		runFunc(nil, []string{})
	})

	t.Run("New command works correctly with projectName", func(t *testing.T) {
		cliN := mocks.NewCliNewI(t)
		runFunc := runNew(cliN)

		cliN.On("ReadFlagsConfig", "testproject").Return(fakeConfig, nil)
		cliN.On("ConfigureDatabase", fakeConfig).Return(nil)
		cliN.On("CreateNewProject", fakeConfig).Return(nil)
		cliN.On("FormatProject", fakeConfig).Return(nil)
		cliN.On("TidyProject", fakeConfig).Return(nil)
		cliN.On("GoModDownloadProject", fakeConfig).Return(nil)

		runFunc(nil, []string{"testproject"})
	})
}
