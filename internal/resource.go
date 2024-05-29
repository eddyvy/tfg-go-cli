package internal

import (
	"fmt"
	"strings"
	"unicode"
)

type ResourceParams struct {
	ProjectConfig *ProjectConfig
	Table         *TableDefinition
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

func (t *TableDefinition) PrimaryKeys() []*ColumnDefinition {
	columns := make([]*ColumnDefinition, 0)
	for _, col := range t.Columns {
		if col.IsPrimaryKey {
			columns = append(columns, col)
		}
	}
	return columns
}

func (t *TableDefinition) NotPrimaryKeys() []*ColumnDefinition {
	columns := make([]*ColumnDefinition, 0)
	for _, col := range t.Columns {
		if !col.IsPrimaryKey {
			columns = append(columns, col)
		}
	}
	return columns
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
		strArr = append(strArr, fmt.Sprintf("%s %s", col.Name, col.Type))
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
		strArr = append(strArr, fmt.Sprintf("&%s.%s", t.Name, col.GoName()))
	}
	return strings.Join(strArr, ", ")
}

func (c *ColumnDefinition) ParserFunc() string {
	return goTypesToParserFunc[c.Type]
}

func (c *ColumnDefinition) GoName() string {
	if c.Name == "" {
		return ""
	}
	runes := []rune(c.Name)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
