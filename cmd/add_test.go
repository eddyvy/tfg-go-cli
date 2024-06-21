package cmd

import (
	"testing"

	"github.com/eddyvy/tfg-go-cli/cmd/mocks"
	"github.com/eddyvy/tfg-go-cli/internal"
)

func TestRunAdd(t *testing.T) {
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

	t.Run("Add command works correctly", func(t *testing.T) {
		cliA := mocks.NewCliAddI(t)
		runFunc := runAdd(cliA)

		cliA.On("ReadYamlConfig", "").Return(fakeConfig, nil)
		cliA.On("ConfigureDatabaseForUpdate", fakeConfig).Return(nil)
		cliA.On("UpdateProject", fakeConfig).Return(nil)
		cliA.On("FormatProject", fakeConfig).Return(nil)
		cliA.On("TidyProject", fakeConfig).Return(nil)

		runFunc(nil, []string{})
	})

	t.Run("Add command works correctly with projectName", func(t *testing.T) {
		cliA := mocks.NewCliAddI(t)
		runFunc := runAdd(cliA)

		cliA.On("ReadYamlConfig", "testproject").Return(fakeConfig, nil)
		cliA.On("ConfigureDatabaseForUpdate", fakeConfig).Return(nil)
		cliA.On("UpdateProject", fakeConfig).Return(nil)
		cliA.On("FormatProject", fakeConfig).Return(nil)
		cliA.On("TidyProject", fakeConfig).Return(nil)

		runFunc(nil, []string{"testproject"})
	})
}
