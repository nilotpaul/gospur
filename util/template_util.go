package util

import (
	"embed"
	"html/template"
	"os"
	"path/filepath"

	"github.com/nilotpaul/gospur/config"
	tmpls "github.com/nilotpaul/gospur/template"
)

type processTemplate struct {
	targetFilePath string
	template       *template.Template
}

func ProcessTemplates(targetDir string, data interface{}) error {
	for targetPath, templatePath := range config.ProjectStructure {
		// Getting the folder containing all base templates (files that doesn't depend on config).
		baseTmplFS := tmpls.GetBaseFiles()

		// `targetFilePath` is the final path where the file will be stored.
		targetFilePath := filepath.Join(targetDir, targetPath)

		processedTmpl, err := parseTemplate(targetFilePath, templatePath, baseTmplFS)
		if err != nil {
			return err
		}

		if err := createTargetFile(processedTmpl.targetFilePath, processedTmpl.template, nil); err != nil {
			return err
		}
	}

	return nil
}

// parseTemplate takes fullWritePath, template path and template embed.
// tmplPath -> path where the template is stored
// tmplFS -> template embed FS which contains all template files.
func parseTemplate(fullWritePath, tmplPath string, tmplFS embed.FS) (*processTemplate, error) {
	baseTmplBytes, err := tmplFS.ReadFile(tmplPath)
	if err != nil {
		return nil, err
	}

	// Parsing the tmpl bytes(file contents) to get the actual template.
	tmpl, err := template.New(filepath.Base(tmplPath)).Parse(string(baseTmplBytes))
	if err != nil {
		return nil, err
	}

	return &processTemplate{targetFilePath: fullWritePath, template: tmpl}, nil
}

// createTargetFile takes a fullWritePath and parsed template where it'll write
// the contents. fullWritePath has to be joined to the project or targetPath. (eg. gospur/config/env.go)
func createTargetFile(fullWritePath string, tmpl *template.Template, data interface{}) error {
	// Create parent directories for the target file.
	// Here second arg of `CreateTargetDir` is false which depicts write even
	// if the directory is not empty.
	if err := CreateTargetDir(filepath.Dir(fullWritePath), false); err != nil {
		return err
	}

	// Create the file in the target file path.
	destFile, err := os.Create(fullWritePath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Execute the template and write the output to the file.
	if err := tmpl.Execute(destFile, data); err != nil {
		return err
	}

	return nil
}
