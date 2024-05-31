package internal

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/gertd/go-pluralize"
)

type ResourceParams struct {
	ProjectConfig *ProjectConfig
	Table         *TableDefinition
}

type UpdateRouterParams struct {
	ProjectConfig *ProjectConfig
	Tables        []*TableDefinition
}

type TableDefinition struct {
	Name    string              `yaml:"name"`
	Columns []*ColumnDefinition `yaml:"columns"`
}

type ColumnDefinition struct {
	Name         string `yaml:"name"`
	Type         string `yaml:"type"`
	Nullable     bool   `yaml:"nullable"`
	IsPrimaryKey bool   `yaml:"is_primary_key"`
	HasDefault   bool   `yaml:"has_default"`
}

func (t *TableDefinition) InputName() string {
	return t.Name + "Input"
}

func (t *TableDefinition) PluralName() string {
	client := pluralize.NewClient()

	if client.IsPlural(t.Name) {
		return t.Name
	} else {
		return client.Plural(t.Name)
	}
}

func (t *TableDefinition) SingularName() string {
	client := pluralize.NewClient()

	if client.IsSingular(t.Name) {
		return t.Name
	} else {
		return client.Singular(t.Name)
	}
}

func (t *TableDefinition) PrimaryKeys() []*ColumnDefinition {
	columns := make([]*ColumnDefinition, 0)
	for _, col := range t.Columns {
		if col.IsPrimaryKey {
			columns = append(columns, col)
		}
	}
	if len(columns) == 0 {
		return t.Columns
	}
	return columns
}

func (t *TableDefinition) CreateInputColumns() []*ColumnDefinition {
	columns := make([]*ColumnDefinition, 0)
	for _, col := range t.Columns {
		if col.IsPrimaryKey && col.HasDefault {
			continue
		}
		columns = append(columns, col)
	}
	if len(columns) == 0 {
		return t.Columns
	}
	return columns
}

func (t *TableDefinition) UpdateInputColumns() []*ColumnDefinition {
	columns := make([]*ColumnDefinition, 0)
	for _, col := range t.Columns {
		if !col.IsPrimaryKey {
			columns = append(columns, col)
		}
	}
	if len(columns) == 0 {
		return t.Columns
	}
	return columns
}

func (t *TableDefinition) ColumnsByComma() string {
	strArr := make([]string, 0)
	for _, col := range t.Columns {
		strArr = append(strArr, col.Name)
	}
	return strings.Join(strArr, ", ")
}

func (t *TableDefinition) PrimaryKeysByComma() string {
	strArr := make([]string, 0)
	for _, col := range t.PrimaryKeys() {
		strArr = append(strArr, col.Name)
	}
	return strings.Join(strArr, ", ")
}

func (t *TableDefinition) PrimaryKeysFuncParams() string {
	strArr := make([]string, 0)
	for _, col := range t.PrimaryKeys() {
		strArr = append(strArr, fmt.Sprintf("%s %s", col.Name, col.GoType()))
	}
	return strings.Join(strArr, ", ")
}

func (t *TableDefinition) PrimaryKeysWhereClause() string {
	strArr := make([]string, 0)
	for i, col := range t.PrimaryKeys() {
		if i == 0 {
			strArr = append(strArr, fmt.Sprintf("WHERE %s = $%d", col.Name, i+1))
		} else {
			strArr = append(strArr, fmt.Sprintf("AND %s = $%d", col.Name, i+1))
		}
	}
	return strings.Join(strArr, " ")
}

func (t *TableDefinition) PrimaryKeysEndpoint() string {
	strArr := make([]string, 0)
	for _, col := range t.PrimaryKeys() {
		strArr = append(strArr, ":"+col.Name)
	}
	return strings.Join(strArr, "/")
}

func (t *TableDefinition) ModelScanParams() string {
	strArr := make([]string, 0)
	for _, col := range t.Columns {
		strArr = append(strArr, fmt.Sprintf("&%s.%s", t.SingularName(), col.GoName()))
	}
	return strings.Join(strArr, ", ")
}

func (t *TableDefinition) CreateInputByComma() string {
	strArr := make([]string, 0)
	for _, col := range t.CreateInputColumns() {
		strArr = append(strArr, col.Name)
	}
	return strings.Join(strArr, ", ")
}

func (t *TableDefinition) CreateInputNumbersByComma() string {
	strArr := make([]string, 0)
	for idx := range t.CreateInputColumns() {
		strArr = append(strArr, fmt.Sprintf("$%d", idx+1))
	}
	return strings.Join(strArr, ", ")
}

func (t *TableDefinition) CreateInputParams() string {
	strArr := make([]string, 0)
	for _, col := range t.CreateInputColumns() {
		strArr = append(strArr, fmt.Sprintf("%s.%s", t.InputName(), col.GoName()))
	}
	return strings.Join(strArr, ", ")
}

func (t *TableDefinition) UpdateInputByComma() string {
	strArr := make([]string, 0)
	for _, col := range t.UpdateInputColumns() {
		strArr = append(strArr, col.Name)
	}
	return strings.Join(strArr, ", ")
}

func (t *TableDefinition) UpdateInputParams() string {
	strArr := make([]string, 0)
	for _, col := range t.UpdateInputColumns() {
		strArr = append(strArr, fmt.Sprintf("%s.%s", t.InputName(), col.GoName()))
	}
	return strings.Join(strArr, ", ")
}

func (t *TableDefinition) UpdateClause() string {
	str := "UPDATE " + t.Name + " SET "
	nParams := 1

	strArr := make([]string, 0)
	for _, col := range t.UpdateInputColumns() {
		strArr = append(strArr, fmt.Sprintf("%s = $%d", col.Name, nParams))
		nParams++
	}
	str += strings.Join(strArr, ", ")

	strArr = make([]string, 0)
	for i, col := range t.PrimaryKeys() {
		if i == 0 {
			strArr = append(strArr, fmt.Sprintf(" WHERE %s = $%d", col.Name, nParams))
		} else {
			strArr = append(strArr, fmt.Sprintf("AND %s = $%d", col.Name, nParams))
		}
		nParams++
	}
	str += strings.Join(strArr, " ")

	return str
}

func (c *ColumnDefinition) GoType() string {
	if c.Nullable {
		return postgresToNullableGoTypes[c.Type]
	}
	return postgresToGoTypes[c.Type]
}

func (c *ColumnDefinition) ParserFunc() string {
	return goTypesToParserFunc[postgresToGoTypes[c.Type]]
}

func (c *ColumnDefinition) GoName() string {
	if c.Name == "" {
		return ""
	}
	runes := []rune(c.Name)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
