package internal

import (
	"bytes"
	"embed"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed * templates/base/.*
var embedTemplates embed.FS

const BASE_PATH = "templates/base"
const RESOURCE_PATH = "templates/resource"
const ROUTER_PATH_ENDPOINTS = "templates/router/endpoints.go.tmpl"
const ROUTER_PATH_IMPORTS = "templates/router/imports.go.tmpl"

func ExecuteTemplatesBase(cfg *GlobalConfig) error {
	return executeTemplatesBaseRec("", cfg)
}

func ExecuteTemplatesResources(cfg *GlobalConfig, useUpdating bool) error {
	resourcesTypes := []string{"handler", "model", "service"}

	// Path to Internal
	internalDir := filepath.Join(cfg.ProjectConfig.ProjectDir, "internal")

	var tables []*TableDefinition
	if useUpdating {
		tables = cfg.DatabaseConfig.UpdatingTables
	} else {
		tables = cfg.DatabaseConfig.Tables
	}

	// Execute the templates for each table
	for _, table := range tables {
		tableDir := filepath.Join(internalDir, table.Name)
		err := os.MkdirAll(tableDir, os.ModePerm)
		if err != nil {
			return err
		}

		resourcesParams := &ResourceParams{
			ProjectConfig: cfg.ProjectConfig,
			Table:         table,
		}

		for _, resType := range resourcesTypes {
			err = createFileFromTmpl(
				RESOURCE_PATH+"/resource_"+resType+".go.tmpl",
				tableDir,
				table.Name+"_"+resType+".go",
				resourcesParams)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func UpdateRouter(tables []*TableDefinition, cfg *GlobalConfig) error {
	config := &UpdateRouterParams{
		ProjectConfig: cfg.ProjectConfig,
		Tables:        tables,
	}

	// Read current router.go file
	routerPath := filepath.Join(cfg.ProjectConfig.ProjectDir, "router.go")
	routerData, err := os.ReadFile(routerPath)
	if err != nil {
		return err
	}
	source := string(routerData)

	// Execute the imports template
	importsTmplFile, err := embedTemplates.Open(ROUTER_PATH_IMPORTS)
	if err != nil {
		return err
	}
	defer importsTmplFile.Close()
	importsTmplBytes, err := io.ReadAll(importsTmplFile)
	if err != nil {
		return err
	}
	importsTmplContent, err := executeTmpl(string(importsTmplBytes), "imports.go.tmpl", config)
	if err != nil {
		return err
	}

	// Execute the endpoints template
	endpointsTmplFile, err := embedTemplates.Open(ROUTER_PATH_ENDPOINTS)
	if err != nil {
		return err
	}
	defer endpointsTmplFile.Close()
	endpointsTmplBytes, err := io.ReadAll(endpointsTmplFile)
	if err != nil {
		return err
	}
	endpointsTmplContent, err := executeTmpl(string(endpointsTmplBytes), "endpoints.go.tmpl", config)
	if err != nil {
		return err
	}

	// Insert the new imports and endpoints into the source code.
	source = strings.Replace(source, `"github.com/gofiber/fiber/v2"`, `"github.com/gofiber/fiber/v2"`+"\n"+importsTmplContent, 1)
	source = strings.Replace(source, `func SetRoutes(app *fiber.App) {`, `func SetRoutes(app *fiber.App) {`+"\n"+endpointsTmplContent, 1)

	// Write the modified source code back to the file.
	err = os.WriteFile(routerPath, []byte(source), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func executeTemplatesBaseRec(relPath string, cfg *GlobalConfig) error {
	files, err := fs.ReadDir(embedTemplates, BASE_PATH+relPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		destPath := filepath.Join(cfg.ProjectConfig.ProjectDir, relPath, file.Name())
		nextPath := relPath + "/" + file.Name()

		if file.IsDir() {
			// Create the directory in the destination

			err = os.MkdirAll(destPath, os.ModePerm)
			if err != nil {
				return err
			}

			// Recursively copy the files in the directory
			err = executeTemplatesBaseRec(nextPath, cfg)
			if err != nil {
				return err
			}
		} else {
			// Get the new file name
			newFileName := file.Name()
			if strings.HasSuffix(file.Name(), ".tmpl") {
				newFileName = strings.TrimSuffix(file.Name(), ".tmpl")
			}

			err = createFileFromTmpl(BASE_PATH+nextPath, filepath.Join(cfg.ProjectConfig.ProjectDir, relPath), newFileName, cfg)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func createFileFromTmpl(tmplEmbedPath, destDir, destFilename string, data interface{}) error {
	// Read template from embedded files
	tmplFile, err := embedTemplates.Open(tmplEmbedPath)
	if err != nil {
		return err
	}
	defer tmplFile.Close()

	// Read template file content bytes
	tmplBytes, err := io.ReadAll(tmplFile)
	if err != nil {
		return err
	}

	// Execute the template and get parsed string
	tmplContent, err := executeTmpl(string(tmplBytes), destFilename, data)
	if err != nil {
		return err
	}

	// Create final file
	tmplFilePath := filepath.Join(destDir, destFilename)
	tmplDestFile, err := os.Create(tmplFilePath)
	if err != nil {
		return err
	}
	defer tmplDestFile.Close()

	// Write the content of the template to the destination file
	_, err = tmplDestFile.WriteString(tmplContent)
	if err != nil {
		return err
	}

	return nil
}

func executeTmpl(tmplText, tmplName string, data interface{}) (string, error) {
	tmpl, err := template.New(tmplName).Parse(tmplText)
	if err != nil {
		return "", err
	}

	// Execute the template
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
