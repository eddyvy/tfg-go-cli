package internal

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
)

const TFG_VERSION = "0.0.1"
const TFG_FILENAME = "tfg.yml"

type GlobalConfig struct {
	ConfigFile     string          `yaml:"-"`
	Version        string          `yaml:"version"`
	ProjectConfig  *ProjectConfig  `yaml:"project"`
	DatabaseConfig *DatabaseConfig `yaml:"database"`
}

type ProjectConfig struct {
	Name string `mapstructure:"project_name"`
	Base string `mapstructure:"project_base"`
}

type DatabaseConfig struct {
	Type     string             `mapstructure:"db_type" yaml:"type"`
	Host     string             `mapstructure:"db_host" yaml:"host"`
	Port     string             `mapstructure:"db_port" yaml:"port"`
	Database string             `mapstructure:"db_database" yaml:"database"`
	Schema   string             `mapstructure:"db_schema" yaml:"schema"`
	User     string             `mapstructure:"db_user" yaml:"-"`
	Password string             `mapstructure:"db_pass" yaml:"-"`
	SSL      string             `mapstructure:"db_ssl" yaml:"ssl"`
	Tables   []*TableDefinition `yaml:"tables"`
}

func (d *DatabaseConfig) ConnectionString() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s&search_path=%s", d.Type, d.User, d.Password, d.Host, d.Port, d.Database, d.SSL, d.Schema)
}

func ReadFlagsConfig(projectName string) (*GlobalConfig, error) {
	var proCfg ProjectConfig

	err := viper.Unmarshal(&proCfg)
	if err != nil {
		return nil, err
	}

	prompt := promptui.Prompt{HideEntered: true}

	if projectName != "" {
		proCfg.Name = projectName
	}

	for proCfg.Name == "" {
		prompt.Label = "Please enter a project name (e.g. my_app)"
		proCfg.Name, err = prompt.Run()
		if err != nil {
			return nil, err
		}
	}
	fmt.Println("Project name:", proCfg.Name)

	for proCfg.Base == "" {
		prompt.Label = "Please enter a project base (e.g. github.com/spf13/)"
		proCfg.Base, err = prompt.Run()
		if err != nil {
			return nil, err
		}

		runes := []rune(proCfg.Base)
		lastChar := runes[len(runes)-1]
		if lastChar != '/' {
			proCfg.Base += "/"
		}
	}
	fmt.Println("Project base:", proCfg.Base)

	var dbCfg DatabaseConfig

	err = viper.Unmarshal(&dbCfg)
	if err != nil {
		return nil, err
	}

	for dbCfg.Type == "" {
		prompt.Label = "Please enter a database type (e.g. postgresql)"
		dbCfg.Type, err = prompt.Run()
		if err != nil {
			return nil, err
		}
	}
	fmt.Println("Database type:", dbCfg.Type)

	for dbCfg.Host == "" {
		prompt.Label = "Please enter a database host (e.g. localhost)"
		dbCfg.Host, err = prompt.Run()
		if err != nil {
			return nil, err
		}
	}
	fmt.Println("Host:", dbCfg.Host)

	for dbCfg.Port == "" {
		prompt.Label = "Please enter a database port (e.g. 5432)"
		dbCfg.Port, err = prompt.Run()
		if err != nil {
			return nil, err
		}
	}
	fmt.Println("Port:", dbCfg.Port)

	for dbCfg.Database == "" {
		prompt.Label = "Please enter the database name (e.g. postgres)"
		dbCfg.Database, err = prompt.Run()
		if err != nil {
			return nil, err
		}
	}
	fmt.Println("Database:", dbCfg.Database)

	for dbCfg.Schema == "" {
		prompt.Label = "Please enter a database schema (e.g. public)"
		dbCfg.Schema, err = prompt.Run()
		if err != nil {
			return nil, err
		}
	}
	fmt.Println("Schema:", dbCfg.Schema)

	for dbCfg.User == "" {
		prompt.Label = "Please enter a database user"
		dbCfg.User, err = prompt.Run()
		if err != nil {
			return nil, err
		}
	}
	fmt.Println("User:", dbCfg.User)

	for dbCfg.Password == "" {
		prompt.Label = "Please enter a database password"
		prompt.Mask = '*'
		dbCfg.Password, err = prompt.Run()
		if err != nil {
			return nil, err
		}
	}

	cfg := &GlobalConfig{
		ConfigFile:     TFG_FILENAME,
		Version:        TFG_VERSION,
		ProjectConfig:  &proCfg,
		DatabaseConfig: &dbCfg,
	}

	return cfg, nil
}
