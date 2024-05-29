package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-yaml/yaml"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DatabaseConfig struct {
	Type     string
	Host     string
	Port     string
	Database string
	Schema   string
	User     string
	Password string
	SSL      string
}

func (d *DatabaseConfig) ConnectionString() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s&search_path=%s",
		d.Type, d.User, d.Password, d.Host, d.Port, d.Database, d.SSL, d.Schema)
}

var DB *sql.DB

func InitDb() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.ConnectionString())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to the database")
	DB = db
}

func loadConfig() (*DatabaseConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile("tfg.yml")
	if err != nil {
		return nil, err
	}

	var cfg map[interface{}]interface{}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	dbCfg := cfg["database"].(map[interface{}]interface{})

	return &DatabaseConfig{
		Type:     dbCfg["type"].(string),
		Host:     dbCfg["host"].(string),
		Port:     dbCfg["port"].(string),
		Database: dbCfg["database"].(string),
		Schema:   dbCfg["schema"].(string),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		SSL:      dbCfg["ssl"].(string),
	}, nil
}
