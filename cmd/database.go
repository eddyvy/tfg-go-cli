package cmd

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Type     string `mapstructure:"db_type"`
	Host     string `mapstructure:"db_host"`
	Port     string `mapstructure:"db_port"`
	Database string `mapstructure:"db_database"`
	User     string `mapstructure:"db_user"`
	Password string `mapstructure:"db_pass"`
	SSL      string `mapstructure:"db_ssl"`
}

func (d *DatabaseConfig) ConnectionString() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s", d.Type, d.User, d.Password, d.Host, d.Port, d.Database, d.SSL)
}

func initDatabaseFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("db_type", "", "", "database type (e.g. postgresql) ONLY postgresql is supported")
	cmd.PersistentFlags().StringP("db_host", "", "", "database host (e.g. 127.0.0.1)")
	cmd.PersistentFlags().StringP("db_port", "", "", "database port (e.g. 5432)")
	cmd.PersistentFlags().StringP("db_database", "", "", "database name")
	cmd.PersistentFlags().StringP("db_user", "", "", "database user")
	cmd.PersistentFlags().StringP("db_pass", "", "", "database password")
	cmd.PersistentFlags().StringP("db_ssl", "", "", "ssl mode enabled")
	viper.SetDefault("db_type", "postgresql")
	viper.SetDefault("db_ssl", "disable")
	viper.BindPFlag("db_type", cmd.PersistentFlags().Lookup("db_type"))
	viper.BindPFlag("db_host", cmd.PersistentFlags().Lookup("db_host"))
	viper.BindPFlag("db_port", cmd.PersistentFlags().Lookup("db_port"))
	viper.BindPFlag("db_database", cmd.PersistentFlags().Lookup("db_database"))
	viper.BindPFlag("db_user", cmd.PersistentFlags().Lookup("db_user"))
	viper.BindPFlag("db_pass", cmd.PersistentFlags().Lookup("db_pass"))
	viper.BindPFlag("db_ssl", cmd.PersistentFlags().Lookup("db_ssl"))
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

	for cfg.Database == "" {
		fmt.Println("Please enter a database name (e.g. postgres):")
		fmt.Scanln(&cfg.Database)
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
		return nil, fmt.Errorf("sorry, database type %s is not supported", cfg.Type)
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

func readTables(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		if err != nil {
			return nil, err
		}
		tables = append(tables, tableName)
	}

	if tables == nil {
		return nil, fmt.Errorf("no tables found in the database")
	}

	return tables, nil
}
