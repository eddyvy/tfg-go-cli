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

type ResourceParams struct {
	ProjectConfig *ProjectConfig
	Table         *TableDefinition
}

func (r *ResourceParams) EndpointOne() string {
	return "/" + r.Table.Name + "/{id}"
}

func ExecuteTemplatesBase(cfg *GlobalConfig) error {
	return executeTemplatesBaseRec("", cfg)
}

func ExecuteTemplatesResources(cfg *GlobalConfig) error {
	resourcesTypes := []string{"handler", "model", "service"}

	// Path to Internal
	execDir, err := os.Getwd()
	if err != nil {
		return err
	}
	internalDir := filepath.Join(execDir, cfg.ProjectConfig.Name, "internal")

	// Execute the templates for each table
	for _, table := range cfg.DatabaseConfig.Tables {
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

func executeTemplatesBaseRec(relPath string, cfg *GlobalConfig) error {
	files, err := fs.ReadDir(embedTemplates, BASE_PATH+relPath)
	if err != nil {
		return err
	}

	execDir, err := os.Getwd()
	if err != nil {
		return err
	}
	currentDir := filepath.Join(execDir, cfg.ProjectConfig.Name)

	for _, file := range files {
		destPath := filepath.Join(currentDir, relPath, file.Name())
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

			err = createFileFromTmpl(BASE_PATH+nextPath, filepath.Join(currentDir, relPath), newFileName, cfg)
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
