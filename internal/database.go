package internal

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/manifoldco/promptui"
)

var postgresToGoTypes = map[string]string{
	"integer":                     "int",
	"bigint":                      "int64",
	"smallint":                    "int16",
	"boolean":                     "bool",
	"real":                        "float32",
	"double precision":            "float64",
	"numeric":                     "*big.Rat", // Arbitrary-precision numeric types can be represented using big.Rat in Go
	"money":                       "float64",
	"character varying":           "string",
	"text":                        "string",
	"date":                        "time.Time",
	"timestamp without time zone": "time.Time",
	"timestamp with time zone":    "time.Time",
	"json":                        "json.RawMessage",
	"jsonb":                       "json.RawMessage",
	"uuid":                        "uuid.UUID", // You'll need to use a package like github.com/google/uuid for this
	"bytea":                       "[]byte",    // Binary data can be represented as a byte slice in Go
	"point":                       "string",    // There's no direct equivalent for geometric types in Go, so you might want to use string and parse them manually
	"line":                        "string",
	"lseg":                        "string",
	"box":                         "string",
	"path":                        "string",
	"polygon":                     "string",
	"circle":                      "string",
	"cidr":                        "net.IPNet", // You can use the net package in Go to work with network addresses
	"inet":                        "net.IP",
	"macaddr":                     "net.HardwareAddr",
	"macaddr8":                    "net.HardwareAddr",
	"bit":                         "[]byte", // Bit strings can be represented as byte slices in Go
	"bit varying":                 "[]byte",
	"array":                       "[]interface{}",     // Arrays can be represented as slices in Go, but you'll need to handle them separately because the type name includes the element type (e.g., integer[])
	"hstore":                      "map[string]string", // Hstore can be represented as a map in Go, but you'll need a package like github.com/lib/pq to work with it
}

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
		tables = append(tables, tableName)
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
		var colType string
		var nullable string
		var defaultVal sql.NullString

		err = rows.Scan(&column.Name, &colType, &nullable, &defaultVal, &column.IsPrimaryKey)
		if err != nil {
			return nil, err
		}
		column.Type = postgresToGoTypes[colType]
		column.Nullable = nullable == "YES"
		column.HasDefault = defaultVal.Valid
		tableDef.Columns = append(tableDef.Columns, &column)
	}

	tableDef.Name = table

	return &tableDef, nil
}
