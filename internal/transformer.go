package internal

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

var postgresToNullableGoTypes = map[string]string{
	"integer":                     "parser.NullInt32",
	"bigint":                      "parser.NullInt64",
	"smallint":                    "parser.NullInt32",
	"boolean":                     "parser.NullBool",
	"real":                        "parser.NullFloat64",
	"double precision":            "parser.NullFloat64",
	"numeric":                     "*big.Rat",
	"money":                       "parser.NullFloat64",
	"character varying":           "parser.NullString",
	"text":                        "parser.NullString",
	"date":                        "parser.NullTime",
	"timestamp without time zone": "parser.NullTime",
	"timestamp with time zone":    "parser.NullTime",
	"json":                        "json.RawMessage",
	"jsonb":                       "json.RawMessage",
	"uuid":                        "uuid.UUID",
	"bytea":                       "[]byte",
	"point":                       "parser.NullString",
	"line":                        "parser.NullString",
	"lseg":                        "parser.NullString",
	"box":                         "parser.NullString",
	"path":                        "parser.NullString",
	"polygon":                     "parser.NullString",
	"circle":                      "parser.NullString",
	"cidr":                        "net.IPNet",
	"inet":                        "net.IP",
	"macaddr":                     "net.HardwareAddr",
	"macaddr8":                    "net.HardwareAddr",
	"bit":                         "[]byte",
	"bit varying":                 "[]byte",
	"array":                       "[]interface{}",
	"hstore":                      "map[string]string",
}

var goTypesToParserFunc = map[string]string{
	"int":               "StringToInt",
	"int64":             "StringToInt64",
	"int16":             "StringToInt16",
	"bool":              "StringToBool",
	"float32":           "StringToFloat32",
	"float64":           "StringToFloat64",
	"*big.Rat":          "StringToRat",
	"time.Time":         "StringToTime",
	"json.RawMessage":   "StringToJSON",
	"uuid.UUID":         "StringToUUID",
	"[]byte":            "StringToBytes",
	"string":            "",
	"net.IPNet":         "StringToIPNet",
	"net.IP":            "StringToIP",
	"net.HardwareAddr":  "StringToHardwareAddr",
	"[]interface{}":     "StringToInterfaceSlice",
	"map[string]string": "StringToStringMap",
}
