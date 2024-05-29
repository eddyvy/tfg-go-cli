package internal

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
	"github.com/manifoldco/promptui"
)

func ConfigureDatabase(cfg *GlobalConfig) error {
	fmt.Println("Connecting to database...")
	db, err := connectDatabase(cfg.DatabaseConfig)
	if err != nil {
		return err
	}
	defer db.Close()
	fmt.Println("Successfully connected to the database")

	tableNames, err := chooseTables(db, cfg.DatabaseConfig.Schema)
	if err != nil {
		return err
	}

	fmt.Println("Reading table details...")
	for _, tableName := range tableNames {
		tableDef, err := readTableData(db, tableName)
		if err != nil {
			return err
		}

		cfg.DatabaseConfig.Tables = append(cfg.DatabaseConfig.Tables, tableDef)
	}
	fmt.Println("Successfully read table details")

	return nil
}

func connectDatabase(cfg *DatabaseConfig) (*sql.DB, error) {
	if cfg.Type != "postgresql" {
		fmt.Println("Sorry, only \"postgresql\" database type is supported")
		return nil, fmt.Errorf("sorry, database type %s is not supported", cfg.Type)
	}

	connStr := cfg.ConnectionString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func chooseTables(db *sql.DB, schema string) ([]string, error) {
	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = $1", schema)
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
		tables = append(tables, strings.ToLower(tableName))
	}

	if tables == nil {
		return nil, fmt.Errorf("no tables found in the database")
	}

	prompt := promptui.Select{
		Label:        "Select table",
		Items:        append([]string{"All tables"}, tables...),
		HideSelected: true,
		HideHelp:     true,
	}

	_, result, err := prompt.Run()
	os.Stdout.Sync()

	if err != nil {
		return nil, err
	}

	if result == "All tables" {
		return tables, nil
	}

	return []string{result}, nil
}

func readTableData(db *sql.DB, table string) (*TableDefinition, error) {
	var tableDef TableDefinition
	rows, err := db.Query(`
			SELECT 
					c.column_name, 
					c.data_type, 
					c.is_nullable, 
					c.column_default,
					(
							SELECT COUNT(*)
							FROM information_schema.key_column_usage AS kcu
							INNER JOIN information_schema.table_constraints AS tc
							ON kcu.constraint_name = tc.constraint_name
							WHERE kcu.table_name = c.table_name
							AND kcu.column_name = c.column_name
							AND tc.constraint_type = 'PRIMARY KEY'
					) > 0 AS is_primary_key
			FROM information_schema.columns AS c
			WHERE c.table_name = $1
	`, table)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var column ColumnDefinition
		var columnName string
		var colType string
		var nullable string
		var defaultVal sql.NullString

		err = rows.Scan(&columnName, &colType, &nullable, &defaultVal, &column.IsPrimaryKey)
		if err != nil {
			return nil, err
		}
		column.Name = strings.ToLower(columnName)
		column.Type = postgresToGoTypes[colType]
		column.Nullable = nullable == "YES"
		column.HasDefault = defaultVal.Valid
		tableDef.Columns = append(tableDef.Columns, &column)
	}

	tableDef.Name = table

	return &tableDef, nil
}
