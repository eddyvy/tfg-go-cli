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

func ExecuteTemplatesBase(cfg *GlobalConfig) error {
	return executeTemplatesBaseRec("", cfg)
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
			// Read the template file
			tmplFile, err := embedTemplates.Open(BASE_PATH + nextPath)
			if err != nil {
				return err
			}
			defer tmplFile.Close()

			// Get the new file name
			newFileName := parseFileName(file.Name())

			// Parse the content of the template
			tmplBytes, err := io.ReadAll(tmplFile)
			if err != nil {
				return err
			}
			tmplContent, err := executeTmpl(string(tmplBytes), newFileName, cfg)
			if err != nil {
				return err
			}

			// Create the destination file
			destPath := filepath.Join(currentDir, relPath, newFileName)
			destFile, err := os.Create(destPath)
			if err != nil {
				return err
			}
			defer destFile.Close()

			// Write the content of the template to the destination file
			_, err = destFile.WriteString(tmplContent)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func parseFileName(fileName string) string {
	newFileName := fileName
	if strings.HasSuffix(fileName, ".tmpl") {
		newFileName = strings.TrimSuffix(fileName, ".tmpl")
	}
	return newFileName
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
