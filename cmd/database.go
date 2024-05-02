package cmd

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Type     string `mapstructure:"db_type"`
	Host     string `mapstructure:"db_host"`
	Port     string `mapstructure:"db_port"`
	User     string `mapstructure:"db_user"`
	Password string `mapstructure:"db_pass"`
}

func (d *DatabaseConfig) ConnectionString() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d", d.Type, d.User, d.Password, d.Host, d.Port)
}

func readDatabaseConfig() (*DatabaseConfig, error) {
	var cfg DatabaseConfig

	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	for cfg.Type == "" {
		fmt.Println("Please enter a database type (e.g. postgresql):")
		fmt.Scanln(&cfg.Type)
	}

	for cfg.Host == "" {
		fmt.Println("Please enter a database host (e.g. localhost):")
		fmt.Scanln(&cfg.Host)
	}

	for cfg.Port == "" {
		fmt.Println("Please enter a database port (e.g. 5432):")
		fmt.Scanln(&cfg.Port)
	}

	for cfg.User == "" {
		fmt.Println("Please enter a database user:")
		fmt.Scanln(&cfg.User)
	}

	for cfg.Password == "" {
		fmt.Println("Please enter a database password:")
		fmt.Scanln(&cfg.Password)
	}

	return &cfg, nil
}

func connectDatabase(cfg *DatabaseConfig) (*sql.DB, error) {
	if cfg.Type != "postgresql" {
		fmt.Println("Sorry, only \"postgresql\" database type is supported")
		return nil, fmt.Errorf("Sorry, database type %s is not supported", cfg.Type)
	}

	fmt.Println("Connecting to database...")

	connStr := cfg.ConnectionString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to the database")

	return db, nil
}
