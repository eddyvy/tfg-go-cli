package internal

var postgresToGoTypes = map[string]string{
	"integer":                     "int",
	"bigint":                      "int64",
	"smallint":                    "int16",
	"boolean":                     "bool",
	"real":                        "float32",
	"double precision":            "float64",
	"numeric":                     "*big.Rat",
	"money":                       "float64",
	"character varying":           "string",
	"text":                        "string",
	"date":                        "time.Time",
	"timestamp without time zone": "time.Time",
	"timestamp with time zone":    "time.Time",
	"json":                        "json.RawMessage",
	"jsonb":                       "json.RawMessage",
	"uuid":                        "uuid.UUID",
	"bytea":                       "[]byte",
	"point":                       "string",
	"line":                        "string",
	"lseg":                        "string",
	"box":                         "string",
	"path":                        "string",
	"polygon":                     "string",
	"circle":                      "string",
	"cidr":                        "net.IPNet",
	"inet":                        "net.IP",
	"macaddr":                     "net.HardwareAddr",
	"macaddr8":                    "net.HardwareAddr",
	"bit":                         "[]byte",
	"bit varying":                 "[]byte",
	"array":                       "[]interface{}",
	"hstore":                      "map[string]string",
	"character":                   "string",
	"time without time zone":      "time.Time",
	"time with time zone":         "time.Time",
	"interval":                    "time.Duration",
	"tsvector":                    "string",
	"tsquery":                     "string",
	"oid":                         "uint32",
	"xml":                         "string",
	"char":                        "string",
	"name":                        "string",
	"bpchar":                      "string",
	"void":                        "interface{}",
	"int2vector":                  "[]int16",
	"int4range":                   "pg.Range",
	"int8range":                   "pg.Range",
	"numrange":                    "pg.Range",
	"tsrange":                     "pg.Range",
	"tstzrange":                   "pg.Range",
	"daterange":                   "pg.Range",
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
	"character":                   "parser.NullString",
	"time without time zone":      "parser.NullTime",
	"time with time zone":         "parser.NullTime",
	"interval":                    "time.Duration",
	"tsvector":                    "parser.NullString",
	"tsquery":                     "parser.NullString",
	"oid":                         "parser.NullInt32",
	"xml":                         "parser.NullString",
	"char":                        "parser.NullString",
	"name":                        "parser.NullString",
	"bpchar":                      "parser.NullString",
	"void":                        "interface{}",
	"int2vector":                  "[]parser.NullInt32",
	"int4range":                   "pg.Range",
	"int8range":                   "pg.Range",
	"numrange":                    "pg.Range",
	"tsrange":                     "pg.Range",
	"tstzrange":                   "pg.Range",
	"daterange":                   "pg.Range",
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
